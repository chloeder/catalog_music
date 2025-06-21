package httpclient

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Client *http.Client
}

func NewHTTPClient() HTTPClient {
	return &Client{
		Client: &http.Client{},
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}
