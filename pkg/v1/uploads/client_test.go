package uploads

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/connellrobert/printify-go/pkg/common"
)

func newUploadsTestClient(register func(mux *http.ServeMux)) (*common.Client, func()) {
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

func ExampleListUploadedImages() {
	c, closeFn := newUploadsTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/uploads/images.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`[{"id":"img_1","file_name":"design.png"}]`))
		})
	})
	defer closeFn()

	items, _ := ListUploadedImages(c)
	fmt.Printf("%#v\n", items[0])
	// Output: uploads.Image{Id:"img_1", FileName:"design.png", Height:0, Width:0, Size:0, MimeType:"", PreviewUrl:"", UploadTime:""}
}

func ExampleGetUploadedImage() {
	c, closeFn := newUploadsTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/uploads/images/img_1.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"img_1","file_name":"design.png"}`))
		})
	})
	defer closeFn()

	item, _ := GetUploadedImage(c, "img_1")
	fmt.Printf("%#v\n", *item)
	// Output: uploads.Image{Id:"img_1", FileName:"design.png", Height:0, Width:0, Size:0, MimeType:"", PreviewUrl:"", UploadTime:""}
}

func ExampleUploadImage() {
	c, closeFn := newUploadsTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/uploads/images.json", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"id":"img_new","file_name":"new.png"}`))
		})
	})
	defer closeFn()

	item, _ := UploadImage(c, ImageUpload{Filename: "new.png", Url: "https://example.com/new.png"})
	fmt.Printf("%#v\n", *item)
	// Output: uploads.Image{Id:"img_new", FileName:"new.png", Height:0, Width:0, Size:0, MimeType:"", PreviewUrl:"", UploadTime:""}
}

func ExampleArchiveUploadedImage() {
	c, closeFn := newUploadsTestClient(func(mux *http.ServeMux) {
		mux.HandleFunc("/v1/uploads/images/img_1/archive.json", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	defer closeFn()

	err := ArchiveUploadedImage(c, "img_1")
	fmt.Println(err == nil)
	// Output: true
}
