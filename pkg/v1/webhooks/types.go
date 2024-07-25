package webhooks

type Webhook struct {
	Id     string `json:"id,omitempty"`
	Topic  string `json:"topic"`
	Url    string `json:"url"`
	ShopId int    `json:"shop_id,omitempty"`
	Secret string `json:"secret,omitempty"`
}
