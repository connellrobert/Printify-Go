package product

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/connellrobert/printify-go/pkg/common"
)

func newProductTestClient(register func(mux *http.ServeMux)) (*common.Client, func()) {
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

func ExampleListProducts() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"id":"prod_1","title":"T-Shirt"}]}`))
		})
	})
	defer closeFn()

	items, _ := ListProducts(c, 123)
	fmt.Printf("%#v\n", items[0])
	// Output: product.Product{Id:"prod_1", Title:"T-Shirt", Description:"", Tags:[]string(nil), Options:[]product.ProductOptions(nil), Variants:[]product.Variant(nil), Images:[]product.MockupImage(nil), CreatedAt:"", UpdateAt:"", Visible:false, BlueprintId:0, PrintProviderId:0, UserId:0, ShopId:0, PrintAreas:[]product.PrintArea(nil), PrintDetails:[]common.PrintDetails(nil), External:product.PublishReference{Id:"", Handle:"", ShippingTemplateId:""}, IsLocked:false, IsPrintifyExpressEligible:false, IsEconomyShippingEligible:false, IsPrintifyExpressEnabled:false, IsEconomyShippingEnabled:false, SalesChannelProperties:[]interface {}(nil)}
}

func ExampleGetProduct() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/456.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"prod_456","title":"Hoodie"}`))
		})
	})
	defer closeFn()

	item, _ := GetProduct(c, 123, 456)
	fmt.Printf("%#v\n", *item)
	// Output: product.Product{Id:"prod_456", Title:"Hoodie", Description:"", Tags:[]string(nil), Options:[]product.ProductOptions(nil), Variants:[]product.Variant(nil), Images:[]product.MockupImage(nil), CreatedAt:"", UpdateAt:"", Visible:false, BlueprintId:0, PrintProviderId:0, UserId:0, ShopId:0, PrintAreas:[]product.PrintArea(nil), PrintDetails:[]common.PrintDetails(nil), External:product.PublishReference{Id:"", Handle:"", ShippingTemplateId:""}, IsLocked:false, IsPrintifyExpressEligible:false, IsEconomyShippingEligible:false, IsPrintifyExpressEnabled:false, IsEconomyShippingEnabled:false, SalesChannelProperties:[]interface {}(nil)}
}

func ExampleCreateProduct() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"prod_new","title":"Poster"}`))
		})
	})
	defer closeFn()

	item, _ := CreateProduct(c, 123, Product{})
	fmt.Printf("%#v\n", *item)
	// Output: product.Product{Id:"prod_new", Title:"Poster", Description:"", Tags:[]string(nil), Options:[]product.ProductOptions(nil), Variants:[]product.Variant(nil), Images:[]product.MockupImage(nil), CreatedAt:"", UpdateAt:"", Visible:false, BlueprintId:0, PrintProviderId:0, UserId:0, ShopId:0, PrintAreas:[]product.PrintArea(nil), PrintDetails:[]common.PrintDetails(nil), External:product.PublishReference{Id:"", Handle:"", ShippingTemplateId:""}, IsLocked:false, IsPrintifyExpressEligible:false, IsEconomyShippingEligible:false, IsPrintifyExpressEnabled:false, IsEconomyShippingEnabled:false, SalesChannelProperties:[]interface {}(nil)}
}

func ExampleUpdateProduct() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/456.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"prod_456","title":"Updated Hoodie"}`))
		})
	})
	defer closeFn()

	item, _ := UpdateProduct(c, 123, 456, Product{})
	fmt.Printf("%#v\n", *item)
	// Output: product.Product{Id:"prod_456", Title:"Updated Hoodie", Description:"", Tags:[]string(nil), Options:[]product.ProductOptions(nil), Variants:[]product.Variant(nil), Images:[]product.MockupImage(nil), CreatedAt:"", UpdateAt:"", Visible:false, BlueprintId:0, PrintProviderId:0, UserId:0, ShopId:0, PrintAreas:[]product.PrintArea(nil), PrintDetails:[]common.PrintDetails(nil), External:product.PublishReference{Id:"", Handle:"", ShippingTemplateId:""}, IsLocked:false, IsPrintifyExpressEligible:false, IsEconomyShippingEligible:false, IsPrintifyExpressEnabled:false, IsEconomyShippingEnabled:false, SalesChannelProperties:[]interface {}(nil)}
}

func ExampleDeleteProduct() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/456.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := DeleteProduct(c, 123, 456)
	fmt.Println(err == nil)
	// Output: true
}

func ExamplePublishProduct() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/456/publish.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := PublishProduct(c, 123, 456, Publish{})
	fmt.Println(err == nil)
	// Output: true
}

func ExampleUpdatePublishStatusToSucceeded() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/456/publishing_succeeded.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := UpdatePublishStatusToSucceeded(c, 123, 456, PublishReference{})
	fmt.Println(err == nil)
	// Output: true
}

func ExampleUpdatePublishStatusToFailed() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/456/publishing_failed.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := UpdatePublishStatusToFailed(c, 123, 456, PublishFailedRequest{})
	fmt.Println(err == nil)
	// Output: true
}

func ExampleNotifyProductUnpublished() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/456/unpublished.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := NotifyProductUnpublished(c, 123, 456)
	fmt.Println(err == nil)
	// Output: true
}
