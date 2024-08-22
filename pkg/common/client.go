package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/connellrobert/printify-go/pkg/v1/pagination"
)

const (
	HOST = "https://api.printify.com"
)

type Client struct {
	Host   string
	Client *http.Client
	PAT    string
	ShopID int
}

func NewClient(pat string, shopId int) *Client {
	return &Client{
		Host:   HOST,
		Client: &http.Client{},
		PAT:    pat,
		ShopID: shopId,
	}
}

func ListResources[T any](endpoint string) func(c *Client) ([]T, error) {
	return func(c *Client) ([]T, error) {
		req, err := http.NewRequest("GET", c.Host+endpoint, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("error: %s", resp.Body)
		}

		var resources []T
		err = json.NewDecoder(resp.Body).Decode(&resources)
		if err != nil {
			return nil, err
		}

		return resources, nil
	}
}

func ListResourceWithId[T any, ID int | string](endpoint string) func(c *Client, id ID) ([]T, error) {
	return func(c *Client, id ID) ([]T, error) {
		url := fmt.Sprintf(c.Host+endpoint, id)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			var r map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&r)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("error: %+v", r)
		}

		var resources pagination.APIPagination[T]
		err = json.NewDecoder(resp.Body).Decode(&resources)
		if err != nil {
			return nil, err
		}
		// Check for nested data field
		return resources.Data, nil
	}
}
func ListResourceWithTwoID[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) ([]T, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO) ([]T, error) {
		url := fmt.Sprintf(c.Host+endpoint, idOne, idTwo)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		var resources []T
		err = json.NewDecoder(resp.Body).Decode(&resources)
		if err != nil {
			return nil, err
		}
		return resources, nil
	}
}

func GetResourceById[T any, ID int | string](endpoint string) func(c *Client, id ID) (*T, error) {
	return func(c *Client, id ID) (*T, error) {
		req, err := http.NewRequest("GET", fmt.Sprintf(c.Host+endpoint, id), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		var resource T
		err = json.NewDecoder(resp.Body).Decode(&resource)
		if err != nil {
			return nil, err
		}
		return &resource, nil
	}
}

func GetResourceWithTwoId[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
		url := fmt.Sprintf(c.Host+endpoint, idOne, idTwo)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		var resource T
		err = json.NewDecoder(resp.Body).Decode(&resource)
		if err != nil {
			return nil, err
		}
		return &resource, nil
	}
}

func PostResourceWithoutReturn[T any](endpoint string) func(c *Client, body T) error {
	return func(c *Client, body T) error {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", c.Host+endpoint, bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")
		resp, err := checkResponse(c, req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return nil
	}
}

func PostResourceWithReturn[T any, R any](endpoint string) func(c *Client, body T) (*R, error) {
	return func(c *Client, body T) (*R, error) {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", c.Host+endpoint, bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")
		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var resource R
		err = json.NewDecoder(resp.Body).Decode(&resource)
		if err != nil {
			return nil, err
		}

		return &resource, nil
	}
}

func PostResourceWithReturnTwoId[T any, R any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO, body T) (*R, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO, body T) (*R, error) {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("POST", fmt.Sprintf(c.Host+endpoint, idOne, idTwo), bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")

		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var resource R
		err = json.NewDecoder(resp.Body).Decode(&resource)
		if err != nil {
			return nil, err
		}

		return &resource, nil
	}
}

func PostResourceWithoutReturnTwoId[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO, body T) error {
	return func(c *Client, idOne IDONE, idTwo IDTWO, body T) error {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("POST", fmt.Sprintf(c.Host+endpoint, idOne, idTwo), bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")
		resp, err := checkResponse(c, req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return nil
	}
}

func PostNoResourceWithReturnTwoId[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
		req, err := http.NewRequest("POST", fmt.Sprintf(c.Host+endpoint, idOne, idTwo), nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.PAT)

		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var resource T
		err = json.NewDecoder(resp.Body).Decode(&resource)
		if err != nil {
			return nil, err
		}
		return &resource, nil
	}
}

func PostResourceWithReturnAndId[T any, R any, ID int | string](endpoint string) func(c *Client, id ID, body T) (*R, error) {
	return func(c *Client, id ID, body T) (*R, error) {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest("POST", fmt.Sprintf(c.Host+endpoint, id), bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")
		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var resource R
		err = json.NewDecoder(resp.Body).Decode(&resource)
		if err != nil {
			return nil, err
		}

		return &resource, nil
	}
}

func PostNoResourceWithoutReturnTwoId[IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) error {
	return func(c *Client, idOne IDONE, idTwo IDTWO) error {
		url := fmt.Sprintf(c.Host+endpoint, idOne, idTwo)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)

		resp, err := checkResponse(c, req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}
}

func PostNoResourceWithoutReturn[ID int | string](endpoint string) func(c *Client, id ID) error {
	return func(c *Client, id ID) error {
		url := fmt.Sprintf(c.Host+endpoint, id)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}
}

func PutResourceWithReturnAndTwoId[T any, R any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO, body T) (*R, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO, body T) (*R, error) {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("PUT", fmt.Sprintf(c.Host+endpoint, idOne, idTwo), bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")

		resp, err := checkResponse(c, req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var resource R
		err = json.NewDecoder(resp.Body).Decode(&resource)
		if err != nil {
			return nil, err
		}

		return &resource, nil
	}
}

func DeleteResourceWithTwoId[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) error {
	return func(c *Client, idOne IDONE, idTwo IDTWO) error {
		url := fmt.Sprintf(c.Host+endpoint, idOne, idTwo)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}
}

func DeleteResourceWithId[T any, ID int | string](endpoint string) func(c *Client, id ID) error {
	return func(c *Client, id ID) error {
		url := fmt.Sprintf(c.Host+endpoint, id)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := checkResponse(c, req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}
}

func checkResponse(c *Client, req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		var r map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error: %+v", r)
	}
	return resp, nil
}
