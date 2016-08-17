package client

import (
	"fmt"

	"github.com/giantswarm/api-schema"
)

type CreateKeyPairRequest struct {
	Description string `json:"description"`
}

type CreateKeyPairResponse struct {
	CertificateAuthorityData string `json:"certificate_authority_data"`
	ClientKeyData            string `json:"client_key_data"`
	ClientCertificateData    string `json:"client_certificate_data"`
	SerialNumber             string `json:"serial_number"`
}

func (c *Client) CreateClusterKeyPair(clusterID string, request CreateKeyPairRequest) (CreateKeyPairResponse, error) {
	resp, err := apischema.FromHTTPResponse(c.postJSON(fmt.Sprintf("/v1/cluster/%s/key-pair", clusterID), request))
	if err != nil {
		return CreateKeyPairResponse{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return CreateKeyPairResponse{}, mapError(maskAny(err))
	}

	var response CreateKeyPairResponse

	if err := resp.UnmarshalData(&response); err != nil {
		return CreateKeyPairResponse{}, maskAny(err)
	}

	return response, nil
}
