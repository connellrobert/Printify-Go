package webhooks

import (
	"fmt"

	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                        = "/v1/shops"
	LIST_WEBHOOKS_FOR_SHOP_ENDPOINT = fmt.Sprintf("%s/%%d/webhooks.json", ENDPOINT)
	CREATE_WEBHOOK_ENDPOINT         = fmt.Sprintf("%s/%%d/webhooks.json", ENDPOINT)
	MODIFY_WEBHOOK_ENDPOINT         = fmt.Sprintf("%s/%%d/webhooks/%%s.json", ENDPOINT)
	DELETE_WEBHOOK_ENDPOINT         = fmt.Sprintf("%s/%%d/webhooks/%%s.json", ENDPOINT)
)

var (
	// ListWebhooksForShop calls GET /v1/shops/{shopId}/webhooks.json.
	//
	// Signature:
	//	func(c *common.Client) ([]Webhook, error)
	//
	// The shop id used by this endpoint is taken from the client that created the request.
	// Create the client with the desired shop id, or discover shops with shop.ListShops.
	ListWebhooksForShop = common.ListResources[Webhook](LIST_WEBHOOKS_FOR_SHOP_ENDPOINT)
	// CreateWebhook calls POST /v1/shops/{shopId}/webhooks.json.
	//
	// Signature:
	//	func(c *common.Client, id int, body Webhook) (*Webhook, error)
	// Parameter mapping:
	//	id -> {shopId}
	//	body -> webhook payload (topic and url)
	//
	// shopId can be discovered with shop.ListShops.
	CreateWebhook = common.PostResourceWithReturnAndId[Webhook, Webhook, int](CREATE_WEBHOOK_ENDPOINT)
	// ModifyWebhook calls PUT /v1/shops/{shopId}/webhooks/{webhookId}.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo string, body Webhook) (*Webhook, error)
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {webhookId}
	//	body -> updated webhook payload
	//
	// shopId can be discovered with shop.ListShops.
	// webhookId can be discovered with ListWebhooksForShop.
	ModifyWebhook = common.PutResourceWithReturnAndTwoId[Webhook, Webhook, int, string](MODIFY_WEBHOOK_ENDPOINT)
	// DeleteWebhook calls DELETE /v1/shops/{shopId}/webhooks/{webhookId}.json.
	//
	// Signature:
	//	func(c *common.Client, idOne int, idTwo string) error
	// Parameter mapping:
	//	idOne -> {shopId}
	//	idTwo -> {webhookId}
	//
	// shopId can be discovered with shop.ListShops.
	// webhookId can be discovered with ListWebhooksForShop.
	DeleteWebhook = common.DeleteResourceWithTwoId[Webhook, int, string](DELETE_WEBHOOK_ENDPOINT)
)
