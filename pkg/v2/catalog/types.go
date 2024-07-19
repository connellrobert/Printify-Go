package catalog

type ShippingList struct {
	Name string `json:"name"`
}

type SpecificShipping struct {
	ShippingType   string       `json:"shipping_type"`
	Country        string       `json:"country"`
	VariantId      int          `json:"variant_id"`
	ShippingPlanId string       `json:"shipping_plan_id"`
	HandlingTime   HandlingTime `json:"handling_time"`
	ShippingCost   ShippingCost `json:"shipping_cost"`
}

type HandlingTime struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type ShippingCost struct {
	FirstItem       Item `json:"first_item"`
	AdditionalItems Item `json:"additional_items"`
}

type Item struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
