package platform

import (
	"net/http"
)

type Client struct {
	c      *http.Client
	apiKey string
}

func NewClient() *Client {
	return &Client{
		c: &http.Client{},
	}
}
