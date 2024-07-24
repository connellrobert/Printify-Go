package webhooks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                        = "/v1/shops"
	LIST_WEBHOOKS_FOR_SHOP_ENDPOINT = ENDPOINT + "/%d/webhooks.json"
	CREATE_WEBHOOK_ENDPOINT         = ENDPOINT + "/%d/webhooks.json"
	MODIFY_WEBHOOK_ENDPOINT         = ENDPOINT + "/%d/webhooks/%s.json"
	DELETE_WEBHOOK_ENDPOINT         = ENDPOINT + "/%d/webhooks/%s.json"
)

func ListWebhooksForShop(c *common.Client) ([]Webhook, error) {
	req, err := http.NewRequest("GET", LIST_WEBHOOKS_FOR_SHOP_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var webhooks []Webhook
	err = json.NewDecoder(resp.Body).Decode(&webhooks)
	if err != nil {
		return nil, err
	}
	return webhooks, nil
}

func CreateWebhook(c *common.Client, shopId int, topic, url, secret string) (*Webhook, error) {
	req, err := http.NewRequest("POST", CREATE_WEBHOOK_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var webhook Webhook
	err = json.NewDecoder(resp.Body).Decode(&webhook)
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

func ModifyWebhook(c *common.Client, shopId int, webhookId string, topic, url, secret string) (*Webhook, error) {
	req, err := http.NewRequest("PUT", MODIFY_WEBHOOK_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.PAT))
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var webhook Webhook
	err = json.NewDecoder(resp.Body).Decode(&webhook)
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

func DeleteWebhook(c *common.Client, shopId int, webhookId string) error {
	req, err := http.NewRequest("DELETE", DELETE_WEBHOOK_ENDPOINT, nil)
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
