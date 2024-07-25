package webhooks

import (
	"github.com/connellrobert/printify-go/pkg/common"
)

var (
	ENDPOINT                        = "/v1/shops"
	LIST_WEBHOOKS_FOR_SHOP_ENDPOINT = ENDPOINT + "/%d/webhooks.json"
	CREATE_WEBHOOK_ENDPOINT         = ENDPOINT + "/%d/webhooks.json"
	MODIFY_WEBHOOK_ENDPOINT         = ENDPOINT + "/%d/webhooks/%s.json"
	DELETE_WEBHOOK_ENDPOINT         = ENDPOINT + "/%d/webhooks/%s.json"
)

var (
	ListWebhooksForShop = common.ListResources[Webhook](LIST_WEBHOOKS_FOR_SHOP_ENDPOINT)
	CreateWebhook       = common.PostResourceWithReturnAndId[Webhook, Webhook, int](CREATE_WEBHOOK_ENDPOINT)
	ModifyWebhook       = common.PutResourceWithReturnAndTwoId[Webhook, Webhook, int, string](MODIFY_WEBHOOK_ENDPOINT)
	DeleteWebhook       = common.DeleteResourceWithTwoId[Webhook, int, string](DELETE_WEBHOOK_ENDPOINT)
)
