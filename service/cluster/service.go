package cluster

import (
	"net/url"

	"github.com/go-resty/resty"

	"github.com/giantswarm/clusterclient/service/cluster/creator"
	"github.com/giantswarm/clusterclient/service/cluster/deleter"
	"github.com/giantswarm/clusterclient/service/cluster/lister"
	"github.com/giantswarm/clusterclient/service/cluster/searcher"
)

// Config represents the configuration used to create a new service.
type Config struct {
	// Dependencies.
	RestClient *resty.Client

	// Settings.
	URL *url.URL
}

// DefaultConfig provides a default configuration to create a new service by
// best effort.
func DefaultConfig() Config {
	return Config{
		// Dependencies.
		RestClient: resty.New(),

		// Settings.
		URL: nil,
	}
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	// Dependencies.
	if config.RestClient == nil {
		return nil, maskAnyf(invalidConfigError, "rest client must not be empty")
	}

	// Settings.
	if config.URL == nil {
		return nil, maskAnyf(invalidConfigError, "URL must not be empty")
	}

	var err error

	var creatorService *creator.Service
	{
		creatorConfig := creator.DefaultConfig()
		creatorConfig.RestClient = config.RestClient
		creatorConfig.URL = config.URL
		creatorService, err = creator.New(creatorConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	var deleterService *deleter.Service
	{
		deleterConfig := deleter.DefaultConfig()
		deleterConfig.RestClient = config.RestClient
		deleterConfig.URL = config.URL
		deleterService, err = deleter.New(deleterConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	var listerService *lister.Service
	{
		listerConfig := lister.DefaultConfig()
		listerConfig.RestClient = config.RestClient
		listerConfig.URL = config.URL
		listerService, err = lister.New(listerConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	var searcherService *searcher.Service
	{
		searcherConfig := searcher.DefaultConfig()
		searcherConfig.RestClient = config.RestClient
		searcherConfig.URL = config.URL
		searcherService, err = searcher.New(searcherConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	newService := &Service{
		Creator:  creatorService,
		Deleter:  deleterService,
		Lister:   listerService,
		Searcher: searcherService,
	}

	return newService, nil
}

type Service struct {
	Creator  *creator.Service
	Deleter  *deleter.Service
	Lister   *lister.Service
	Searcher *searcher.Service
}
