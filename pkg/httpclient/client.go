package httpclient

import "net/http"

//go:generate mockgen -source=client.go -destination=client_mock_test.go -package=httpclient
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
