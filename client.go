package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/juju/errgo"
)

type Config struct {
	User     string
	Password string
	Endpoint string
	Insecure bool
}

type Client struct {
	Config
}

func New(config Config) (*Client, error) {
	c := &Client{
		Config: config,
	}

	_, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, errgo.NoteMask(err, fmt.Sprintf("parsing endpoint failed", c.Endpoint), errgo.Any)
	}

	return c, nil
}

func (c *Client) get(path string) (*http.Response, error) {
	resp, err := c.request("GET", path, nil, nil)
	if err != nil {
		return nil, errgo.NoteMask(err, fmt.Sprintf("GET '%s' failed", path), errgo.Any)
	}

	return resp, nil
}

func (c *Client) postJSON(path string, data interface{}) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errgo.NoteMask(err, fmt.Sprintf("marshaling %#v failed", data), errgo.Any)
	}
	r := bytes.NewReader(b)

	resp, err := c.request("POST", path, headers, r)
	if err != nil {
		return nil, errgo.NoteMask(err, fmt.Sprintf("POST json '%s' failed", path), errgo.Any)
	}

	return resp, nil
}

func (c *Client) request(method, path string, headers map[string]string, body io.Reader) (*http.Response, error) {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, errgo.NoteMask(err, fmt.Sprintf("parsing endpoint failed", c.Endpoint), errgo.Any)
	}
	u.Path = path

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, errgo.NoteMask(err, fmt.Sprintf("request to '%s' failed", u), errgo.Any)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if c.User != "" || c.Password != "" {
		req.SetBasicAuth(c.User, c.Password)
	}

	var client *http.Client
	if c.Insecure {
		client = &http.Client{}
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errgo.NoteMask(err, fmt.Sprintf("request to '%s' failed", u), errgo.Any)
	}

	return res, nil
}
