// Package clusterclient implements business logic to request the
// cluster-service.
package clusterclient

import (
	"net/url"

	micrologger "github.com/giantswarm/microkit/logger"
	"github.com/go-resty/resty"

	"github.com/giantswarm/clusterclient/service/cluster"
	"github.com/giantswarm/clusterclient/service/keypair"
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
	return Config{
		// Dependencies.
		Logger:     nil,
		RestClient: resty.New(),

		// Settings.
		Address: "http://127.0.0.1:8080",
	}
}

// New creates a new configured client.
func New(config Config) (*Client, error) {
	// Dependencies.
	if config.RestClient == nil {
		return nil, maskAnyf(invalidConfigError, "rest client must not be empty")
	}

	// Settings.
	if config.Address == "" {
		return nil, maskAnyf(invalidConfigError, "address must not be empty")
	}

	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, maskAny(err)
	}

	var clusterService *cluster.Service
	{
		clusterConfig := cluster.DefaultConfig()
		clusterConfig.Logger = config.Logger
		clusterConfig.RestClient = config.RestClient
		clusterConfig.URL = u
		clusterService, err = cluster.New(clusterConfig)
		if err != nil {
			return nil, maskAny(err)
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
			return nil, maskAny(err)
		}
	}

	var rootService *root.Service
	{
		rootConfig := root.DefaultConfig()
		rootConfig.RestClient = config.RestClient
		rootConfig.URL = u
		rootService, err = root.New(rootConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	newClient := &Client{
		Cluster: clusterService,
		KeyPair: keypairService,
		Root:    rootService,
	}

	return newClient, nil
}

type Client struct {
	Cluster *cluster.Service
	KeyPair *keypair.Service
	Root    *root.Service
}
