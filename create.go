package client

import (
	"time"
)

type V3CreateClusterReq struct {
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

// TODO
func (c *Client) V3CreateCluster(request V3CreateClusterReq) error {
	return nil
}
