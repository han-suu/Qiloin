package auth

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	// Type     string `json:"type" binding:"required"`
	// Title       string `json:"title" binding:"required"`
	// Price       int    `json:"price" binding:"required,number"`
	// Description string `json:"description" binding:"required"`
	// Rating      int    `json:"rating" binding:"required,number"`
	// Discount    int    `json:"discount" binding:"required,number"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddressInput struct {
	Address string `json:"address" binding:"required"`
	City    string `json:"city" binding:"required"`
}

type TagInput struct {
	Name string `json:"name" binding:"required"`
}
