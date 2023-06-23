package item

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(filter string, sort string) ([]Item, error)
	// FindAllSongTags() ([]SongTag, error)
	// FindByID(ID int) (Tag, error)
	Create(item Item) (Item, error)
	// UpdateTag(tag Tag) (Tag, error)
	// Delete(tag Tag) (Tag, error)
	// // ======================================
	// FindAllSongTag() ([]SongTag, error)
	// FindByIDSongTag(ID int) (SongTag, error)
	// FindSongTag(tagid int, songid int) ([]SongTag, error)
	// AddTag(songTag SongTag) (SongTag, error)
	// DeleteSongTag(songTag SongTag) (SongTag, error)
	// // ==========================================
	// GetSongByTag(tagid int) ([]SongTag, error)
	// GetTagsBySong(songidid int) ([]SongTag, error)
	// FilterTag(tagid []int) ([]SongTag, error)
	Order(order Orders) (Orders, error)
	FindOrderByID(id int) (Orders, error)
	ACC(order Orders) (Orders, error)
	CreateOrderItem(orderItem OrderItem) (OrderItem, error)
	UpdateOrderInput(order Orders) (Orders, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(filter string, sort string) ([]Item, error) {
	var item []Item
	// if filter != "" {
	// 	// err := r.db.Find(&item).Error

	// 	err := r.db.Where(&Item{Category: filter}).Find(&item).Error
	// 	if err != nil {
	// 		println("=====================")
	// 		println("ERROR WHILE F")
	// 		println("=====================")
	// 	}

	// 	return item, err
	// } else {
	// 	err := r.db.Find(&item).Error
	// 	if err != nil {
	// 		println("=====================")
	// 		println("ERROR WHILE F")
	// 		println("=====================")
	// 	}

	// 	return item, err
	// }

	base := r.db.Debug()
	// if filter != "" {
	// 	base = base.Where(&Item{Category: filter}) // Adding this condition just if someThing is true
	// }
	// if sort != "" {
	// 	list := strings.Split(sort, "_")
	// 	val := list[0]
	// 	order := list[1]
	// 	query := fmt.Sprintf("%s %s", val, order)
	// 	base = base.Order(query)
	// }
	err := base.Find(&item).Error // Query your results
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return item, err

}

// func (r *repository) FindAllSongTags() ([]SongTag, error) {
// 	var songtag []SongTag

// 	err := r.db.Find(&songtag).Error
// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE F")
// 		println("=====================")
// 	}

// 	return songtag, err
// }

// func (r *repository) FindByID(ID int) (Tag, error) {
// 	var tag Tag

// 	err := r.db.Find(&tag, ID).Error
// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE FBI")
// 		println("=====================")
// 	}

// 	return tag, err
// }

func (r *repository) Create(item Item) (Item, error) {

	err := r.db.Create(&item).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return item, err
}

func (r *repository) Order(order Orders) (Orders, error) {

	err := r.db.Create(&order).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE ORDERING")
		println("=====================")
	}

	return order, err
}

func (r *repository) FindOrderByID(ID int) (Orders, error) {
	var tag Orders

	err := r.db.Find(&tag, ID).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE FBI")
		println("=====================")
	}

	return tag, err
}

func (r *repository) ACC(order Orders) (Orders, error) {

	err := r.db.Save(&order).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE Updating")
		println("=====================")
	}

	return order, err
}

func (r *repository) CreateOrderItem(orderItem OrderItem) (OrderItem, error) {

	err := r.db.Create(&orderItem).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE MakeOrderItem")
		println("=====================")
	}

	return orderItem, err
}

func (r *repository) UpdateOrderInput(order Orders) (Orders, error) {

	err := r.db.Save(&order).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE Updating")
		println("=====================")
	}

	return order, err
}

// func (r *repository) UpdateTag(tag Tag) (Tag, error) {

// 	err := r.db.Save(&tag).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE Updating")
// 		println("=====================")
// 	}

// 	return tag, err
// }

// func (r *repository) Delete(tag Tag) (Tag, error) {

// 	err := r.db.Delete(&tag).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE Deleting")
// 		println("=====================")
// 	}

// 	return tag, err
// }

// =============================================================

// func (r *repository) FindAllSongTag() ([]SongTag, error) {
// 	var songtag []SongTag

// 	err := r.db.Find(&songtag).Error
// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE F")
// 		println("=====================")
// 	}

// 	return songtag, err
// }

// func (r *repository) FindByIDSongTag(ID int) (SongTag, error) {
// 	var songtag SongTag

// 	err := r.db.Find(&songtag, ID).Error
// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE FBI")
// 		println("=====================")
// 	}

// 	return songtag, err
// }

// func (r *repository) FindSongTag(tagid int, songid int) ([]SongTag, error) {
// 	var songtag []SongTag

// 	// err := r.db.Find(&songtag, ID).Error
// 	err := r.db.Where(&SongTag{TagID: tagid, SongID: songid}).Find(&songtag).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE FST")
// 		println("=====================")
// 	}

// 	return songtag, err
// }

// func (r *repository) GetSongByTag(tagid int) ([]SongTag, error) {
// 	var songtag []SongTag

// 	// err := r.db.Find(&songtag, ID).Error
// 	err := r.db.Where(&SongTag{TagID: tagid}).Find(&songtag).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE GSBT")
// 		println("=====================")
// 	}

// 	return songtag, err
// }

// func (r *repository) FilterTag(tagid []int) ([]SongTag, error) {
// 	var songtag []SongTag
// 	var song []SongTag
// 	err := r.db.Where("tag_id IN ?", tagid).Find(&songtag).Error

// 	var arr []int
// 	for _, item := range songtag {
// 		arr = append(arr, item.SongID)
// 	}

// 	// MENCARI FREKUENSI VIDEO
// 	freq := make(map[int]int)
// 	for _, num := range arr {
// 		freq[num] = freq[num] + 1
// 	}

// 	// APAKAH FREKUENSI SAMA DENGAN JUMLAH TAG ID? YA : masukkan ke filtered ID
// 	var filteredid []int
// 	for _, item := range arr {
// 		if freq[item] == len(tagid) {
// 			filteredid = append(filteredid, item)
// 		}
// 	}
// 	// Menghapus data duplicate
// 	filteredid = removeDuplicateValues(filteredid)

// 	// KALAU MAU RETURN NYA SONG AJA BISA MAKE :
// 	// var song []song.Song
// 	// db.Find(&song, filteredid)

// 	fmt.Println(filteredid)
// 	for _, id := range filteredid {
// 		// r.db.First(&songtag, id)
// 		r.db.Where("song_id = ?", id).First(&songtag)
// 		song = append(song, songtag[0])
// 	}

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE FILTERING")
// 		println("=====================")
// 	}

// 	return song, err
// }

// func (r *repository) GetTagsBySong(songid int) ([]SongTag, error) {
// 	var songtag []SongTag

// 	// err := r.db.Find(&songtag, ID).Error
// 	err := r.db.Where(&SongTag{SongID: songid}).Find(&songtag).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE GSBT")
// 		println("=====================")
// 	}

// 	return songtag, err
// }

// func (r *repository) AddTag(songTag SongTag) (SongTag, error) {
// 	fmt.Println("MASUK REPO")
// 	err := r.db.Create(&songTag).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE CREATING")
// 		println("=====================")
// 	}

// 	return songTag, err
// }

// func (r *repository) DeleteSongTag(songtag SongTag) (SongTag, error) {

// 	err := r.db.Delete(&songtag).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE Deleting")
// 		println("=====================")
// 	}

// 	return songtag, err
// }

// func removeDuplicateValues(intSlice []int) []int {
// 	keys := make(map[int]bool)
// 	list := []int{}

// 	// If the key(values of the slice) is not equal
// 	// to the already present value in new slice (list)
// 	// then we append it. else we jump on another element.
// 	for _, entry := range intSlice {
// 		if _, value := keys[entry]; !value {
// 			keys[entry] = true
// 			list = append(list, entry)
// 		}
// 	}
// 	return list
// }
