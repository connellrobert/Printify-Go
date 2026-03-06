package shop

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

// Client defines shop operations and enables dependency injection.
type Client interface {
	ListShops() ([]Shop, error)
	DeleteShop(id int) error
}

type client struct {
	c *common.Client
}

// NewClient creates a shop client implementation backed by common.Client.
func NewClient(c *common.Client) Client {
	return &client{c: c}
}

func (cl *client) ListShops() ([]Shop, error) {
	return ListShops(cl.c)
}

func (cl *client) DeleteShop(id int) error {
	return DeleteShop(cl.c, id)
}

var (
	ENDPOINT             = "/v1/shops"
	LIST_SHOPS_ENDPOINT  = fmt.Sprintf("%s.json", ENDPOINT)
	DELETE_SHOP_ENDPOINT = fmt.Sprintf("%s/%%d.json", ENDPOINT)
)

var (
	// ListShops calls GET /v1/shops.json and returns all shops accessible by the PAT.
	//
	// Signature:
	//	func(c *common.Client) ([]Shop, error)
	ListShops = common.ListResources[Shop](LIST_SHOPS_ENDPOINT)
	// DeleteShop calls DELETE /v1/shops/{shopId}.json.
	//
	// Signature:
	//	func(c *common.Client, id int) error
	// Parameter mapping:
	//	id -> {shopId}
	//
	// shopId can be discovered with ListShops.
	DeleteShop = common.DeleteResourceWithId[Shop, int](DELETE_SHOP_ENDPOINT)
)
