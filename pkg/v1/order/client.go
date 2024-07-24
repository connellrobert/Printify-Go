package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                               = "/v1/shops"
	LIST_ORDERS_ENDPOINT                   = fmt.Sprintf("%s/%%d/orders.json", ENDPOINT)
	GET_ORDER_DETAILS_ENDPOINT             = fmt.Sprintf("%s/%%d/orders/%%d.json", ENDPOINT)
	SUBMIT_ORDER_ENDPOINT                  = fmt.Sprintf("%s/%%d/orders.json", ENDPOINT)
	SUBMIT_PRINTIFY_EXPRESS_ORDER_ENDPOINT = fmt.Sprintf("%s/%%d/orders/express.json", ENDPOINT)
	SEND_ORDER_TO_PRODUCTION_ENDPOINT      = fmt.Sprintf("%s/%%d/orders/%%d/send_to_production.json", ENDPOINT)
	CALCULATE_SHIPPING_COSTS_ENDPOINT      = fmt.Sprintf("%s/%%d/orders/shipping.json", ENDPOINT)
	CANCEL_ORDER_ENDPOINT                  = fmt.Sprintf("%s/%%d/orders/%%d/cancel.json", ENDPOINT)
)

func ListOrders(c *common.Client, shopId int) ([]Order, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(LIST_ORDERS_ENDPOINT, shopId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var orders []Order
	err = json.NewDecoder(resp.Body).Decode(&orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderDetails(c *common.Client, shopId, orderId int) (*Order, error) {
	req, err := http.NewRequest("GET", c.Host+fmt.Sprintf(GET_ORDER_DETAILS_ENDPOINT, shopId, orderId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var order Order
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func SubmitOrder(c *common.Client, shopId int, order Order) (*Order, error) {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(SUBMIT_ORDER_ENDPOINT, shopId), nil)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(order)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(body)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submittedOrder Order
	err = json.NewDecoder(resp.Body).Decode(&submittedOrder)
	if err != nil {
		return nil, err
	}

	return &submittedOrder, nil
}

func SubmitPrintifyExpressOrder(c *common.Client, shopId int, order Order) (*Order, error) {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(SUBMIT_PRINTIFY_EXPRESS_ORDER_ENDPOINT, shopId), nil)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(order)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(body)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var submittedOrder Order
	err = json.NewDecoder(resp.Body).Decode(&submittedOrder)
	if err != nil {
		return nil, err
	}

	return &submittedOrder, nil
}

func SendOrderToProduction(c *common.Client, shopId, orderId int, order Order) error {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(SEND_ORDER_TO_PRODUCTION_ENDPOINT, shopId, orderId), nil)
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(order)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(body)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func CalculateShippingCosts(c *common.Client, shopId int, details ShipmentCalculationRequest) (*ShipmentCalculationResponse, error) {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(CALCULATE_SHIPPING_COSTS_ENDPOINT, shopId), nil)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(details)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(body)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var calculation ShipmentCalculationResponse
	err = json.NewDecoder(resp.Body).Decode(&calculation)
	if err != nil {
		return nil, err
	}

	return &calculation, nil
}

func CancelOrder(c *common.Client, shopId, orderId int) (*Order, error) {
	req, err := http.NewRequest("POST", c.Host+fmt.Sprintf(CANCEL_ORDER_ENDPOINT, shopId, orderId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var order Order
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
