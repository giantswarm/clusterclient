package client

import (
	"time"
)

type ClusterListResp struct {
	Clusters []ClusterItemResp `json:"clusters"`
}

type ClusterItemResp struct {
	CreateDate time.Time `json:"create_date"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
}

// TODO
func (c *Client) ListClustersForOrg(orgID string) (ClusterListResp, error) {
	return ClusterListResp{}, nil
}
