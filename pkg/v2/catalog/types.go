package catalog

// ShippingListAttributeCountry represents a country object in v2 shipping attributes.
type ShippingListAttributeCountry struct {
	Code string `json:"code,omitempty"`
}

// ShippingListAttributeHandlingTime represents shipping handling time ranges.
type ShippingListAttributeHandlingTime struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

// ShippingListAttributeShippingCostItem represents a single shipping cost item.
type ShippingListAttributeShippingCostItem struct {
	Amount   int    `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// ShippingListAttributeShippingCost represents first-item and additional-item costs.
type ShippingListAttributeShippingCost struct {
	FirstItem       ShippingListAttributeShippingCostItem `json:"first_item,omitempty"`
	AdditionalItems ShippingListAttributeShippingCostItem `json:"additional_items,omitempty"`
}

// ShippingListAttribute is the attributes block for a shipping record in the v2 list response.
type ShippingListAttribute struct {
	Name           string                            `json:"name,omitempty"`
	ShippingType   string                            `json:"shipping_type,omitempty"`
	VariantId      int                               `json:"variant_id,omitempty"`
	ShippingPlanId string                            `json:"shipping_plan_id,omitempty"`
	Country        ShippingListAttributeCountry      `json:"country,omitempty"`
	HandlingTime   ShippingListAttributeHandlingTime `json:"handling_time,omitempty"`
	ShippingCost   ShippingListAttributeShippingCost `json:"shipping_cost,omitempty"`
}

// ShippingList represents a single data item from v2 shipping list responses.
type ShippingList struct {
	Type      string                `json:"type"`
	Id        string                `json:"id"`
	Attribute ShippingListAttribute `json:"attributes"`
}

// ShippingInfo represents the top-level response from v2 shipping endpoints.
type ShippingInfo struct {
	Data  []SpecificShipping `json:"data"`
	Links map[string]string  `json:"links"`
}

// SpecificShipping represents shipping data for a country/variant combination.
type SpecificShipping struct {
	ShippingType   string       `json:"shipping_type"`
	Country        string       `json:"country"`
	VariantId      int          `json:"variant_id"`
	ShippingPlanId string       `json:"shipping_plan_id"`
	HandlingTime   HandlingTime `json:"handling_time"`
	ShippingCost   ShippingCost `json:"shipping_cost"`
}

// HandlingTime represents minimum and maximum handling durations.
type HandlingTime struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// ShippingCost represents first-item and additional-item shipping cost values.
type ShippingCost struct {
	FirstItem       Item `json:"first_item"`
	AdditionalItems Item `json:"additional_items"`
}

// Item represents a currency amount object used in shipping cost fields.
type Item struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
