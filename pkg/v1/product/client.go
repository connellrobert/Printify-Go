package product

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                                    = "/v1/shops"
	LIST_PRODUCTS_ENDPOINT                      = fmt.Sprintf("%s/%%d/products.json", ENDPOINT)
	GET_PRODUCT_ENDPOINT                        = fmt.Sprintf("%s/%%d/products/%%d.json", ENDPOINT)
	CREATE_PRODUCT_ENDPOINT                     = fmt.Sprintf("%s/%%d/products.json", ENDPOINT)
	UPDATE_PRODUCT_ENDPOINT                     = fmt.Sprintf("%s/%%d/products/%%d.json", ENDPOINT)
	DELETE_PRODUCT_ENDPOINT                     = fmt.Sprintf("%s/%%d/products/%%d.json", ENDPOINT)
	PUBLISH_PRODUCT_ENDPOINT                    = fmt.Sprintf("%s/%%d/products/%%d/publish.json", ENDPOINT)
	UPDATE_PUBLISH_STATUS_TO_SUCCEEDED_ENDPOINT = fmt.Sprintf("%s/%%d/products/%%d/publishing_succeeded.json", ENDPOINT)
	UPDATE_PUBLISH_STATUS_TO_FAILED_ENDPOINT    = fmt.Sprintf("%s/%%d/products/%%d/publishing_failed.json", ENDPOINT)
	NOTIFY_PRODUCT_UNPUBLISHED_ENDPOINT         = fmt.Sprintf("%s/%%d/products/%%d/unpublished.json", ENDPOINT)
)

var (
	ListProducts                   = common.ListResourceWithId[Product, int](LIST_PRODUCTS_ENDPOINT)
	GetProduct                     = common.GetResourceWithTwoId[Product, int, int](GET_PRODUCT_ENDPOINT)
	CreateProduct                  = common.PostResourceWithReturnAndId[Product, Product, int](CREATE_PRODUCT_ENDPOINT)
	UpdateProduct                  = common.PutResourceWithReturnAndTwoId[Product, Product, int, int](UPDATE_PRODUCT_ENDPOINT)
	DeleteProduct                  = common.DeleteResourceWithTwoId[Product, int, int](DELETE_PRODUCT_ENDPOINT)
	PublishProduct                 = common.PostResourceWithoutReturnTwoId[Publish, int, int](PUBLISH_PRODUCT_ENDPOINT)
	UpdatePublishStatusToSucceeded = common.PostResourceWithoutReturnTwoId[PublishReference, int, int](UPDATE_PUBLISH_STATUS_TO_SUCCEEDED_ENDPOINT)
	UpdatePublishStatusToFailed    = common.PostResourceWithoutReturnTwoId[PublishFailedRequest, int, int](UPDATE_PUBLISH_STATUS_TO_FAILED_ENDPOINT)
	NotifyProductUnpublished       = common.PostNoResourceWithoutReturnTwoId[int, int](NOTIFY_PRODUCT_UNPUBLISHED_ENDPOINT)
)
