package user

import (
	"context"
	"log"

	domain "github.com/jwilyandi19/simple-auth/domain/user"
	"github.com/jwilyandi19/simple-auth/utils"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

type UserUsecase interface {
	FetchUsers(ctx context.Context, req domain.FetchUserRequest) ([]domain.User, error)
	Login(ctx context.Context, username string, password string) (string, error)
	Register(ctx context.Context, req domain.CreateUserRequest) (string, error)
}

func NewUserUsecase(u domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) FetchUsers(ctx context.Context, req domain.FetchUserRequest) ([]domain.User, error) {
	users, err := u.userRepo.FetchUser(ctx, req)
	if err != nil {
		log.Println("[UserUsecase-FetchUsers] error to get list of users: ", err.Error())
		return []domain.User{}, err
	}
	return users, nil
}

func (u *userUsecase) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := u.userRepo.GetUser(ctx, username)
	if err != nil {
		log.Println("[UserUsecase-Login] error to get user: ", err.Error())
		return "", err
	}
	check := utils.CheckPasswordHash(password, user.HashedPassword)
	if !check {
		log.Println("[UserUsecase-Login] unauthorized: ", err.Error())
		return "", nil
	}

	token, err := utils.GenerateToken(username)
	if err != nil {
		log.Println("[UserUsecase-Login] error to call bcrypt: ", err.Error())
		return "", err
	}
	return token, nil
}

func (u *userUsecase) Register(ctx context.Context, req domain.CreateUserRequest) (string, error) {
	exist, err := u.userRepo.IsUserExist(ctx, req.Username)
	if err != nil {
		log.Println("[UserUsecase-Register] error to check exist: ", err.Error())
		return "", err
	}

	if exist {
		return "", ErrDataExist
	}

	cryptPass, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Println("[UserUsecase-Register] error to crypt: ", err.Error())
		return "", err
	}
	req.Password = cryptPass
	user, err := u.userRepo.CreateUser(ctx, req)
	if err != nil {
		log.Println("[UserUsecase-Register] error: ", err.Error())
		return "", err
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		log.Println("[UserUsecase-Register] error to call bcrypt: ", err.Error())
		return "", err
	}

	return token, nil
}
