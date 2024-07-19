package pagination

type APIPagination struct {
	FirstPageUrl    string `json:"first_page_url"`
	PreviousPageUrl string `json:"prev_page_url"`
	NextPageUrl     string `json:"next_page_url"`
	LastPageUrl     string `json:"last_page_url"`
	CurrentPage     int    `json:"current_page"`
	LastPage        int    `json:"last_page"`
	Total           int    `json:"total"`
	PerPage         int    `json:"per_page"`
	From            int    `json:"from"`
	To              int    `json:"to"`
}
