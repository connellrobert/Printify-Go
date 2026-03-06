package catalog

// Blueprint describes a catalog blueprint returned by /v1/catalog/blueprints endpoints.
type Blueprint struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Brand  string   `json:"brand"`
	Model  string   `json:"model"`
	Images []string `json:"images"`
}

// PrintProvider describes a print provider available for one or more blueprints.
type PrintProvider struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Location Location `json:"location"`
}

// Location describes the provider address/location metadata in the catalog payload.
type Location struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Zip      string `json:"zip"`
}

// Variant describes a blueprint variant for a specific print provider.
type Variant struct {
	Id           int            `json:"id"`
	Title        string         `json:"title"`
	Options      VariantOptions `json:"options"`
	Placeholders []Placeholder  `json:"placeholders"`
}

// VariantOptions contains option labels (for example color and size) for a variant.
type VariantOptions struct {
	Color string `json:"color"`
	Size  string `json:"size"`
}

// Placeholder describes a printable area placeholder for a variant.
type Placeholder struct {
	Position int `json:"position"`
	Height   int `json:"height"`
	Width    int `json:"width"`
}

// Shipping contains shipping details for a blueprint/provider pair.
type Shipping struct {
	HandlingTime HandlingTime `json:"handling_time"`
	Profiles     []Profile    `json:"profiles"`
}

// HandlingTime describes production handling duration for shipping.
type HandlingTime struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

// Profile describes shipping rates and country coverage for specific variant ids.
type Profile struct {
	VariantIds      []int           `json:"variant_ids"`
	FirstItem       FirstItem       `json:"first_item"`
	AdditionalItems AdditionalItems `json:"additional_items"`
	Countries       []string        `json:"countries"`
}

// FirstItem is the first-item shipping price component.
type FirstItem struct {
	Currency string `json:"currency"`
	Cost     string `json:"cost"`
}

// AdditionalItems is the additional-items shipping price component.
type AdditionalItems struct {
	Currency string `json:"currency"`
	Cost     string `json:"cost"`
}

// PrintDetails represents print-side metadata returned in catalog responses.
type PrintDetails struct {
	PrintOnSide   string `json:"print_on_side"`
	SeparatorType string `json:"separator_type"`
}
