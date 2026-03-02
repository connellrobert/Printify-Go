package catalog

type ShippingListAttributeCountry struct {
	Code string `json:"code,omitempty"`
}

type ShippingListAttributeHandlingTime struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type ShippingListAttributeShippingCostItem struct {
	Amount   int    `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type ShippingListAttributeShippingCost struct {
	FirstItem       ShippingListAttributeShippingCostItem `json:"first_item,omitempty"`
	AdditionalItems ShippingListAttributeShippingCostItem `json:"additional_items,omitempty"`
}

type ShippingListAttribute struct {
	Name           string                            `json:"name,omitempty"`
	ShippingType   string                            `json:"shipping_type,omitempty"`
	VariantId      int                               `json:"variant_id,omitempty"`
	ShippingPlanId string                            `json:"shipping_plan_id,omitempty"`
	Country        ShippingListAttributeCountry      `json:"country,omitempty"`
	HandlingTime   ShippingListAttributeHandlingTime `json:"handling_time,omitempty"`
	ShippingCost   ShippingListAttributeShippingCost `json:"shipping_cost,omitempty"`
}

type ShippingList struct {
	Type      string                `json:"type"`
	Id        string                `json:"id"`
	Attribute ShippingListAttribute `json:"attributes"`
}

type ShippingInfo struct {
	Data  []SpecificShipping `json:"data"`
	Links map[string]string  `json:"links"`
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
