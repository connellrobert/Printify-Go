package shop

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/connellrobert/printify-go/pkg/common"
)

func newShopTestClient(register func(mux *http.ServeMux)) (*common.Client, func()) {
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

func ExampleListShops() {
	c, closeFn := newShopTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`[{"id":123,"title":"Main Shop","sales_channel":"api"}]`))
		})
	})
	defer closeFn()

	items, _ := ListShops(c)
	fmt.Printf("%#v\n", items[0])
	// Output: shop.Shop{Id:123, Title:"Main Shop", SalesChannel:"api"}
}

func ExampleDeleteShop() {
	c, closeFn := newShopTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := DeleteShop(c, 123)
	fmt.Println(err == nil)
	// Output: true
}
