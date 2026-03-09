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
		mux.HandleFunc("/v1/shops/123/products/5f2e9a3b7c1d4e8f90ab12cd.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"5f2e9a3b7c1d4e8f90ab12cd","title":"Hoodie"}`))
		})
	})
	defer closeFn()

	item, _ := GetProduct(c, 123, "5f2e9a3b7c1d4e8f90ab12cd")
	fmt.Printf("%#v\n", *item)
	// Output: product.Product{Id:"5f2e9a3b7c1d4e8f90ab12cd", Title:"Hoodie", Description:"", Tags:[]string(nil), Options:[]product.ProductOptions(nil), Variants:[]product.Variant(nil), Images:[]product.MockupImage(nil), CreatedAt:"", UpdateAt:"", Visible:false, BlueprintId:0, PrintProviderId:0, UserId:0, ShopId:0, PrintAreas:[]product.PrintArea(nil), PrintDetails:[]common.PrintDetails(nil), External:product.PublishReference{Id:"", Handle:"", ShippingTemplateId:""}, IsLocked:false, IsPrintifyExpressEligible:false, IsEconomyShippingEligible:false, IsPrintifyExpressEnabled:false, IsEconomyShippingEnabled:false, SalesChannelProperties:[]interface {}(nil)}
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
		mux.HandleFunc("/v1/shops/123/products/5f2e9a3b7c1d4e8f90ab12cd.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"5f2e9a3b7c1d4e8f90ab12cd","title":"Updated Hoodie"}`))
		})
	})
	defer closeFn()

	item, _ := UpdateProduct(c, 123, "5f2e9a3b7c1d4e8f90ab12cd", Product{})
	fmt.Printf("%#v\n", *item)
	// Output: product.Product{Id:"5f2e9a3b7c1d4e8f90ab12cd", Title:"Updated Hoodie", Description:"", Tags:[]string(nil), Options:[]product.ProductOptions(nil), Variants:[]product.Variant(nil), Images:[]product.MockupImage(nil), CreatedAt:"", UpdateAt:"", Visible:false, BlueprintId:0, PrintProviderId:0, UserId:0, ShopId:0, PrintAreas:[]product.PrintArea(nil), PrintDetails:[]common.PrintDetails(nil), External:product.PublishReference{Id:"", Handle:"", ShippingTemplateId:""}, IsLocked:false, IsPrintifyExpressEligible:false, IsEconomyShippingEligible:false, IsPrintifyExpressEnabled:false, IsEconomyShippingEnabled:false, SalesChannelProperties:[]interface {}(nil)}
}

func ExampleDeleteProduct() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/5f2e9a3b7c1d4e8f90ab12cd.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := DeleteProduct(c, 123, "5f2e9a3b7c1d4e8f90ab12cd")
	fmt.Println(err == nil)
	// Output: true
}

func ExamplePublishProduct() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/5f2e9a3b7c1d4e8f90ab12cd/publish.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := PublishProduct(c, 123, "5f2e9a3b7c1d4e8f90ab12cd", Publish{})
	fmt.Println(err == nil)
	// Output: true
}

func ExampleUpdatePublishStatusToSucceeded() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/5f2e9a3b7c1d4e8f90ab12cd/publishing_succeeded.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := UpdatePublishStatusToSucceeded(c, 123, "5f2e9a3b7c1d4e8f90ab12cd", PublishReference{})
	fmt.Println(err == nil)
	// Output: true
}

func ExampleUpdatePublishStatusToFailed() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/5f2e9a3b7c1d4e8f90ab12cd/publishing_failed.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := UpdatePublishStatusToFailed(c, 123, "5f2e9a3b7c1d4e8f90ab12cd", PublishFailedRequest{})
	fmt.Println(err == nil)
	// Output: true
}

func ExampleNotifyProductUnpublished() {
	c, closeFn := newProductTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/products/5f2e9a3b7c1d4e8f90ab12cd/unpublished.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := NotifyProductUnpublished(c, 123, "5f2e9a3b7c1d4e8f90ab12cd")
	fmt.Println(err == nil)
	// Output: true
}
