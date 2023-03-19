package seed

import (
	"context"
	"log"

	domain "github.com/jwilyandi19/simple-auth/domain/user"
	"github.com/jwilyandi19/simple-auth/utils"
)

type Seed struct {
	userRepo domain.UserRepository
}

func NewSeeder(u domain.UserRepository) Seed {
	return Seed{
		userRepo: u,
	}
}

type UserData struct {
	ID       int
	Username string
	FullName string
	Password string
}

var userData = []UserData{
	{
		ID:       1,
		Username: "abc",
		FullName: "ababab",
		Password: "aaaaa",
	},
	{
		ID:       2,
		Username: "abd",
		FullName: "Jus",
		Password: "bbbbb",
	},
	{
		ID:       3,
		Username: "jwilyandi19",
		FullName: "Jason Wilyandi",
		Password: "password",
	},
}

func (s Seed) Seed() {
	for _, data := range userData {
		password, _ := utils.HashPassword(data.Password)
		_, err := s.userRepo.CreateUser(context.TODO(), domain.CreateUserRequest{
			Username: data.Username,
			FullName: data.FullName,
			Password: password,
		})
		if err != nil {
			log.Println("[Seed] error: ", err.Error())
		}
	}
}
