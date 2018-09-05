package searcher

import (
	"time"

	"github.com/giantswarm/versionbundle"

	"github.com/giantswarm/clusterclient/service/cluster/searcher/response"
	"github.com/giantswarm/clusterclient/service/cluster/searcher/response/aws"
	"github.com/giantswarm/clusterclient/service/cluster/searcher/response/kvm"
)

// Response is the return value of the service action.
type Response struct {
	APIEndpoint    string                 `json:"api_endpoint"`
	AWS            aws.Cluster            `json:"aws,omitempty"`
	CreateDate     time.Time              `json:"create_date"`
	ID             string                 `json:"id"`
	KVM            kvm.Cluster            `json:"kvm,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Owner          string                 `json:"owner,omitempty"`
	ReleaseVersion string                 `json:"release_version,omitempty"`
	VersionBundles []versionbundle.Bundle `json:"version_bundles,omitempty"`
	Workers        []response.Worker      `json:"workers,omitempty"`
}

// DefaultResponse provides a default response by best effort.
func DefaultResponse() *Response {
	return &Response{
		APIEndpoint:    "",
		AWS:            aws.DefaultCluster(),
		CreateDate:     time.Time{},
		ID:             "",
		KVM:            kvm.Cluster{},
		Name:           "",
		Owner:          "",
		ReleaseVersion: "",
		VersionBundles: []versionbundle.Bundle{},
		Workers:        []response.Worker{},
	}
}
