package client

import (
	"time"

	"github.com/giantswarm/api-schema"
)

type ClusterReq struct {
	APIEndpoint              string              `json:"api_endpoint"`
	CertificateAuthorityData string              `json:"certificate_authority_data"`
	CreateDate               time.Time           `json:"create_date"`
	ID                       string              `json:"id"`
	Name                     string              `json:"name"`
	ServiceAccounts          []ServiceAccountReq `json:"service_accounts"`
}

type ServiceAccountReq struct {
	ClientCertificateData string `json:"client_certificate_data"`
	ClientKeyData         string `json:"client_key_data"`
	Name                  string `json:"name"`
}

type CreateClusterResp struct {
	ID string `json:"id"`
}

func (c *Client) CreateCluster(request ClusterReq) (CreateClusterResp, error) {
	resp, err := apischema.FromHTTPResponse(c.postJSON("/v1/cluster", request))
	if err != nil {
		return CreateClusterResp{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return CreateClusterResp{}, mapError(err)
	}

	var createResp CreateClusterResp

	if err := resp.UnmarshalData(&createResp); err != nil {
		return CreateClusterResp{}, maskAny(err)
	}

	return createResp, nil
}
