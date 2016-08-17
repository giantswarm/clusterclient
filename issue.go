package client

import (
	"fmt"

	"github.com/giantswarm/api-schema"
)

type KeyPairResp struct {
	CertificateAuthorityData string `json:"certificate_authority_data"`
	ClientKeyData            string `json:"client_key_data"`
	ClientCertificateData    string `json:"client_certificate_data"`
}

func (c *Client) GetClusterKeyPair(clusterID, keyName string) (KeyPairResp, error) {
	resp, err := apischema.FromHTTPResponse(c.get(fmt.Sprintf("/v1/cluster/%s/key-pair/%s", clusterID, keyName)))
	if err != nil {
		return KeyPairResp{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return KeyPairResp{}, mapError(err)
	}

	var response KeyPairResp

	if err := resp.UnmarshalData(&response); err != nil {
		return KeyPairResp{}, maskAny(err)
	}

	return response, nil
}
