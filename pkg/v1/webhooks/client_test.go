package webhooks

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/connellrobert/printify-go/pkg/common"
)

func newWebhooksTestClient(register func(mux *http.ServeMux)) (*common.Client, func()) {
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

func ExampleListWebhooksForShop() {
	c, closeFn := newWebhooksTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/webhooks.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"data":[{"id":"wh_1","topic":"order:created","url":"https://example.com/hook"}]}`))
		})
	})
	defer closeFn()

	items, _ := ListWebhooksForShop(c)
	fmt.Printf("%#v\n", items[0])
	// Output: webhooks.Webhook{Id:"wh_1", Topic:"order:created", Url:"https://example.com/hook", ShopId:0, Secret:""}
}

func ExampleCreateWebhook() {
	c, closeFn := newWebhooksTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/webhooks.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"wh_new","topic":"order:created","url":"https://example.com/hook"}`))
		})
	})
	defer closeFn()

	item, _ := CreateWebhook(c, 123, Webhook{Topic: "order:created", Url: "https://example.com/hook"})
	fmt.Printf("%#v\n", *item)
	// Output: webhooks.Webhook{Id:"wh_new", Topic:"order:created", Url:"https://example.com/hook", ShopId:0, Secret:""}
}

func ExampleModifyWebhook() {
	c, closeFn := newWebhooksTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/webhooks/wh_1.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"wh_1","topic":"order:updated","url":"https://example.com/hook"}`))
		})
	})
	defer closeFn()

	item, _ := ModifyWebhook(c, 123, "wh_1", Webhook{Topic: "order:updated", Url: "https://example.com/hook"})
	fmt.Printf("%#v\n", *item)
	// Output: webhooks.Webhook{Id:"wh_1", Topic:"order:updated", Url:"https://example.com/hook", ShopId:0, Secret:""}
}

func ExampleDeleteWebhook() {
	c, closeFn := newWebhooksTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/shops/123/webhooks/wh_1.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := DeleteWebhook(c, 123, "wh_1")
	fmt.Println(err == nil)
	// Output: true
}
