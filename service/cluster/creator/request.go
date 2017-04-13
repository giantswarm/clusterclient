package creator

import (
	"time"

	"github.com/giantswarm/clusterclient/service/cluster/creator/config"
)

// Request is the configuration for the service action.
type Request struct {
	// TODO remove this when v3 is no more used
	APIEndpoint string    `json:"api_endpoint"`
	CreateDate  time.Time `json:"create_date"`
	ID          string    `json:"id"`
	// TODO remove this when v3 is no more used

	Name string `json:"name,omitempty"`

	Owner string `json:"owner,omitempty"`

	KubernetesVersion string `json:"kubernetes_version,omitempty"`

	AWSSettings config.ClusterAWSSettings `json:"aws,omitempty"`

	Masters []*config.Master `json:"masters,omitempty"`
	Workers []*config.Worker `json:"workers,omitempty"`
}

// DefaultRequest provides a default request object by best effort.
func DefaultRequest() Request {
	return Request{
		Name: "",

		Owner: "",

		KubernetesVersion: "",

		AWSSettings: config.ClusterAWSSettings{},

		Masters: []*config.Master{},
		Workers: []*config.Worker{},
	}
}
