package representations

// APIResponse represents the full structure of the API response
type APIResponse struct {
	Data []Player `json:"data"`
	Meta struct {
		TotalPages  int  `json:"total_pages"`
		CurrentPage int  `json:"current_page"`
		NextPage    *int `json:"next_page"` // Nullable field
		PerPage     int  `json:"per_page"`
		TotalCount  int  `json:"total_count"`
	} `json:"meta"`
}
