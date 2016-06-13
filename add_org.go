package client

import (
	"fmt"

	"github.com/giantswarm/api-schema"
)

type AddOwnersReq struct {
	Owners []OwnerReq `json:"owners"`
}

type OwnerReq struct {
	ID        string `json:"id"`
	OwnerType string `json:"owner_type"`
}

func (c *Client) AddOwnersToCluster(clusterID string, request AddOwnersReq) error {
	resp, err := apischema.FromHTTPResponse(c.postJSON(fmt.Sprintf("/v1/cluster/%s/owners", clusterID), request))
	if err != nil {
		return maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_RESOURCE_CREATED); err != nil {
		// TODO
		return maskAny(err)
	}

	return nil
}
