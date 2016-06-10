package client

import (
	"time"
)

type ClusterResp struct {
	APIEndpoint     string               `json:"api_endpoint"`
	CreateDate      time.Time            `json:"create_date"`
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	ServiceAccounts []ServiceAccountResp `json:"service_accounts"`
}

type ServiceAccountResp struct {
	ClientCertificateData string `json:"client_certificate_data"`
	ClientKeyData         string `json:"client_key_data"`
	Name                  string `json:"name"`
}

// TODO
func (c *Client) GetClusterByID(clusterID string) (ClusterResp, error) {
	return ClusterResp{}, nil
}
