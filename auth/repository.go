package auth

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	// FindByID(ID int) (User, error)
	Create(song User) (User, error)
	SignIn(signin SignIn) (User, error)
	UpdateAddress(user User) (User, error)
	FindByEmail(email string) (User, error)
	// UpdateSong(song User) (User, error)
	// Delete(song User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user User) (User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return user, err
}

func (r *repository) FindAll() ([]User, error) {
	var user []User

	err := r.db.Find(&user).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return user, err
}

func (r *repository) SignIn(signin SignIn) (User, error) {
	var user User
	err := r.db.Debug().Where(&User{Email: signin.Email}).First(&user).Error
	if err != nil {
		println("=====================")
		println("ERROR LOGIN1")
		println("=====================")
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signin.Password))
	// err := r.db.Debug().Where("email = ? AND password = ?", signin.Email, signin.Password).Find(&user).Error
	// err := r.db.Create(&user).Error

	if err != nil {
		println("=====================")
		println("ERROR LOGIN2")
		println("=====================")
	}

	return user, err
}

func (r *repository) UpdateAddress(user User) (User, error) {
	// user_emails := string(user_email)
	// err := r.db.Debug().Where(&User{Email: user_email}).Updates(User{Address: "hello", City: 18, Active: false}).Error

	err := r.db.Save(&user).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE Updating")
		println("=====================")
	}

	return user, err
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE FB-EMAIL")
		println("=====================")
	}

	return user, err
}

// func (r *repository) FindByID(ID int) (User, error) {
// 	var song User

// 	err := r.db.Find(&song, ID).Error
// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE FBI")
// 		println("=====================")
// 	}

// 	return song, err
// }

// func (r *repository) UpdateSong(song User) (User, error) {

// 	err := r.db.Save(&song).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE Updating")
// 		println("=====================")
// 	}

// 	return song, err
// }

// func (r *repository) Delete(song User) (User, error) {

// 	err := r.db.Delete(&song).Error

// 	if err != nil {
// 		println("=====================")
// 		println("ERROR WHILE Deleting")
// 		println("=====================")
// 	}

// 	return song, err
// }
