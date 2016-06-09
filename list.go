package client

import (
	"time"
)

type ListClustersForOrgResp struct {
	Clusters []ListClusterReq `json:"clusters"`
}

type ListClusterReq struct {
	CreateDate time.Time `json:"create_date"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
}

// TODO
func (c *Client) ListClustersForOrg(orgName string) (ListClustersForOrgResp, error) {
	return ListClustersForOrgResp{}, nil
}
