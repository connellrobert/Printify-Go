package order

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

// Client defines order operations and enables dependency injection.
type Client interface {
	ListOrders() ([]Order, error)
	GetOrderDetails(idOne int, idTwo int) (*Order, error)
	SubmitOrder(idOne int, idTwo int, body Order) (*Order, error)
	SubmitPrintifyExpressOrder(idOne int, idTwo int, body Order) (*Order, error)
	SendOrderToProduction(idOne int, idTwo int, body Order) error
	CalculateShippingCosts(id int, body ShipmentCalculationRequest) (*ShipmentCalculationResponse, error)
	CancelOrder(idOne int, idTwo int) (*Order, error)
}

type client struct {
	c *common.Client
}

// NewClient creates an order client implementation backed by common.Client.
func NewClient(c *common.Client) Client {
	return &client{c: c}
}

func (cl *client) ListOrders() ([]Order, error) {
	return ListOrders(cl.c)
}

func (cl *client) GetOrderDetails(idOne int, idTwo int) (*Order, error) {
	return GetOrderDetails(cl.c, idOne, idTwo)
}

func (cl *client) SubmitOrder(idOne int, idTwo int, body Order) (*Order, error) {
	return SubmitOrder(cl.c, idOne, idTwo, body)
}

func (cl *client) SubmitPrintifyExpressOrder(idOne int, idTwo int, body Order) (*Order, error) {
	return SubmitPrintifyExpressOrder(cl.c, idOne, idTwo, body)
}

func (cl *client) SendOrderToProduction(idOne int, idTwo int, body Order) error {
	return SendOrderToProduction(cl.c, idOne, idTwo, body)
}

func (cl *client) CalculateShippingCosts(id int, body ShipmentCalculationRequest) (*ShipmentCalculationResponse, error) {
	return CalculateShippingCosts(cl.c, id, body)
}

func (cl *client) CancelOrder(idOne int, idTwo int) (*Order, error) {
	return CancelOrder(cl.c, idOne, idTwo)
}

var (
	ENDPOINT                               = "/v1/shops"
	LIST_ORDERS_ENDPOINT                   = fmt.Sprintf("%s/%%d/orders.json", ENDPOINT)
	GET_ORDER_DETAILS_ENDPOINT             = fmt.Sprintf("%s/%%d/orders/%%d.json", ENDPOINT)
	SUBMIT_ORDER_ENDPOINT                  = fmt.Sprintf("%s/%%d/orders.json", ENDPOINT)
	SUBMIT_PRINTIFY_EXPRESS_ORDER_ENDPOINT = fmt.Sprintf("%s/%%d/orders/express.json", ENDPOINT)
	SEND_ORDER_TO_PRODUCTION_ENDPOINT      = fmt.Sprintf("%s/%%d/orders/%%d/send_to_production.json", ENDPOINT)
	CALCULATE_SHIPPING_COSTS_ENDPOINT      = fmt.Sprintf("%s/%%d/orders/shipping.json", ENDPOINT)
	CANCEL_ORDER_ENDPOINT                  = fmt.Sprintf("%s/%%d/orders/%%d/cancel.json", ENDPOINT)
)

var (
	// ListOrders calls GET /v1/shops/{shopId}/orders.json and returns orders for a shop.
	//
	// Signature:
	//	func(c *common.Client) ([]Order, error)
	//
	// The shop id used by this endpoint is taken from the client that created the request.
	// Create the client with the desired shop id, or discover shops with shop.ListShops.
	ListOrders = func(c *common.Client) ([]Order, error) {
		return common.ListResourceWithId[Order, int](LIST_ORDERS_ENDPOINT)(c, c.ShopID)
	}
	// GetOrderDetails calls GET /v1/shops/{shopId}/orders/{orderId}.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*Order, error)
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {orderId}
	//
	// shopId can be discovered with shop.ListShops.
	// orderId can be discovered with ListOrders for the same shop.
	GetOrderDetails = common.GetResourceWithTwoId[Order, int, int](GET_ORDER_DETAILS_ENDPOINT)
	// SubmitOrder calls POST /v1/shops/{shopId}/orders.json to create an order.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int, body Order) (*Order, error)
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> ignored by this endpoint
	//	body -> order payload
	//
	// shopId can be discovered with shop.ListShops.
	// idTwo is unused because this endpoint has only one path identifier.
	SubmitOrder = func(c *common.Client, idOne int, idTwo int, body Order) (*Order, error) {
		_ = idTwo
		return common.PostResourceWithReturnAndId[Order, Order, int](SUBMIT_ORDER_ENDPOINT)(c, idOne, body)
	}
	// SubmitPrintifyExpressOrder calls POST /v1/shops/{shopId}/orders/express.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int, body Order) (*Order, error)
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> ignored by this endpoint
	//	body -> order payload
	//
	// shopId can be discovered with shop.ListShops.
	// idTwo is unused because this endpoint has only one path identifier.
	SubmitPrintifyExpressOrder = func(c *common.Client, idOne int, idTwo int, body Order) (*Order, error) {
		_ = idTwo
		return common.PostResourceWithReturnAndId[Order, Order, int](SUBMIT_PRINTIFY_EXPRESS_ORDER_ENDPOINT)(c, idOne, body)
	}
	// SendOrderToProduction calls POST /v1/shops/{shopId}/orders/{orderId}/send_to_production.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int, body Order) error
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {orderId}
	//	body -> optional request payload (usually an empty Order)
	//
	// shopId can be discovered with shop.ListShops.
	// orderId can be discovered with ListOrders for the same shop.
	SendOrderToProduction = common.PostResourceWithoutReturnTwoId[Order, int, int](SEND_ORDER_TO_PRODUCTION_ENDPOINT)
	// CalculateShippingCosts calls POST /v1/shops/{shopId}/orders/shipping.json.
	//
	// Signature:
	//	func(c *common.Client, id int, body ShipmentCalculationRequest) (*ShipmentCalculationResponse, error)
	// Parameter mapping:
	//	id -> {shopId}
	//	body -> shipping calculation payload
	//
	// shopId can be discovered with shop.ListShops.
	CalculateShippingCosts = common.PostResourceWithReturnAndId[ShipmentCalculationRequest, ShipmentCalculationResponse, int](CALCULATE_SHIPPING_COSTS_ENDPOINT)
	// CancelOrder calls POST /v1/shops/{shopId}/orders/{orderId}/cancel.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo int) (*Order, error)
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {orderId}
	//
	// shopId can be discovered with shop.ListShops.
	// orderId can be discovered with ListOrders for the same shop.
	CancelOrder = common.PostNoResourceWithReturnTwoId[Order, int, int](CANCEL_ORDER_ENDPOINT)
)
