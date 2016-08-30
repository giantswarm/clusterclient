package client

import (
	"fmt"
	"time"

	"github.com/giantswarm/api-schema"
)

type ClusterResp struct {
	APIEndpoint string    `json:"api_endpoint"`
	CreateDate  time.Time `json:"create_date"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
}

func (c *Client) GetClusterByID(clusterID string) (ClusterResp, error) {
	resp, err := apischema.FromHTTPResponse(c.get(fmt.Sprintf("/v1/cluster/%s", clusterID)))
	if err != nil {
		return ClusterResp{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return ClusterResp{}, mapError(err)
	}

	var response ClusterResp

	if err := resp.UnmarshalData(&response); err != nil {
		return ClusterResp{}, maskAny(err)
	}

	return response, nil
}
