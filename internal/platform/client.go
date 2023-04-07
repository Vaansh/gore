package platform

import (
	"net/http"
)

type Client struct {
	c   *http.Client
	key string
}

func NewClient() *Client {
	return &Client{
		c: &http.Client{},
	}
}
