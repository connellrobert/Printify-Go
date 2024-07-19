package events

type Event struct {
	Id        string        `json:"id"`
	Type      string        `json:"type"`
	CreatedAt string        `json:"created_at"`
	Resource  EventResource `json:"resource"`
}

type EventResource struct {
	Id   string            `json:"id"`
	Type string            `json:"type"`
	Data EventResourceData `json:"data"`
}

type EventResourceData struct {
	ShopId int    `json:"shop_id"`
	Reason string `json:"reason"`
}
