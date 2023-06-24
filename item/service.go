package item

import "ntika/auth"

// import "fmt"

type Service interface {
	FindAll(filter string, sort string) ([]Item, error)
	// FindAllSongTags() ([]SongTag, error)
	// FindSongTag(tagid int, sondid int) ([]SongTag, error)
	// FindByID(ID int) (Tag, error)
	Create(tagInput ItemInput) (Item, error)
	// UpdateTag(ID int, tagInput TagInput) (Tag, error)
	// Delete(ID int) (Tag, error)
	// // ===========================================
	// AddTag(songTagInput SongTagInput, tag string, ytid string) (SongTag, error)
	// DeleteSongTag(ID int) (SongTag, error)
	// // ===========================================
	// GetSongByTag(tagid int) ([]SongTag, error)
	// GetTagsBySong(songid int) ([]SongTag, error)
	// FilterTag(tagid FilterInput) ([]SongTag, error)
	Order(user auth.User) (Orders, error)
	ACC(order OrderInput) (Orders, error)
	UpdateOrder(update_order UpdateOrderInput) (Orders, error)
	UserOrders(userID int) ([]Orders, error)
	AdminOrders() ([]Orders, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(filter string, sort string) ([]Item, error) {
	items, err := s.repository.FindAll(filter, sort)
	return items, err

}

// func (s *service) FindAllSongTags() ([]SongTag, error) {
// 	songtags, err := s.repository.FindAllSongTags()
// 	return songtags, err

// }

// func (s *service) FindByID(ID int) (Tag, error) {
// 	tag, err := s.repository.FindByID(ID)
// 	return tag, err
// }

func (s *service) Create(itemInput ItemInput) (Item, error) {
	// fmt.Println(tagInput.Tag)
	item := Item{
		Name:        itemInput.Name,
		Category:    itemInput.Category,
		Price:       itemInput.Price,
		Description: itemInput.Description,
		// Title:       bookInput.Title,
		// Price:       bookInput.Price,
		// Description: bookInput.Description,
		// Rating:      bookInput.Rating,
		// Discount:    bookInput.Discount,
	}
	newtag, err := s.repository.Create(item)
	return newtag, err
}

func (s *service) Order(user auth.User) (Orders, error) {
	// fmt.Println(tagInput.Tag)
	order := Orders{
		User_ID: user.ID,
		Address: user.Address,
		Status:  "Menunggu ACC",
	}
	newtag, err := s.repository.Order(order)
	return newtag, err
}

func (s *service) UserOrders(userID int) ([]Orders, error) {
	// fmt.Println(tagInput.Tag)
	orders, err := s.repository.FindOrdersByUser(userID)

	// newtag, err := s.repository.Order(order)
	return orders, err
}

func (s *service) AdminOrders() ([]Orders, error) {
	// fmt.Println(tagInput.Tag)
	orders, err := s.repository.AdminOrders()

	// newtag, err := s.repository.Order(order)
	return orders, err
}

func (s *service) ACC(orderInput OrderInput) (Orders, error) {
	// fmt.Println(tagInput.Tag)
	order, _ := s.repository.FindOrderByID(orderInput.ID)
	// order := Orders{
	// 	Status: "Menunggu ACC",
	// }
	order.Status = "ACC, Driver OTW"
	newtag, err := s.repository.ACC(order)
	return newtag, err
}

func (s *service) UpdateOrder(orderInput UpdateOrderInput) (Orders, error) {
	// fmt.Println(tagInput.Tag)
	order, _ := s.repository.FindOrderByID(orderInput.ID)
	// order := Orders{
	// 	Status: "Menunggu ACC",
	// }
	for _, order_item := range orderInput.OrderItems {
		order_item := OrderItem{
			Order_ID:   orderInput.ID,
			Product_ID: order_item.Product_ID,
			Quantity:   order_item.Quantity,
			Price:      order_item.Price,
		}
		_, _ = s.repository.CreateOrderItem(order_item)
	}

	orderItems, _ := s.repository.FindOrderItemByOrderID(orderInput.ID)

	subtotal := 0
	for _, order_item := range orderItems {
		subtotal += order_item.Price

	}

	order.Status = "Menunggu Konfirmasi Pelanggan"
	order.Sub_Total = subtotal
	order.Shipping_Fee = orderInput.Shipping_Fee
	total_price := subtotal + orderInput.Shipping_Fee
	order.Total_Price = total_price
	order.Shipping_Method = orderInput.Shipping_Method
	newtag, err := s.repository.UpdateOrderInput(order)
	return newtag, err
}

// func (s *service) UpdateTag(ID int, tagInput TagInput) (Tag, error) {

// 	tag, err := s.repository.FindByID(ID)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	tag.Tag = tagInput.Tag
// 	// book.Title = bookInput.Title
// 	// book.Price = bookInput.Price
// 	// book.Description = bookInput.Description
// 	// book.Rating = bookInput.Rating
// 	// book.Discount = bookInput.Discount

// 	newtag, err := s.repository.UpdateTag(tag)
// 	return newtag, err
// }

// func (s *service) Delete(ID int) (Tag, error) {
// 	tag, err := s.repository.FindByID(ID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	deleteTag, err := s.repository.Delete(tag)

// 	return deleteTag, err
// }

// func (s *service) AddTag(songTagInput SongTagInput, tag string, ytid string) (SongTag, error) {

// 	// song, err := song.s.repository.FindByID(ID)
// 	// tag := "temp"
// 	// ytid := "temp"
// 	songtag := SongTag{
// 		TagID:  songTagInput.TagID,
// 		Tag:    tag,
// 		SongID: songTagInput.SongID,
// 		YtID:   ytid,
// 		// Title:       bookInput.Title,
// 		// Price:       bookInput.Price,
// 		// Description: bookInput.Description,
// 		// Rating:      bookInput.Rating,
// 		// Discount:    bookInput.Discount,
// 	}
// 	fmt.Println("MASUK SERVICE")
// 	newsongtag, err := s.repository.AddTag(songtag)
// 	return newsongtag, err
// }

// func (s *service) DeleteSongTag(ID int) (SongTag, error) {
// 	songtag, err := s.repository.FindByIDSongTag(ID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	deleteTag, err := s.repository.DeleteSongTag(songtag)

// 	return deleteTag, err
// }

// func (s *service) FindSongTag(tagid int, songid int) ([]SongTag, error) {
// 	// var songtags = []SongTag
// 	songtags, err := s.repository.FindSongTag(tagid, songid)
// 	return songtags, err

// }

// func (s *service) GetSongByTag(tagid int) ([]SongTag, error) {
// 	// var songtags = []SongTag
// 	songtags, err := s.repository.GetSongByTag(tagid)

// 	return songtags, err
// }

// func (s *service) FilterTag(tagid FilterInput) ([]SongTag, error) {
// 	// var songtags = []SongTag
// 	ti := tagid.TagID
// 	fmt.Println(ti)
// 	songtags, err := s.repository.FilterTag(ti)

// 	return songtags, err
// }

// func (s *service) GetTagsBySong(songid int) ([]SongTag, error) {
// 	// var songtags = []SongTag
// 	songtags, err := s.repository.GetTagsBySong(songid)

// 	return songtags, err
// }
