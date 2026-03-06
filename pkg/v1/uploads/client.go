package uploads

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

// Client defines uploads operations and enables dependency injection.
type Client interface {
	ListUploadedImages() ([]Image, error)
	GetUploadedImage(id string) (*Image, error)
	UploadImage(body ImageUpload) (*Image, error)
	ArchiveUploadedImage(id string) error
}

type client struct {
	c *common.Client
}

// NewClient creates an uploads client implementation backed by common.Client.
func NewClient(c *common.Client) Client {
	return &client{c: c}
}

func (cl *client) ListUploadedImages() ([]Image, error) {
	return ListUploadedImages(cl.c)
}

func (cl *client) GetUploadedImage(id string) (*Image, error) {
	return GetUploadedImage(cl.c, id)
}

func (cl *client) UploadImage(body ImageUpload) (*Image, error) {
	return UploadImage(cl.c, body)
}

func (cl *client) ArchiveUploadedImage(id string) error {
	return ArchiveUploadedImage(cl.c, id)
}

var (
	ENDPOINT                        = "/v1/uploads"
	LIST_UPLOADED_IMAGES_ENDPOINT   = fmt.Sprintf("%s/images.json", ENDPOINT)
	GET_UPLOADED_IMAGE_ENDPOINT     = fmt.Sprintf("%s/images/%%s.json", ENDPOINT)
	UPLOAD_IMAGE_ENDPOINT           = fmt.Sprintf("%s/images.json", ENDPOINT)
	ARCHIVE_UPLOADED_IMAGE_ENDPOINT = fmt.Sprintf("%s/images/%%s/archive.json", ENDPOINT)
)

var (
	// ListUploadedImages calls GET /v1/uploads/images.json and returns uploaded images.
	//
	// Signature:
	//	func(c *common.Client) ([]Image, error)
	ListUploadedImages = common.ListResources[Image](LIST_UPLOADED_IMAGES_ENDPOINT)
	// GetUploadedImage calls GET /v1/uploads/images/{imageId}.json.
	//
	// Signature:
	//	func(c *common.Client, id string) (*Image, error)
	// Parameter mapping:
	//	id -> {imageId}
	//
	// imageId can be discovered with ListUploadedImages.
	GetUploadedImage = common.GetResourceById[Image, string](GET_UPLOADED_IMAGE_ENDPOINT)
	// UploadImage calls POST /v1/uploads/images.json.
	//
	// Signature:
	//	func(c *common.Client, body ImageUpload) (*Image, error)
	//
	// The request body supports uploading by raw bytes (Contents) or by Url.
	UploadImage = common.PostResourceWithReturn[ImageUpload, Image](UPLOAD_IMAGE_ENDPOINT)
	// ArchiveUploadedImage calls POST /v1/uploads/images/{imageId}/archive.json.
	//
	// Signature:
	//	func(c *common.Client, id string) error
	// Parameter mapping:
	//	id -> {imageId}
	//
	// imageId can be discovered with ListUploadedImages or GetUploadedImage.
	ArchiveUploadedImage = common.PostNoResourceWithoutReturn[string](ARCHIVE_UPLOADED_IMAGE_ENDPOINT)
)
