package searcher

import (
	"time"

	"github.com/giantswarm/clusterclient/service/cluster/searcher/config"
)

// Response is the return value of the service action.
type Response struct {
	APIEndpoint string    `json:"api_endpoint"`
	CreateDate  time.Time `json:"create_date"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`

	Owner string `json:"owner"`

	KubernetesVersion string `json:"kubernetes_version"`

	AWSSettings config.ClusterAWSSettings `json:"aws,omitempty"`

	Masters []*config.Master `json:"masters"`
	Workers []*config.Worker `json:"workers"`
}

// DefaultResponse provides a default response object by best effort.
func DefaultResponse() *Response {
	return &Response{
		APIEndpoint: "",
		CreateDate:  time.Time{},
		ID:          "",
		Name:        "",

		Owner: "",

		KubernetesVersion: "",

		AWSSettings: config.ClusterAWSSettings{},

		Masters: []*config.Master{},
		Workers: []*config.Worker{},
	}
}
