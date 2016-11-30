package client

import apischema "github.com/giantswarm/api-schema"

type DeleteClusterRequest struct {
	ID string `json:"id"`
}

type DeleteClusterResponse struct {
	ID string `json:"id"`
}

func (c *Client) DeleteCluster(request DeleteClusterRequest) (DeleteClusterResponse, error) {
	resp, err := apischema.FromHTTPResponse(c.request("DELETE", "/v1/cluster/"+request.ID, nil, nil))
	if err != nil {
		return DeleteClusterResponse{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return DeleteClusterResponse{}, mapErrors(err)
	}

	var response DeleteClusterResponse

	if err := resp.UnmarshalData(&response); err != nil {
		return DeleteClusterResponse{}, maskAny(err)
	}

	return response, nil
}
