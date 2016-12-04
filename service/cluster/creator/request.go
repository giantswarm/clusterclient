package creator

import (
	"time"

	"github.com/giantswarm/clusterclient/service/cluster/creator/config"
)

type Request struct {
	APIEndpoint string    `json:"api_endpoint"`
	CreateDate  time.Time `json:"create_date"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`

	Owner string `json:"owner"`

	KubernetesVersion string `json:"kubernetes_version"`

	Masters []*config.Master `json:"masters"`
	Workers []*config.Worker `json:"workers"`
}

// DefaultRequest provides a default request object by best effort.
func DefaultRequest() Request {
	return Request{
		APIEndpoint: "",
		CreateDate:  time.Time{},
		ID:          "",
		Name:        "",

		Owner: "",

		KubernetesVersion: "",

		Masters: []*config.Master{},
		Workers: []*config.Worker{},
	}
}
