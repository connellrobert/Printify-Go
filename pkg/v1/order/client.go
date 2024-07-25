package order

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

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
	ListOrders                 = common.ListResources[Order](LIST_ORDERS_ENDPOINT)
	GetOrderDetails            = common.GetResourceWithTwoId[Order, int, int](GET_ORDER_DETAILS_ENDPOINT)
	SubmitOrder                = common.PostResourceWithReturnTwoId[Order, Order, int, int](SUBMIT_ORDER_ENDPOINT)
	SubmitPrintifyExpressOrder = common.PostResourceWithReturnTwoId[Order, Order, int, int](SUBMIT_PRINTIFY_EXPRESS_ORDER_ENDPOINT)
	SendOrderToProduction      = common.PostResourceWithoutReturnTwoId[Order, int, int](SEND_ORDER_TO_PRODUCTION_ENDPOINT)
	CalculateShippingCosts     = common.PostResourceWithReturnAndId[ShipmentCalculationRequest, ShipmentCalculationResponse, int](CALCULATE_SHIPPING_COSTS_ENDPOINT)
	CancelOrder                = common.PostNoResourceWithReturnTwoId[Order, int, int](CANCEL_ORDER_ENDPOINT)
)
