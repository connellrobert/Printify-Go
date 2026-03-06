package catalog

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/connellrobert/printify-go/pkg/common"
)

func newV2CatalogTestClient(register func(mux *http.ServeMux)) (*common.Client, func()) {
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

func ExampleGetShippingForVariantsOfBlueprintById() {
	c, closeFn := newV2CatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v2/catalog/blueprints/1/print_providers/2/shipping.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"shipping_type":"standard","country":"US","variant_id":100,"shipping_plan_id":"sp1","handling_time":{"from":1,"to":2},"shipping_cost":{"first_item":{"amount":500,"currency":"USD"},"additional_items":{"amount":300,"currency":"USD"}}}],"links":{}}`))
		})
	})
	defer closeFn()

	item, _ := GetShippingForVariantsOfBlueprintById(c, 1, 2)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.ShippingInfo{Data:[]catalog.SpecificShipping{catalog.SpecificShipping{ShippingType:"standard", Country:"US", VariantId:100, ShippingPlanId:"sp1", HandlingTime:catalog.HandlingTime{From:1, To:2}, ShippingCost:catalog.ShippingCost{FirstItem:catalog.Item{Amount:500, Currency:"USD"}, AdditionalItems:catalog.Item{Amount:300, Currency:"USD"}}}}, Links:map[string]string{}}
}

func ExampleGetShippingStandardInfoForVariantsOfBlueprintById() {
	c, closeFn := newV2CatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v2/catalog/blueprints/1/print_providers/2/shipping/standard.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"shipping_type":"standard","country":"US","variant_id":100,"shipping_plan_id":"sp1","handling_time":{"from":1,"to":2},"shipping_cost":{"first_item":{"amount":500,"currency":"USD"},"additional_items":{"amount":300,"currency":"USD"}}}],"links":{}}`))
		})
	})
	defer closeFn()

	item, _ := GetShippingStandardInfoForVariantsOfBlueprintById(c, 1, 2)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.ShippingInfo{Data:[]catalog.SpecificShipping{catalog.SpecificShipping{ShippingType:"standard", Country:"US", VariantId:100, ShippingPlanId:"sp1", HandlingTime:catalog.HandlingTime{From:1, To:2}, ShippingCost:catalog.ShippingCost{FirstItem:catalog.Item{Amount:500, Currency:"USD"}, AdditionalItems:catalog.Item{Amount:300, Currency:"USD"}}}}, Links:map[string]string{}}
}

func ExampleGetShippingPriorityInfoForVariantsOfBlueprintById() {
	c, closeFn := newV2CatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v2/catalog/blueprints/1/print_providers/2/shipping/priority.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"shipping_type":"priority","country":"US","variant_id":100,"shipping_plan_id":"sp1","handling_time":{"from":1,"to":2},"shipping_cost":{"first_item":{"amount":700,"currency":"USD"},"additional_items":{"amount":400,"currency":"USD"}}}],"links":{}}`))
		})
	})
	defer closeFn()

	item, _ := GetShippingPriorityInfoForVariantsOfBlueprintById(c, 1, 2)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.ShippingInfo{Data:[]catalog.SpecificShipping{catalog.SpecificShipping{ShippingType:"priority", Country:"US", VariantId:100, ShippingPlanId:"sp1", HandlingTime:catalog.HandlingTime{From:1, To:2}, ShippingCost:catalog.ShippingCost{FirstItem:catalog.Item{Amount:700, Currency:"USD"}, AdditionalItems:catalog.Item{Amount:400, Currency:"USD"}}}}, Links:map[string]string{}}
}

func ExampleGetShippingExpressInfoForVariantsOfBlueprintById() {
	c, closeFn := newV2CatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v2/catalog/blueprints/1/print_providers/2/shipping/express.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"shipping_type":"express","country":"US","variant_id":100,"shipping_plan_id":"sp1","handling_time":{"from":1,"to":2},"shipping_cost":{"first_item":{"amount":900,"currency":"USD"},"additional_items":{"amount":500,"currency":"USD"}}}],"links":{}}`))
		})
	})
	defer closeFn()

	item, _ := GetShippingExpressInfoForVariantsOfBlueprintById(c, 1, 2)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.ShippingInfo{Data:[]catalog.SpecificShipping{catalog.SpecificShipping{ShippingType:"express", Country:"US", VariantId:100, ShippingPlanId:"sp1", HandlingTime:catalog.HandlingTime{From:1, To:2}, ShippingCost:catalog.ShippingCost{FirstItem:catalog.Item{Amount:900, Currency:"USD"}, AdditionalItems:catalog.Item{Amount:500, Currency:"USD"}}}}, Links:map[string]string{}}
}

func ExampleGetShippingEconomyInfoForVariantsOfBlueprintById() {
	c, closeFn := newV2CatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v2/catalog/blueprints/1/print_providers/2/shipping/economy.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"shipping_type":"economy","country":"US","variant_id":100,"shipping_plan_id":"sp1","handling_time":{"from":1,"to":2},"shipping_cost":{"first_item":{"amount":300,"currency":"USD"},"additional_items":{"amount":200,"currency":"USD"}}}],"links":{}}`))
		})
	})
	defer closeFn()

	item, _ := GetShippingEconomyInfoForVariantsOfBlueprintById(c, 1, 2)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.ShippingInfo{Data:[]catalog.SpecificShipping{catalog.SpecificShipping{ShippingType:"economy", Country:"US", VariantId:100, ShippingPlanId:"sp1", HandlingTime:catalog.HandlingTime{From:1, To:2}, ShippingCost:catalog.ShippingCost{FirstItem:catalog.Item{Amount:300, Currency:"USD"}, AdditionalItems:catalog.Item{Amount:200, Currency:"USD"}}}}, Links:map[string]string{}}
}
