package lister

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"
	"golang.org/x/net/context"
)

const (
	Endpoint = "/v1/releases/"
	Name     = "release/lister"
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

type Lister struct {
	logger     micrologger.Logger
	restClient *resty.Client

	url *url.URL
}

func New(config Config) (*Lister, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "logger must not be empty")
	}
	if config.RestClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "rest client must not be empty")
	}

	if config.URL == nil {
		return nil, microerror.Maskf(invalidConfigError, "URL must not be empty")
	}

	l := &Lister{
		logger:     config.Logger,
		restClient: config.RestClient,

		url: config.URL,
	}

	return l, nil
}

func (l *Lister) List(ctx context.Context) ([]Response, error) {
	var err error

	var u *url.URL
	{
		u, err = l.url.Parse(Endpoint)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var response []Response
	{
		l.logger.Log("debug", fmt.Sprintf("sending GET request to '%s'", u.String()), "service", Name)
		r, err := l.restClient.R().SetResult(DefaultResponse()).Get(u.String())
		if err != nil {
			return nil, microerror.Mask(err)
		}
		l.logger.Log("debug", fmt.Sprintf("received status code '%d'", r.StatusCode()), "service", Name)

		if r.StatusCode() != http.StatusOK {
			return nil, microerror.Mask(fmt.Errorf(string(r.Body())))
		}

		response = *(r.Result().(*[]Response))
	}

	return response, nil
}