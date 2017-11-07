package creator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/giantswarm/microerror"
	microserver "github.com/giantswarm/microkit/server"
	transactionid "github.com/giantswarm/microkit/transaction/context/id"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"
	"golang.org/x/net/context"

	"github.com/giantswarm/clusterclient/service/cluster/searcher"
)

const (
	// Endpoint is the API endpoint of the service this client action interacts
	// with.
	Endpoint = "/v1/clusters/"
	// Name is the service name being implemented. This can be used for e.g.
	// logging.
	Name = "cluster/creator"
)

// Config represents the configuration used to create a creator service.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new creator
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

type Service struct {
	Config
}

// New creates a new configured creator service.
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

func (s *Service) Create(ctx context.Context, request Request) (*Response, error) {
	// At first we are going to create a new cluster resource. The result in case
	// the requested resource was created successfully will be a response
	// containing information about the location of the created resource.
	var resourceLocation string
	{
		req := s.RestClient.R()
		req.SetBody(request)

		transactionID, ok := transactionid.FromContext(ctx)
		if ok {
			req.SetHeader(microserver.TransactionIDHeader, transactionID)
		}

		u, err := s.URL.Parse(Endpoint)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		s.Logger.Log("debug", fmt.Sprintf("sending POST request to %s", u.String()), "service", Name)

		res, err := req.Post(u.String())
		if err != nil {
			return nil, microerror.Mask(err)
		}
		s.Logger.Log("debug", fmt.Sprintf("received status code %d", res.StatusCode()), "service", Name)

		if res.StatusCode() == http.StatusBadRequest {
			responseError := responseError{}

			parseErr := json.Unmarshal(res.Body(), &responseError)
			if parseErr != nil {
				return nil, microerror.Maskf(invalidRequestError, string(res.Body()))
			}

			return nil, microerror.Maskf(invalidRequestError, responseError.Error)
		} else if res.StatusCode() != http.StatusCreated {
			return nil, microerror.Mask(fmt.Errorf(string(res.Body())))
		}

		resourceLocation = res.Header().Get("Location")
	}

	// We know the location of the created resource from the response location
	// header. Now we request it to return relevant information about the created
	// resource in our response.
	var response *Response
	{
		u, err := s.URL.Parse(resourceLocation)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		s.Logger.Log("debug", fmt.Sprintf("sending GET request to %s", u.String()), "service", Name)
		r, err := s.RestClient.R().SetResult(searcher.DefaultResponse()).Get(u.String())
		if err != nil {
			return nil, microerror.Mask(err)
		}
		s.Logger.Log("debug", fmt.Sprintf("received status code %d", r.StatusCode()), "service", Name)

		if r.StatusCode() == http.StatusNotFound {
			return nil, microerror.Mask(notFoundError)
		} else if r.StatusCode() != http.StatusOK {
			return nil, microerror.Mask(fmt.Errorf(string(r.Body())))
		}

		clientResponse := r.Result().(*searcher.Response)
		response = DefaultResponse()
		response.Cluster.ID = clientResponse.ID
	}

	return response, nil
}
