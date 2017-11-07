package creator

import (
	"github.com/giantswarm/clusterclient/service/cluster/creator/request"
	"github.com/giantswarm/clusterclient/service/cluster/creator/request/aws"
)

// Request is the configuration for the service action.
type Request struct {
	AWS               aws.Cluster      `json:"aws,omitempty"`
	KubernetesVersion string           `json:"kubernetes_version,omitempty"`
	Masters           []request.Master `json:"masters,omitempty"`
	Name              string           `json:"name,omitempty"`
	Owner             string           `json:"owner,omitempty"`
	ReleaseVersion    string           `json:"release_version,omitempty"`
	Workers           []request.Worker `json:"workers,omitempty"`
}

// DefaultRequest provides a default request object by best effort.
func DefaultRequest() Request {
	return Request{
		AWS:               aws.DefaultCluster(),
		KubernetesVersion: "",
		Masters:           []request.Master{},
		Name:              "",
		Owner:             "",
		ReleaseVersion:    "",
		Workers:           []request.Worker{},
	}
}
