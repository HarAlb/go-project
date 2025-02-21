package api

import (
	"fmt"

	"github.com/HarAlb/go-project/internal/repositories"
	"github.com/icrowley/fake"
)

type Api struct {
}

func NewAPI() *Api {
	return &Api{}
}

func (a *Api) Start() {
	var num int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	
	userRepo := repositories.NewUserRepository();

	for i := 0; i < num; i++ {
		fmt.Println(userRepo.CreateUser(fake.FirstName(), fake.CharactersN(10), fake.EmailAddress()))
	}

	fmt.Println("Creating", num, "users")
}
