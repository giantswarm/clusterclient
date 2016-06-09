package client

import (
	"time"
)

type GetClusterByIDResp struct {
	APIEndpoint     string               `json:"api_endpoint"`
	CreateDate      time.Time            `json:"create_date"`
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	ServiceAccounts []ServiceAccountResp `json:"service_accounts"` // TODO make PR in mock-api
}

type ServiceAccountResp struct {
	ClientCertificateData string `json:"client_certificate_data"`
	ClientKeyData         string `json:"client_key_data"`
	Name                  string `json:"name"`
}

// TODO
func (c *Client) GetClusterByID(clusterID string) (GetClusterByIDResp, error) {
	return GetClusterByIDResp{}, nil
}
