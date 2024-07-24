package uploads

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                        = "/v1/uploads"
	LIST_UPLOADED_IMAGES_ENDPOINT   = fmt.Sprintf("%s/images.json", ENDPOINT)
	GET_UPLOADED_IMAGE_ENDPOINT     = fmt.Sprintf("%s/images/%%s.json", ENDPOINT)
	UPLOAD_IMAGE_ENDPOINT           = fmt.Sprintf("%s/images.json", ENDPOINT)
	ARCHIVE_UPLOADED_IMAGE_ENDPOINT = fmt.Sprintf("%s/images/%%s/archive.json", ENDPOINT)
)

func ListUploadedImages(c *common.Client) ([]Image, error) {
	req, err := http.NewRequest("GET", c.Host+LIST_UPLOADED_IMAGES_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var images []Image
	err = json.NewDecoder(resp.Body).Decode(&images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func GetUploadedImage(c *common.Client, id string) (*Image, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(GET_UPLOADED_IMAGE_ENDPOINT, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var image Image
	err = json.NewDecoder(resp.Body).Decode(&image)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func UploadImage(c *common.Client, upload ImageUpload) (*Image, error) {
	req, err := http.NewRequest("POST", c.Host+UPLOAD_IMAGE_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(upload)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var image Image
	err = json.NewDecoder(resp.Body).Decode(&image)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func ArchiveUploadedImage(c *common.Client, id string) error {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(ARCHIVE_UPLOADED_IMAGE_ENDPOINT, id), nil)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
