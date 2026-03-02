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
	ListWebhooksForShop = common.ListResources[Webhook](LIST_WEBHOOKS_FOR_SHOP_ENDPOINT)
	CreateWebhook       = common.PostResourceWithReturnAndId[Webhook, Webhook, int](CREATE_WEBHOOK_ENDPOINT)
	ModifyWebhook       = common.PutResourceWithReturnAndTwoId[Webhook, Webhook, int, string](MODIFY_WEBHOOK_ENDPOINT)
	DeleteWebhook       = common.DeleteResourceWithTwoId[Webhook, int, string](DELETE_WEBHOOK_ENDPOINT)
)
