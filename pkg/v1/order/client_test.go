package order

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/connellrobert/printify-go/pkg/common"
)

func newOrderTestClient(register func(mux *http.ServeMux)) (*common.Client, func()) {
	mux := http.NewServeMux()
	register(mux)
	srv := httptest.NewServer(mux)
	c := common.NewClient("printify_pat", 123)
	c.Host = srv.URL
	return c, srv.Close
}

func ExampleNewClient() {
	c := common.NewClient("printify_pat", 123)
	svc := NewClient(c)
	fmt.Println(svc != nil)
	// Output: true
}

func ExampleListOrders() {
	c, closeFn := newOrderTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/orders.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"id":"ord_1"}]}`))
		})
	})
	defer closeFn()

	items, _ := ListOrders(c)
	fmt.Printf("%#v\n", items[0])
	// Output: order.Order{Id:"ord_1", AddressTo:order.Address{FirstName:"", LastName:"", Region:"", Address1:"", City:"", Zip:"", Email:"", Phone:"", Country:"", Company:""}, LineItems:[]order.LineItem(nil), Metadata:order.OrderMetadata{OrderType:"", ShopOrderId:0, ShopOrderLabel:"", ShopFulfilledAt:""}, TotalPrice:0, TotalShipping:0, TotalTax:0, Status:"", ShippingMethod:0, IsPrintifyExpress:false, IsEconomyShipping:false, Shipments:[]order.Shipment(nil), CreatedAt:"", SentToProductionAt:"", FulfilledAt:"", PrintifyConnect:order.PrintifyConnect{Url:"", Id:""}}
}

func ExampleGetOrderDetails() {
	c, closeFn := newOrderTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/orders/456.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"ord_456"}`))
		})
	})
	defer closeFn()

	item, _ := GetOrderDetails(c, 123, 456)
	fmt.Printf("%#v\n", *item)
	// Output: order.Order{Id:"ord_456", AddressTo:order.Address{FirstName:"", LastName:"", Region:"", Address1:"", City:"", Zip:"", Email:"", Phone:"", Country:"", Company:""}, LineItems:[]order.LineItem(nil), Metadata:order.OrderMetadata{OrderType:"", ShopOrderId:0, ShopOrderLabel:"", ShopFulfilledAt:""}, TotalPrice:0, TotalShipping:0, TotalTax:0, Status:"", ShippingMethod:0, IsPrintifyExpress:false, IsEconomyShipping:false, Shipments:[]order.Shipment(nil), CreatedAt:"", SentToProductionAt:"", FulfilledAt:"", PrintifyConnect:order.PrintifyConnect{Url:"", Id:""}}
}

func ExampleSubmitOrder() {
	c, closeFn := newOrderTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/orders.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"ord_submit"}`))
		})
	})
	defer closeFn()

	item, _ := SubmitOrder(c, 123, 0, Order{})
	fmt.Printf("%#v\n", *item)
	// Output: order.Order{Id:"ord_submit", AddressTo:order.Address{FirstName:"", LastName:"", Region:"", Address1:"", City:"", Zip:"", Email:"", Phone:"", Country:"", Company:""}, LineItems:[]order.LineItem(nil), Metadata:order.OrderMetadata{OrderType:"", ShopOrderId:0, ShopOrderLabel:"", ShopFulfilledAt:""}, TotalPrice:0, TotalShipping:0, TotalTax:0, Status:"", ShippingMethod:0, IsPrintifyExpress:false, IsEconomyShipping:false, Shipments:[]order.Shipment(nil), CreatedAt:"", SentToProductionAt:"", FulfilledAt:"", PrintifyConnect:order.PrintifyConnect{Url:"", Id:""}}
}

func ExampleSubmitPrintifyExpressOrder() {
	c, closeFn := newOrderTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/orders/express.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"ord_express"}`))
		})
	})
	defer closeFn()

	item, _ := SubmitPrintifyExpressOrder(c, 123, 0, Order{})
	fmt.Printf("%#v\n", *item)
	// Output: order.Order{Id:"ord_express", AddressTo:order.Address{FirstName:"", LastName:"", Region:"", Address1:"", City:"", Zip:"", Email:"", Phone:"", Country:"", Company:""}, LineItems:[]order.LineItem(nil), Metadata:order.OrderMetadata{OrderType:"", ShopOrderId:0, ShopOrderLabel:"", ShopFulfilledAt:""}, TotalPrice:0, TotalShipping:0, TotalTax:0, Status:"", ShippingMethod:0, IsPrintifyExpress:false, IsEconomyShipping:false, Shipments:[]order.Shipment(nil), CreatedAt:"", SentToProductionAt:"", FulfilledAt:"", PrintifyConnect:order.PrintifyConnect{Url:"", Id:""}}
}

func ExampleSendOrderToProduction() {
	c, closeFn := newOrderTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/orders/456/send_to_production.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := SendOrderToProduction(c, 123, 456, Order{})
	fmt.Println(err == nil)
	// Output: true
}

func ExampleCalculateShippingCosts() {
	c, closeFn := newOrderTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/orders/shipping.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"standard":500,"express":900,"priority":700,"printify_express":1100,"economy":300}`))
		})
	})
	defer closeFn()

	item, _ := CalculateShippingCosts(c, 123, ShipmentCalculationRequest{})
	fmt.Printf("%#v\n", *item)
	// Output: order.ShipmentCalculationResponse{Standard:500, Express:900, Priority:700, PrintifyExpress:1100, Economy:300}
}

func ExampleCancelOrder() {
	c, closeFn := newOrderTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/orders/456/cancel.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"ord_cancel"}`))
		})
	})
	defer closeFn()

	item, _ := CancelOrder(c, 123, 456)
	fmt.Printf("%#v\n", *item)
	// Output: order.Order{Id:"ord_cancel", AddressTo:order.Address{FirstName:"", LastName:"", Region:"", Address1:"", City:"", Zip:"", Email:"", Phone:"", Country:"", Company:""}, LineItems:[]order.LineItem(nil), Metadata:order.OrderMetadata{OrderType:"", ShopOrderId:0, ShopOrderLabel:"", ShopFulfilledAt:""}, TotalPrice:0, TotalShipping:0, TotalTax:0, Status:"", ShippingMethod:0, IsPrintifyExpress:false, IsEconomyShipping:false, Shipments:[]order.Shipment(nil), CreatedAt:"", SentToProductionAt:"", FulfilledAt:"", PrintifyConnect:order.PrintifyConnect{Url:"", Id:""}}
}
