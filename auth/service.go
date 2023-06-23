package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindAll() ([]User, error)
	// FindByID(ID int) (Song, error)
	FindByEmail(email string) (User, error)
	Create(userInput UserInput) (User, error)
	SignIn(signin SignIn) (User, error)
	// UpdateSong(ID int, songInput SongInput) (Song, error)
	// Delete(ID int) (Song, error)
	UpdateAddress(addressInput AddressInput, user_email string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(userInput UserInput) (User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
	user := User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hash),
		Phone:    userInput.Phone,
		Type:     "buyer",
		Balance:  0,
		// Type:     userInput.Type,
		// Title:       bookInput.Title,
		// Price:       bookInput.Price,
		// Description: bookInput.Description,
		// Rating:      bookInput.Rating,
		// Discount:    bookInput.Discount,
	}
	newuser, err := s.repository.Create(user)
	return newuser, err
}

func (s *service) SignIn(signin SignIn) (User, error) {

	user, err := s.repository.SignIn(signin)
	return user, err
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err

}

func (s *service) UpdateAddress(addressInput AddressInput, user_email string) (User, error) {

	user, err := s.repository.FindByEmail(user_email)
	if err != nil {
		fmt.Println(err)
	}
	user.City = addressInput.City
	user.Address = addressInput.Address
	// Title:       bookInput.Title,
	// Price:       bookInput.Price,
	// Description: bookInput.Description,
	// Rating:      bookInput.Rating,
	// Discount:    bookInput.Discount,

	newaddress, err := s.repository.UpdateAddress(user)
	return newaddress, err
}

func (s *service) FindByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	return user, err
}

// func (s *service) FindByID(ID int) (Song, error) {
// 	song, err := s.repository.FindByID(ID)
// 	return song, err
// }

// func (s *service) UpdateSong(ID int, songInput SongInput) (Song, error) {

// 	song, err := s.repository.FindByID(ID)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	song.YtID = songInput.YtID
// 	// book.Title = bookInput.Title
// 	// book.Price = bookInput.Price
// 	// book.Description = bookInput.Description
// 	// book.Rating = bookInput.Rating
// 	// book.Discount = bookInput.Discount

// 	newsong, err := s.repository.UpdateSong(song)
// 	return newsong, err
// }

// func (s *service) Delete(ID int) (Song, error) {
// 	song, err := s.repository.FindByID(ID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	deleteSong, err := s.repository.Delete(song)

// 	return deleteSong, err
// }
