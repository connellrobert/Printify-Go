package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                                    = "/v1/shops"
	LIST_PRODUCTS_ENDPOINT                      = fmt.Sprintf("%s/%%d/products.json", ENDPOINT)
	GET_PRODUCT_ENDPOINT                        = fmt.Sprintf("%s/%%d/products/%%d.json", ENDPOINT)
	CREATE_PRODUCT_ENDPOINT                     = fmt.Sprintf("%s/%%d/products.json", ENDPOINT)
	UPDATE_PRODUCT_ENDPOINT                     = fmt.Sprintf("%s/%%d/products/%%d.json", ENDPOINT)
	DELETE_PRODUCT_ENDPOINT                     = fmt.Sprintf("%s/%%d/products/%%d.json", ENDPOINT)
	PUBLISH_PRODUCT_ENDPOINT                    = fmt.Sprintf("%s/%%d/products/%%d/publish.json", ENDPOINT)
	UPDATE_PUBLISH_STATUS_TO_SUCCEEDED_ENDPOINT = fmt.Sprintf("%s/%%d/products/%%d/publishing_succeeded.json", ENDPOINT)
	UPDATE_PUBLISH_STATUS_TO_FAILED_ENDPOINT    = fmt.Sprintf("%s/%%d/products/%%d/publishing_failed.json", ENDPOINT)
	NOTIFY_PRODUCT_UNPUBLISHED_ENDPOINT         = fmt.Sprintf("%s/%%d/products/%%d/unpublished.json", ENDPOINT)
)

func ListProducts(c *common.Client, shopId int) ([]Product, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(LIST_PRODUCTS_ENDPOINT, shopId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(c *common.Client, shopId, productId int) (*Product, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(GET_PRODUCT_ENDPOINT, shopId, productId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var product Product
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func CreateProduct(c *common.Client, shopId int, product Product) (*Product, error) {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(CREATE_PRODUCT_ENDPOINT, shopId), nil)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(product)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var newProduct Product
	err = json.NewDecoder(resp.Body).Decode(&newProduct)
	if err != nil {
		return nil, err
	}

	return &newProduct, nil
}

func UpdateProduct(c *common.Client, shopId, productId int, product Product) (*Product, error) {
	req, err := http.NewRequest("PUT", c.Host+fmt.Sprintf(UPDATE_PRODUCT_ENDPOINT, shopId, productId), nil)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(product)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedProduct Product
	err = json.NewDecoder(resp.Body).Decode(&updatedProduct)
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

func DeleteProduct(c *common.Client, shopId, productId int) error {
	req, err := http.NewRequest("DELETE", c.Host+fmt.Sprintf(DELETE_PRODUCT_ENDPOINT, shopId, productId), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func PublishProduct(c *common.Client, shopId, productId int, publish Publish) error {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(PUBLISH_PRODUCT_ENDPOINT, shopId, productId), nil)
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(publish)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func UpdatePublishStatusToSucceeded(c *common.Client, shopId, productId int, external PublishReference) error {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(PUBLISH_PRODUCT_ENDPOINT, shopId, productId), nil)
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(external)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func UpdatePublishStatusToFailed(c *common.Client, shopId, productId int, reason PublishFailedRequest) error {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(PUBLISH_PRODUCT_ENDPOINT, shopId, productId), nil)
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(reason)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func NotifyProductUnpublished(c *common.Client, shopId, productId int) error {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(NOTIFY_PRODUCT_UNPUBLISHED_ENDPOINT, shopId, productId), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
