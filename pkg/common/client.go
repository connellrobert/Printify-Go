package common

import (
	"net/http"
)

const (
	HOST = "https://api.printify.com"
)

type Client struct {
	Host   string
	Client *http.Client
	PAT    string
	ShopID string
}

func NewClient(pat string, shopId int) *Client {
	return &Client{
		Host:   HOST,
		Client: &http.Client{},
		PAT:    pat,
	}
}
