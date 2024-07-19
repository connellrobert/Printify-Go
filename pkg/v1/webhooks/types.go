package webhooks

type Webhook struct {
	Id     string `json:"id"`
	Topic  string `json:"topic"`
	Url    string `json:"url"`
	ShopId int    `json:"shop_id"`
	Secret string `json:"secret"`
}
