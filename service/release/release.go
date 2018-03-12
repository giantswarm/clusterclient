package release

import (
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"

	"github.com/giantswarm/clusterclient/service/release/lister"
	"github.com/giantswarm/clusterclient/service/release/searcher"
	"github.com/giantswarm/clusterclient/service/release/validator"
)

type Config struct {
	Logger     micrologger.Logger
	RestClient *resty.Client

	URL *url.URL
}

func DefaultConfig() Config {
	return Config{
		Logger:     nil,
		RestClient: nil,

		URL: nil,
	}
}

type Release struct {
	Lister    *lister.Lister
	Searcher  *searcher.Searcher
	Validator *validator.Validator
}

func New(config Config) (*Release, error) {
	var err error

	var newLister *lister.Lister
	{
		c := lister.DefaultConfig()

		c.Logger = config.Logger
		c.RestClient = config.RestClient

		c.URL = config.URL

		newLister, err = lister.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var newSearcher *searcher.Searcher
	{
		c := searcher.DefaultConfig()

		c.Logger = config.Logger
		c.RestClient = config.RestClient

		c.URL = config.URL

		newSearcher, err = searcher.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var newValidator *validator.Validator
	{
		c := validator.Config{
			Logger:     config.Logger,
			RestClient: config.RestClient,

			URL: config.URL,
		}

		newValidator, err = validator.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	r := &Release{
		Lister:    newLister,
		Searcher:  newSearcher,
		Validator: newValidator,
	}

	return r, nil
}
