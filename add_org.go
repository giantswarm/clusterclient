package client

type AddOrgsToClusterReq struct {
	Owners []OwnerReq `json:"owners"`
}

type OwnerReq struct {
	ID        string `json:"id"`
	OwnerType string `json:"owner_type"`
}

// TODO
func (c *Client) AddOrgsToCluster(clusterID string, request AddOrgsToClusterReq) error {
	return nil
}
