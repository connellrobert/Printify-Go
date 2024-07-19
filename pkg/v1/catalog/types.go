package catalog

type Blueprint struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Brand  string   `json:"brand"`
	Model  string   `json:"model"`
	Images []string `json:"images"`
}

type PrintProvider struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Location Location `json:"location"`
}

type Location struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Zip      string `json:"zip"`
}

type Variant struct {
	Id           int            `json:"id"`
	Title        string         `json:"title"`
	Options      VariantOptions `json:"options"`
	Placeholders []Placeholder  `json:"placeholders"`
}

type VariantOptions struct {
	Color string `json:"color"`
	Size  string `json:"size"`
}

type Placeholder struct {
	Position int `json:"position"`
	Height   int `json:"height"`
	Width    int `json:"width"`
}

type Shipping struct {
	HandlingTime HandlingTime `json:"handling_time"`
	Profiles     []Profile    `json:"profiles"`
}

type HandlingTime struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Profile struct {
	VariantIds      []int           `json:"variant_ids"`
	FirstItem       FirstItem       `json:"first_item"`
	AdditionalItems AdditionalItems `json:"additional_items"`
	Countries       []string        `json:"countries"`
}

type FirstItem struct {
	Currency string `json:"currency"`
	Cost     string `json:"cost"`
}

type AdditionalItems struct {
	Currency string `json:"currency"`
	Cost     string `json:"cost"`
}

type PrintDetails struct {
	PrintOnSide   string `json:"print_on_side"`
	SeparatorType string `json:"separator_type"`
}
