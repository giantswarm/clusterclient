package creator

import "github.com/giantswarm/clusterclient/service/cluster/creator/config"

// Request is the configuration for the service action.
type Request struct {
	Cluster  *config.Cluster  `json:"cluster"`
	Customer *config.Customer `json:"customer"`
}

// DefaultRequest provides a default request by best effort.
func DefaultRequest() Request {
	return Request{
		Cluster:  config.DefaultCluster(),
		Customer: config.DefaultCustomer(),
	}
}
