package item

import "time"

type Item struct {
	ID          int
	Name        string
	Category    string
	Price       int
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ID         int
	User_ID    int
	Product_ID int
	Quantity   int
	Size       string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Orders struct {
	ID              int
	User_ID         int
	Sub_Total       int
	Shipping_Method string
	Shipping_Fee    int
	Total_Price     int
	Address         string
	City            string
	Status          string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	ID         int
	Order_ID   int
	Product_ID int
	Quantity   int
	Price      int

	CreatedAt time.Time
	UpdatedAt time.Time
}
