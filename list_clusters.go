package client

import (
	"fmt"
	"time"

	"github.com/giantswarm/api-schema"
)

type ClusterListResp struct {
	Clusters []ClusterItemResp `json:"clusters"`
}

type ClusterItemResp struct {
	CreateDate time.Time `json:"create_date"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
}

func (c *Client) ListClustersForOrg(orgID string) (ClusterListResp, error) {
	resp, err := apischema.FromHTTPResponse(c.get(fmt.Sprintf("/v1/cluster/list/%s", orgID)))
	if err != nil {
		return ClusterListResp{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return ClusterListResp{}, mapErrors(err)
	}

	var response ClusterListResp

	if err := resp.UnmarshalData(&response); err != nil {
		return ClusterListResp{}, maskAny(err)
	}

	return response, nil
}
