package representations

// Player represents the structure of each NBA player in the API response
type Player struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Position     string `json:"position"`
	HeightFeet   *int   `json:"height_feet"`   // Using pointer for nullable fields
	HeightInches *int   `json:"height_inches"` // Using pointer for nullable fields
	WeightPounds *int   `json:"weight_pounds"` // Using pointer for nullable fields
	Team         Team   `json:"team"`
}
