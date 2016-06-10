package client

type AddOwnersReq struct {
	Owners []OwnerReq `json:"owners"`
}

type OwnerReq struct {
	ID        string `json:"id"`
	OwnerType string `json:"owner_type"`
}

// TODO
func (c *Client) AddOwnersToCluster(clusterID string, request AddOwnersReq) error {
	return nil
}
