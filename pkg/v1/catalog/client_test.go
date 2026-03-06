package catalog

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/connellrobert/printify-go/pkg/common"
)

func newCatalogTestClient(register func(mux *http.ServeMux)) (*common.Client, func()) {
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

func ExampleListBlueprints() {
	c, closeFn := newCatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/catalog/blueprints.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`[{"id":1,"title":"Classic Tee","brand":"Brand","model":"M1","images":["img-1"]}]`))
		})
	})
	defer closeFn()

	items, _ := ListBlueprints(c)
	fmt.Printf("%#v\n", items[0])
	// Output: catalog.Blueprint{Id:1, Title:"Classic Tee", Brand:"Brand", Model:"M1", Images:[]string{"img-1"}}
}

func ExampleGetBlueprint() {
	c, closeFn := newCatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/catalog/blueprints/1.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":1,"title":"Classic Tee","brand":"Brand","model":"M1","images":["img-1"]}`))
		})
	})
	defer closeFn()

	item, _ := GetBlueprint(c, 1)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.Blueprint{Id:1, Title:"Classic Tee", Brand:"Brand", Model:"M1", Images:[]string{"img-1"}}
}

func ExampleListPrintProvidersByBlueprint() {
	c, closeFn := newCatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/catalog/blueprints/1/print_providers.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"id":10,"title":"Provider A","location":{"address1":"A1","address2":"","city":"Riga","country":"LV","region":"","zip":"1000"}}]}`))
		})
	})
	defer closeFn()

	items, _ := ListPrintProvidersByBlueprint(c, 1)
	fmt.Printf("%#v\n", items[0])
	// Output: catalog.PrintProvider{Id:10, Title:"Provider A", Location:catalog.Location{Address1:"A1", Address2:"", City:"Riga", Country:"LV", Region:"", Zip:"1000"}}
}

func ExampleListVariantsByBlueprintPrintProvider() {
	c, closeFn := newCatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/catalog/blueprints/1/print_providers/10/variants.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`[{"id":100,"title":"Black / M","options":{"color":"Black","size":"M"},"placeholders":[]}]`))
		})
	})
	defer closeFn()

	items, _ := ListVariantsByBlueprintPrintProvider(c, 1, 10)
	fmt.Printf("%#v\n", items[0])
	// Output: catalog.Variant{Id:100, Title:"Black / M", Options:catalog.VariantOptions{Color:"Black", Size:"M"}, Placeholders:[]catalog.Placeholder{}}
}

func ExampleGetShippingInformation() {
	c, closeFn := newCatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/catalog/blueprints/1/print_providers/10/shipping.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"handling_time":{"value":2,"unit":"day"},"profiles":[]}`))
		})
	})
	defer closeFn()

	item, _ := GetShippingInformation(c, 1, 10)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.Shipping{HandlingTime:catalog.HandlingTime{Value:2, Unit:"day"}, Profiles:[]catalog.Profile{}}
}

func ExampleListPrintProviders() {
	c, closeFn := newCatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/catalog/print_providers.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`[{"id":10,"title":"Provider A","location":{"address1":"A1","address2":"","city":"Riga","country":"LV","region":"","zip":"1000"}}]`))
		})
	})
	defer closeFn()

	items, _ := ListPrintProviders(c)
	fmt.Printf("%#v\n", items[0])
	// Output: catalog.PrintProvider{Id:10, Title:"Provider A", Location:catalog.Location{Address1:"A1", Address2:"", City:"Riga", Country:"LV", Region:"", Zip:"1000"}}
}

func ExampleGetPrintProvider() {
	c, closeFn := newCatalogTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/catalog/print_providers/10.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":10,"title":"Provider A","location":{"address1":"A1","address2":"","city":"Riga","country":"LV","region":"","zip":"1000"}}`))
		})
	})
	defer closeFn()

	item, _ := GetPrintProvider(c, 10)
	fmt.Printf("%#v\n", *item)
	// Output: catalog.PrintProvider{Id:10, Title:"Provider A", Location:catalog.Location{Address1:"A1", Address2:"", City:"Riga", Country:"LV", Region:"", Zip:"1000"}}
}
