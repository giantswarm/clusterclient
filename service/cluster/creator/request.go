package creator

import (
	"github.com/giantswarm/clusterclient/service/cluster/creator/request"
	"github.com/giantswarm/clusterclient/service/cluster/creator/request/aws"
)

// Request is the configuration for the service action.
type Request struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Owner string `json:"owner,omitempty"`

	KubernetesVersion string `json:"kubernetes_version,omitempty"`

	AWS aws.Cluster `json:"aws,omitempty"`

	Masters []request.Master `json:"masters,omitempty"`
	Workers []request.Worker `json:"workers,omitempty"`
}

// DefaultRequest provides a default request object by best effort.
func DefaultRequest() Request {
	return Request{
		ID:    "",
		Name:  "",
		Owner: "",

		KubernetesVersion: "",

		AWS: aws.DefaultCluster(),

		Masters: []request.Master{},
		Workers: []request.Worker{},
	}
}
