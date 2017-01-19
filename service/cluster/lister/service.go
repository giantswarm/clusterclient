package lister

import (
	"fmt"
	"net/http"
	"net/url"

	micrologger "github.com/giantswarm/microkit/logger"
	"github.com/go-resty/resty"
)

const (
	// Endpoint is the API endpoint of the service this client action interacts
	// with.
	Endpoint = "/v1/organizations/%s/clusters/"
	// Name is the service name being implemented. This can be used for e.g.
	// logging.
	Name = "cluster/lister"
)

// Config represents the configuration used to create a lister service.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new lister
// service by best effort.
func DefaultConfig() Config {
	var err error

	var newLogger micrologger.Logger
	{
		loggerConfig := micrologger.DefaultConfig()
		newLogger, err = micrologger.New(loggerConfig)
		if err != nil {
			panic(err)
		}
	}

	config := Config{
		// Dependencies.
		Logger:     newLogger,
		RestClient: resty.New(),

		// Settings.
		URL: nil,
	}

	return config
}

// New creates a new configured lister service.
func New(config Config) (*Service, error) {
	// Dependencies.
	if config.Logger == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if config.RestClient == nil {
		return nil, maskAnyf(invalidConfigError, "rest client must not be empty")
	}

	// Settings.
	if config.URL == nil {
		return nil, maskAnyf(invalidConfigError, "URL must not be empty")
	}

	newService := &Service{
		Config: config,
	}

	return newService, nil
}

type Service struct {
	Config
}

func (s *Service) List(request Request) ([]*Response, error) {
	u, err := s.URL.Parse(fmt.Sprintf(Endpoint, request.Organization.ID))
	if err != nil {
		return nil, maskAny(err)
	}

	s.Logger.Log("debug", fmt.Sprintf("sending GET request to %s", u.String()), "service", Name)
	r, err := s.RestClient.R().SetBody(request).SetResult(DefaultResponse()).Get(u.String())
	if err != nil {
		return nil, maskAny(err)
	}
	s.Logger.Log("debug", fmt.Sprintf("received status code %d", r.StatusCode()), "service", Name)

	if r.StatusCode() == http.StatusNotFound {
		return nil, maskAny(notFoundError)
	} else if r.StatusCode() != http.StatusOK {
		return nil, maskAny(fmt.Errorf(string(r.Body())))
	}

	response := *(r.Result().(*[]*Response))

	return response, nil
}
