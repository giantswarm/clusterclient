package release

import (
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"

	"github.com/giantswarm/clusterclient/service/release/lister"
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
	Lister *lister.Lister
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

	r := &Release{
		Lister: newLister,
	}

	return r, nil
}
