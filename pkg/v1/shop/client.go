package shop

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

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
