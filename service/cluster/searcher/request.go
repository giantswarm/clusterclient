package searcher

import "github.com/giantswarm/clusterclient/service/cluster/searcher/config"

// Request is the configuration for the service action.
type Request struct {
	Cluster *config.Cluster `json:"cluster"`
}

// DefaultRequest provides a default request by best effort.
func DefaultRequest() Request {
	return Request{
		Cluster: config.DefaultCluster(),
	}
}
