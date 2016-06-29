package client

import (
	"fmt"
	"time"

	"github.com/giantswarm/api-schema"
)

// response

type KeyPairResponse struct {
	Crt string `json:"crt"`
	Key string `json:"key"`
}

// ca

type CAKeyPairRequest struct {
	// AllowedDomains represents a comma separate list of valid domain names the
	// generated certificate authority is valid for.
	AllowedDomains string `json:"allowed_domains"`

	CommonName string `json:"common_name"`

	// TTL configures the time to live for the requested certificate. This is a
	// golang time string with the allowed units s, m and h.
	TTL string `json:"ttl"`
}

func (c *Client) GetOrCreateCAKeyPair(clusterID string, request CAKeyPairRequest) (KeyPairResponse, error) {
	resp, err := apischema.FromHTTPResponse(c.postJSON(fmt.Sprintf("/v1/cluster/%s/certificate/key-pair/ca", clusterID), request))
	if err != nil {
		return KeyPairResponse{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return KeyPairResponse{}, mapError(err)
	}

	var response KeyPairResp

	if err := resp.UnmarshalData(&response); err != nil {
		return KeyPairResponse{}, maskAny(err)
	}

	return response, nil
}

// signed

type SignedKeyPairRequest struct {
	CommonName string `json:"common_name"`

	// IPSANs represents a comma separate lists of IPs.
	IPSANs string `json:"ip_sans"`

	// TTL configures the time to live for the requested certificate. This is a
	// golang time string with the allowed units s, m and h.
	TTL string `json:"ttl"`
}

func (c *Client) GetOrCreateSignedKeyPair(clusterID, certName string, request SignedKeyPairRequest) (KeyPairResponse, error) {
	resp, err := apischema.FromHTTPResponse(c.postJSON(fmt.Sprintf("/v1/cluster/%s/certificate/key-pair/%s", clusterID, certName), request))
	if err != nil {
		return KeyPairResponse{}, maskAny(err)
	}

	if err := resp.EnsureStatusCodes(apischema.STATUS_CODE_DATA); err != nil {
		return KeyPairResponse{}, mapError(err)
	}

	var response KeyPairResp

	if err := resp.UnmarshalData(&response); err != nil {
		return KeyPairResponse{}, maskAny(err)
	}

	return response, nil
}
