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
	ListShops  = common.ListResources[Shop](LIST_SHOPS_ENDPOINT)
	DeleteShop = common.DeleteResourceWithId[Shop, int](DELETE_SHOP_ENDPOINT)
)
