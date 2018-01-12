package info

import (
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"

	"github.com/giantswarm/clusterclient/service/info/aws"
	"github.com/giantswarm/clusterclient/service/info/kvm"
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
	return Config{
		Logger:     nil,
		RestClient: nil,

		URL: nil,
	}
}

type Service struct {
	AWS *aws.Service
	KVM *kvm.Service
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	var err error

	var awsService *aws.Service
	{
		awsConfig := aws.DefaultConfig()

		awsConfig.Logger = config.Logger
		awsConfig.RestClient = config.RestClient
		awsConfig.URL = config.URL

		awsService, err = aws.New(awsConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var kvmService *kvm.Service
	{
		kvmConfig := kvm.DefaultConfig()

		kvmConfig.Logger = config.Logger
		kvmConfig.RestClient = config.RestClient
		kvmConfig.URL = config.URL

		kvmService, err = kvm.New(kvmConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newService := &Service{
		AWS: awsService,
		KVM: kvmService,
	}

	return newService, nil
}
