package uploads

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

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
