package product

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

// Client defines product operations and enables dependency injection.
type Client interface {
	ListProducts(id int) ([]Product, error)
	GetProduct(idOne int, idTwo int) (*Product, error)
	CreateProduct(id int, body Product) (*Product, error)
	UpdateProduct(idOne int, idTwo int, body Product) (*Product, error)
	DeleteProduct(idOne int, idTwo int) error
	PublishProduct(idOne int, idTwo int, body Publish) error
	UpdatePublishStatusToSucceeded(idOne int, idTwo int, body PublishReference) error
	UpdatePublishStatusToFailed(idOne int, idTwo int, body PublishFailedRequest) error
	NotifyProductUnpublished(idOne int, idTwo int) error
}

type client struct {
	c *common.Client
}

// NewClient creates a product client implementation backed by common.Client.
func NewClient(c *common.Client) Client {
	return &client{c: c}
}

func (cl *client) ListProducts(id int) ([]Product, error) {
	return ListProducts(cl.c, id)
}

func (cl *client) GetProduct(idOne int, idTwo int) (*Product, error) {
	return GetProduct(cl.c, idOne, idTwo)
}

func (cl *client) CreateProduct(id int, body Product) (*Product, error) {
	return CreateProduct(cl.c, id, body)
}

func (cl *client) UpdateProduct(idOne int, idTwo int, body Product) (*Product, error) {
	return UpdateProduct(cl.c, idOne, idTwo, body)
}

func (cl *client) DeleteProduct(idOne int, idTwo int) error {
	return DeleteProduct(cl.c, idOne, idTwo)
}

func (cl *client) PublishProduct(idOne int, idTwo int, body Publish) error {
	return PublishProduct(cl.c, idOne, idTwo, body)
}

func (cl *client) UpdatePublishStatusToSucceeded(idOne int, idTwo int, body PublishReference) error {
	return UpdatePublishStatusToSucceeded(cl.c, idOne, idTwo, body)
}

func (cl *client) UpdatePublishStatusToFailed(idOne int, idTwo int, body PublishFailedRequest) error {
	return UpdatePublishStatusToFailed(cl.c, idOne, idTwo, body)
}

func (cl *client) NotifyProductUnpublished(idOne int, idTwo int) error {
	return NotifyProductUnpublished(cl.c, idOne, idTwo)
}

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
	// ListProducts calls GET /v1/shops/{shopId}/products.json and returns products in a shop.
	//
	// Signature:
	//	func(c *common.Client, id int) ([]Product, error)
	// Parameter mapping:
	//	id -> {shopId}
	//
	// shopId can be discovered with shop.ListShops.
	ListProducts = common.ListResourceWithId[Product, int](LIST_PRODUCTS_ENDPOINT)
	// GetProduct calls GET /v1/shops/{shopId}/products/{productId}.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*Product, error)
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {productId}
	//
	// shopId can be discovered with shop.ListShops.
	// productId can be discovered with ListProducts(idOne).
	GetProduct = common.GetResourceWithTwoId[Product, int, int](GET_PRODUCT_ENDPOINT)
	// CreateProduct calls POST /v1/shops/{shopId}/products.json.
	//
	// Signature:
	//	func(c *common.Client, id int, body Product) (*Product, error)
	// Parameter mapping:
	//	id -> {shopId}
	//	body -> product payload
	//
	// shopId can be discovered with shop.ListShops.
	CreateProduct = common.PostResourceWithReturnAndId[Product, Product, int](CREATE_PRODUCT_ENDPOINT)
	// UpdateProduct calls PUT /v1/shops/{shopId}/products/{productId}.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int, body Product) (*Product, error)
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {productId}
	//	body -> product payload
	//
	// shopId can be discovered with shop.ListShops.
	// productId can be discovered with ListProducts(idOne).
	UpdateProduct = common.PutResourceWithReturnAndTwoId[Product, Product, int, int](UPDATE_PRODUCT_ENDPOINT)
	// DeleteProduct calls DELETE /v1/shops/{shopId}/products/{productId}.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) error
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {productId}
	//
	// shopId can be discovered with shop.ListShops.
	// productId can be discovered with ListProducts(idOne).
	DeleteProduct = common.DeleteResourceWithTwoId[Product, int, int](DELETE_PRODUCT_ENDPOINT)
	// PublishProduct calls POST /v1/shops/{shopId}/products/{productId}/publish.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int, body Publish) error
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {productId}
	//	body -> publish options
	//
	// shopId can be discovered with shop.ListShops.
	// productId can be discovered with ListProducts(idOne).
	PublishProduct = common.PostResourceWithoutReturnTwoId[Publish, int, int](PUBLISH_PRODUCT_ENDPOINT)
	// UpdatePublishStatusToSucceeded calls
	// POST /v1/shops/{shopId}/products/{productId}/publishing_succeeded.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int, body PublishReference) error
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {productId}
	//	body -> external channel references
	//
	// shopId can be discovered with shop.ListShops.
	// productId can be discovered with ListProducts(idOne).
	UpdatePublishStatusToSucceeded = common.PostResourceWithoutReturnTwoId[PublishReference, int, int](UPDATE_PUBLISH_STATUS_TO_SUCCEEDED_ENDPOINT)
	// UpdatePublishStatusToFailed calls
	// POST /v1/shops/{shopId}/products/{productId}/publishing_failed.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int, body PublishFailedRequest) error
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {productId}
	//	body -> failure reason payload
	//
	// shopId can be discovered with shop.ListShops.
	// productId can be discovered with ListProducts(idOne).
	UpdatePublishStatusToFailed = common.PostResourceWithoutReturnTwoId[PublishFailedRequest, int, int](UPDATE_PUBLISH_STATUS_TO_FAILED_ENDPOINT)
	// NotifyProductUnpublished calls POST /v1/shops/{shopId}/products/{productId}/unpublished.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) error
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {productId}
	//
	// shopId can be discovered with shop.ListShops.
	// productId can be discovered with ListProducts(idOne).
	NotifyProductUnpublished = common.PostNoResourceWithoutReturnTwoId[int, int](NOTIFY_PRODUCT_UNPUBLISHED_ENDPOINT)
)
