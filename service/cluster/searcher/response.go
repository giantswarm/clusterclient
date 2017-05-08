package searcher

import (
	"time"

	"github.com/giantswarm/clusterclient/service/cluster/searcher/response"
	"github.com/giantswarm/clusterclient/service/cluster/searcher/response/aws"
)

// Response is the return value of the service action.
type Response struct {
	APIEndpoint string    `json:"api_endpoint"`
	CreateDate  time.Time `json:"create_date"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`

	Owner string `json:"owner"`

	KubernetesVersion string `json:"kubernetes_version"`

	AWS aws.Cluster `json:"aws,omitempty"`

	Masters []response.Master `json:"masters"`
	Workers []response.Worker `json:"workers"`
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

		AWS: aws.DefaultCluster(),

		Masters: []response.Master{},
		Workers: []response.Worker{},
	}
}
