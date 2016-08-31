package client

import (
	"fmt"
	"time"

	"github.com/giantswarm/api-schema"
)

type ListKeyPairsResponse struct {
	KeyPairs []KeyPairResponse `json:key_pairs`
}

type KeyPairResponse struct {
	// CreateDate represents the timestamp of the serial items creation.
	CreateDate time.Time `json:"create_date"`

	// Description represents the description associated with an issued key-pair.
	Description string `json:"description"`

	// SerialNumber represents the fingerprint of an issued key-pair.
	SerialNumber string `json:"serial_number"`

	// TTL represents the time the number of hours the key-pair associated with
	// the current serial item is valid.
	TTL int `json:"ttl"`
}

func (c *Client) ListKeyPairs(clusterID string) (ListKeyPairsResponse, error) {
	resp, err := apischema.FromHTTPResponse(c.get(fmt.Sprintf("/v1/cluster/%s/list/key-pair", clusterID)))
	if err != nil {
		return ListKeyPairsResponse{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return ListKeyPairsResponse{}, mapErrors(err)
	}

	var response ListKeyPairsResponse

	if err := resp.UnmarshalData(&response); err != nil {
		return ListKeyPairsResponse{}, maskAny(err)
	}

	return response, nil
}
