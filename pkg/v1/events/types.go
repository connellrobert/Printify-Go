package events

// Event represents a webhook event payload from Printify.
type Event struct {
	Id        string        `json:"id"`
	Type      string        `json:"type"`
	CreatedAt string        `json:"created_at"`
	Resource  EventResource `json:"resource"`
}

// EventResource identifies the resource associated with an event.
type EventResource struct {
	Id   string            `json:"id"`
	Type string            `json:"type"`
	Data EventResourceData `json:"data"`
}

// EventResourceData contains resource-specific event metadata.
type EventResourceData struct {
	ShopId int    `json:"shop_id"`
	Reason string `json:"reason"`
}

type EventTypeEnum string

const (
	SHOP_DISCONNECTED        EventTypeEnum = "shop:disconnected"
	PRODUCT_DELETED          EventTypeEnum = "product:deleted"
	PRODUCT_PUBLISH_STARTED  EventTypeEnum = "product:publish_started"
	ORDER_CREATED            EventTypeEnum = "order:created"
	ORDER_UPDATED            EventTypeEnum = "order:updated"
	ORDER_SENT_TO_PRODUCTION EventTypeEnum = "order:sent-to-production"
	ORDER_SHIPMENT_CREATED   EventTypeEnum = "order:shipment:created"
	ORDER_SHIPMENT_DELIVERED EventTypeEnum = "order:shipment:delivered"
)
