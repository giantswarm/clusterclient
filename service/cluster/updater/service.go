package updater

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"
	"golang.org/x/net/context"
)

const (
	// Endpoint is the API endpoint of the service this client action interacts
	// with.
	Endpoint = "/v1/clusters/%s/"
	// Name is the service name being implemented. This can be used for e.g.
	// logging.
	Name = "cluster/updater"
)

// Config represents the configuration used to create a service.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new service by
// best effort.
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

// New creates a new configured service.
func New(config Config) (*Service, error) {
	// Dependencies.
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "logger must not be empty")
	}
	if config.RestClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "rest client must not be empty")
	}

	// Settings.
	if config.URL == nil {
		return nil, microerror.Maskf(invalidConfigError, "URL must not be empty")
	}

	newService := &Service{
		Config: config,
	}

	return newService, nil
}

type Service struct {
	Config
}

func (s *Service) Update(ctx context.Context, request Request) (*Response, error) {
	u, err := s.URL.Parse(fmt.Sprintf(Endpoint, request.Cluster.ID))
	if err != nil {
		return nil, microerror.Mask(err)
	}

	s.Logger.Log("debug", fmt.Sprintf("sending PATCH request to %s", u.String()), "service", Name)
	r, err := s.RestClient.R().SetBody(request.Cluster.Patch).SetResult(DefaultResponse()).Patch(u.String())
	if err != nil {
		return nil, microerror.Mask(err)
	}
	s.Logger.Log("debug", fmt.Sprintf("received status code %d", r.StatusCode()), "service", Name)

	if r.StatusCode() == http.StatusBadRequest {
		responseError := responseError{}

		parseErr := json.Unmarshal(r.Body(), &responseError)
		if parseErr != nil {
			return nil, microerror.Maskf(invalidRequestError, string(r.Body()))
		}

		return nil, microerror.Maskf(invalidRequestError, responseError.Error)
	} else if r.StatusCode() != http.StatusOK {
		return nil, microerror.Mask(fmt.Errorf(string(r.Body())))
	}

	response := r.Result().(*Response)

	return response, nil
}
