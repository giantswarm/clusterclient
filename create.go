package client

import (
	"time"

	"github.com/giantswarm/api-schema"
)

type ClusterReq struct {
	APIEndpoint     string              `json:"api_endpoint"`
	CreateDate      time.Time           `json:"create_date"`
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	ServiceAccounts []ServiceAccountReq `json:"service_accounts"`
}

type ServiceAccountReq struct {
	ClientCertificateData string `json:"client_certificate_data"`
	ClientKeyData         string `json:"client_key_data"`
	Name                  string `json:"name"`
}

func (c *Client) CreateCluster(request ClusterReq) error {
	resp, err := apischema.FromHTTPResponse(c.postJSON("/v1/cluster", request))
	if err != nil {
		return maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_CREATED); err != nil {
		return mapError(err)
	}

	return nil
}
