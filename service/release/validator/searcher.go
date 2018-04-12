package validator

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty"

	"github.com/giantswarm/microclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	Endpoint = "/v1/releases/%s/%s/validator/"
	Name     = "release/validator"
)

type Config struct {
	Logger     micrologger.Logger
	RestClient *resty.Client

	URL *url.URL
}

type Validator struct {
	logger     micrologger.Logger
	restClient *resty.Client

	url *url.URL
}

func New(config Config) (*Validator, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.RestClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.RestClient must not be empty", config)
	}

	if config.URL == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.URL must not be empty", config)
	}

	s := &Validator{
		logger:     config.Logger,
		restClient: config.RestClient,

		url: config.URL,
	}

	return s, nil
}

func (s *Validator) IsAllowed(ctx context.Context, request Request) (Response, error) {
	var err error

	var u *url.URL
	{
		u, err = s.url.Parse(fmt.Sprintf(Endpoint, request.CurrentRelease, request.DesiredRelease))
		if err != nil {
			return Response{}, microerror.Mask(err)
		}
	}

	var response Response
	{
		s.logger.Log("level", "debug", "message", "sending GET request", "service", Name, "url", u.String())

		r, err := microclient.Do(ctx, s.restClient.R().SetResult(Response{}).Get, u.String())
		if err != nil {
			return Response{}, microerror.Mask(err)
		}

		s.logger.Log("code", r.StatusCode(), "level", "debug", "message", "received status code", "service", Name, "url", u.String())

		if r.StatusCode() == http.StatusNotFound {
			return Response{}, microerror.Mask(notFoundError)
		} else if r.StatusCode() != http.StatusOK {
			return Response{}, microerror.Mask(fmt.Errorf(string(r.Body())))
		}

		response = *(r.Result().(*Response))
	}

	return response, nil
}
