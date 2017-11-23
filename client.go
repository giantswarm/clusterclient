// Package clusterclient implements business logic to request the
// cluster-service.
package clusterclient

import (
	"net/url"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/go-resty/resty"

	"github.com/giantswarm/clusterclient/service/cluster"
	"github.com/giantswarm/clusterclient/service/info"
	"github.com/giantswarm/clusterclient/service/keypair"
	"github.com/giantswarm/clusterclient/service/release"
	"github.com/giantswarm/clusterclient/service/root"
)

// Config represents the configuration used to create a new client.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	RestClient *resty.Client

	// Settings.
	Address string
}

// DefaultConfig provides a default configuration to create a new client by best
// effort.
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
		Address: "http://127.0.0.1:8080",
	}

	return config
}

type Client struct {
	Cluster *cluster.Service
	Info    *info.Service
	KeyPair *keypair.Service
	Release *release.Release
	Root    *root.Service
}

func New(config Config) (*Client, error) {
	if config.Address == "" {
		return nil, microerror.Maskf(invalidConfigError, "address must not be empty")
	}

	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var clusterService *cluster.Service
	{
		clusterConfig := cluster.DefaultConfig()
		clusterConfig.Logger = config.Logger
		clusterConfig.RestClient = config.RestClient
		clusterConfig.URL = u
		clusterService, err = cluster.New(clusterConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var infoService *info.Service
	{
		infoConfig := info.DefaultConfig()
		infoConfig.Logger = config.Logger
		infoConfig.RestClient = config.RestClient
		infoConfig.URL = u
		infoService, err = info.New(infoConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var keypairService *keypair.Service
	{
		keypairConfig := keypair.DefaultConfig()
		keypairConfig.Logger = config.Logger
		keypairConfig.RestClient = config.RestClient
		keypairConfig.URL = u
		keypairService, err = keypair.New(keypairConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var releaseService *release.Release
	{
		c := release.DefaultConfig()

		c.Logger = config.Logger
		c.RestClient = config.RestClient

		c.URL = u

		releaseService, err = release.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var rootService *root.Service
	{
		rootConfig := root.DefaultConfig()
		rootConfig.RestClient = config.RestClient
		rootConfig.URL = u
		rootService, err = root.New(rootConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newClient := &Client{
		Cluster: clusterService,
		Info:    infoService,
		KeyPair: keypairService,
		Release: releaseService,
		Root:    rootService,
	}

	return newClient, nil
}
