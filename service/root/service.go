package root

import (
	"fmt"
	"net/url"

	micrologger "github.com/giantswarm/microkit/logger"
	"github.com/go-resty/resty"
)

const (
	// Endpoint is the API endpoint of the service this client action interacts
	// with.
	Endpoint = "/"
)

// Config represents the configuration used to create a root service.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new root service
// by best effort.
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

// New creates a new configured root service.
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

// Service implements the root service action.
type Service struct {
	Config
}

func (s *Service) Get(request Request) (*Response, error) {
	u, err := s.URL.Parse(Endpoint)
	if err != nil {
		return nil, maskAny(err)
	}

	r, err := s.RestClient.R().SetResult(DefaultResponse()).Get(u.String())
	if err != nil {
		return nil, maskAny(err)
	}

	if r.StatusCode() != 200 {
		return nil, maskAny(fmt.Errorf(string(r.Body())))
	}

	response := r.Result().(*Response)

	return response, nil
}
