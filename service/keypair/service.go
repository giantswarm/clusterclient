package keypair

import (
	"net/url"

	micrologger "github.com/giantswarm/microkit/logger"
	"github.com/go-resty/resty"

	"github.com/giantswarm/clusterclient/service/keypair/creator"
	"github.com/giantswarm/clusterclient/service/keypair/lister"
)

// Config represents the configuration used to create a new service.
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

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	var err error

	var creatorService *creator.Service
	{
		creatorConfig := creator.DefaultConfig()
		creatorConfig.Logger = config.Logger
		creatorConfig.RestClient = config.RestClient
		creatorConfig.URL = config.URL
		creatorService, err = creator.New(creatorConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	var listerService *lister.Service
	{
		listerConfig := lister.DefaultConfig()
		listerConfig.Logger = config.Logger
		listerConfig.RestClient = config.RestClient
		listerConfig.URL = config.URL
		listerService, err = lister.New(listerConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	newService := &Service{
		Creator: creatorService,
		Lister:  listerService,
	}

	return newService, nil
}

type Service struct {
	Creator *creator.Service
	Lister  *lister.Service
}
