package shop

import (
	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT             = "/v1/shops"
	LIST_SHOPS_ENDPOINT  = ENDPOINT + ".json"
	DELETE_SHOP_ENDPOINT = ENDPOINT + "/%d.json"
)

var (
	ListShops  = common.ListResources[Shop](LIST_SHOPS_ENDPOINT)
	DeleteShop = common.DeleteResourceWithId[Shop, int](DELETE_SHOP_ENDPOINT)
)
