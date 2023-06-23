package auth

type SongResponse struct {
	ID   int    `json:"id"`
	YtID string `json:"YtID"`
	// Title       string `json:"title"`
	// Price       int    `json:"price"`
	// Description string `json:"description"`
	// Rating      int    `json:"rating"`
	// Discount    int    `json:"discount"`
}

type UserResponse struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address string
}
