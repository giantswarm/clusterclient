package creator

import "github.com/giantswarm/clusterclient/service/cluster/creator/config"

// Request is the configuration for the service action.
type Request struct {
	Name string `json:"name,omitempty"`

	Owner string `json:"owner,omitempty"`

	KubernetesVersion string `json:"kubernetes_version,omitempty"`

	Masters []*config.Master `json:"masters,omitempty"`
	Workers []*config.Worker `json:"workers,omitempty"`
}

// DefaultRequest provides a default request object by best effort.
func DefaultRequest() Request {
	return Request{
		Name: "",

		Owner: "",

		KubernetesVersion: "",

		Masters: []*config.Master{},
		Workers: []*config.Worker{},
	}
}
