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
	ListUploadedImages   = common.ListResources[Image](LIST_UPLOADED_IMAGES_ENDPOINT)
	GetUploadedImage     = common.GetResourceById[Image, string](GET_UPLOADED_IMAGE_ENDPOINT)
	UploadImage          = common.PostResourceWithReturn[ImageUpload, Image](UPLOAD_IMAGE_ENDPOINT)
	ArchiveUploadedImage = common.PostNoResourceWithoutReturn[string](ARCHIVE_UPLOADED_IMAGE_ENDPOINT)
)
