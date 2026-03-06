package shop

// Shop represents a connected Printify shop.
type Shop struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	SalesChannel string `json:"sales_channel"`
}
