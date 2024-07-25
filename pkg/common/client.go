package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	HOST = "https://api.printify.com"
)

type Client struct {
	Host   string
	Client *http.Client
	PAT    string
	ShopID string
}

func NewClient(pat string, shopId int) *Client {
	return &Client{
		Host:   HOST,
		Client: &http.Client{},
		PAT:    pat,
	}
}

func ListResources[T any](endpoint string) func(c *Client) ([]T, error) {
	return func(c *Client) ([]T, error) {
		req, err := http.NewRequest("GET", c.Host+endpoint, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)

		resp, err := c.Client.Do(req)
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
		req, err := http.NewRequest("GET", c.Host+endpoint, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := c.Client.Do(req)
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

func PostResourceWithoutReturn[T any](c *Client, endpoint string, body T) error {
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

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
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

		resp, err := c.Client.Do(req)
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

		req, err := http.NewRequest("POST", c.Host+endpoint, bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.Client.Do(req)
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

		req, err := http.NewRequest("POST", c.Host+endpoint, bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.Client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return nil
	}
}

func PostNoResourceWithReturnTwoId[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
		req, err := http.NewRequest("POST", c.Host+endpoint, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)

		resp, err := c.Client.Do(req)
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

func PutResourceWithReturnAndTwoId[T any, R any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO, body T) (*R, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO, body T) (*R, error) {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest("PUT", c.Host+endpoint, bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.Client.Do(req)
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

func PostResourceWithReturnAndId[T any, R any, ID int | string](endpoint string) func(c *Client, id ID, body T) (*R, error) {
	return func(c *Client, id ID, body T) (*R, error) {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		url := fmt.Sprintf(c.Host+endpoint, id)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+c.PAT)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.Client.Do(req)
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

func ListResourceWithId[T any, ID int | string](endpoint string) func(c *Client, id ID) ([]T, error) {
	return func(c *Client, id ID) ([]T, error) {
		url := fmt.Sprintf(c.Host+endpoint, id)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := c.Client.Do(req)
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
func ListResourceWithTwoID[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) ([]T, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO) ([]T, error) {
		url := fmt.Sprintf(c.Host+endpoint, idOne, idTwo)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := c.Client.Do(req)
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
func GetResourceWithTwoId[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
	return func(c *Client, idOne IDONE, idTwo IDTWO) (*T, error) {
		url := fmt.Sprintf(c.Host+endpoint, idOne, idTwo)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := c.Client.Do(req)
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

func DeleteResourceWithTwoId[T any, IDONE, IDTWO int | string](endpoint string) func(c *Client, idOne IDONE, idTwo IDTWO) error {
	return func(c *Client, idOne IDONE, idTwo IDTWO) error {
		url := fmt.Sprintf(c.Host+endpoint, idOne, idTwo)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", "Bearer "+c.PAT)
		resp, err := c.Client.Do(req)
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
		resp, err := c.Client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
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
		resp, err := c.Client.Do(req)
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
		resp, err := c.Client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}
}