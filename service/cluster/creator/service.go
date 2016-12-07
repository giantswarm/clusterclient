package creator

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
)

const (
	// Endpoint is the API endpoint of the service this client action interacts
	// with.
	Endpoint = "/v1/clusters/"
)

// Config represents the configuration used to create a creator service.
type Config struct {
	// Dependencies.
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new creator
// service by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		RestClient: resty.New(),

		// Settings.
		URL: nil,
	}

	return newConfig
}

// New creates a new configured creator service.
func New(config Config) (*Service, error) {
	// Dependencies.
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

func (s *Service) Create(request Request) (*Response, error) {
	// At first we are going to create a new cluster resource. The result in case
	// the requested resource was created successfully will be a response
	// containing information about the location of the created resource.
	var resourceLocation string
	{
		u, err := s.URL.Parse(Endpoint)
		if err != nil {
			return nil, maskAny(err)
		}
		r, err := s.RestClient.R().SetBody(request).Post(u.String())
		if err != nil {
			return nil, maskAny(err)
		}
		if r.StatusCode() != 201 {
			return nil, maskAny(fmt.Errorf(string(r.Body())))
		}
		resourceLocation = r.Header().Get("Location")
	}

	// We know the location of the created resource from the response location
	// header. Now we request it to return relevant information about the created
	// resource in our response.
	var response *Response
	{
		u, err := s.URL.Parse(resourceLocation)
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
		response = r.Result().(*Response)
	}

	return response, nil
}
