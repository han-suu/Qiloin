package item

type ItemInput struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// type SongTagInput struct {
// 	TagID  int `json:"tagid" binding:"required"`
// 	SongID int `json:"songid" binding:"required"`
// 	// Tag    string `json:"tag" binding:"required"`
// 	// YtID   string `json:"ytid" binding:"required"`
// }

// type FilterInput struct {
// 	TagID []int `json:"tags" binding:"required"`
// }

type OrderInput struct {
	ID int `json:"id" binding:"required"`
}

type OrderItemInput struct {
	Product_ID int `json:"product_id" binding:"required"`
	Quantity   int `json:"qty" binding:"required"`
	Price      int `json:"price" binding:"required"`
}

type UpdateOrderInput struct {
	ID              int              `json:"id" binding:"required"`
	Shipping_Fee    int              `json:"shipping_fee"`
	Shipping_Method string           `json:"shipping_method"`
	OrderItems      []OrderItemInput `json:"order_items" binding:"required"`
}
