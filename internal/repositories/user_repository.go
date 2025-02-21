package repositories

import (
	"github.com/HarAlb/go-project/internal/entities"
)

var users = []entities.User{}

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetUserByID(id int) *entities.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (u *UserRepository) CreateUser(name string, password string, email string) *entities.User {

	index := 0

	if len(users) > 0 {
		index = users[len(users)-1].ID + 1
	}else {
		index = 1
	}

	user := entities.User{
		ID: index,
		Username: name,
		Password: password,
		Email: email,
	}
	users = append(users, user)

	return &user
}

func (u *UserRepository) All() []entities.User {
	return users
}