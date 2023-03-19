package domain

import "context"

type UserRepository interface {
	FetchUser(ctx context.Context, req FetchUserRequest) ([]User, error)
	GetUser(ctx context.Context, username string) (User, error)
	CreateUser(ctx context.Context, req CreateUserRequest) (User, error)
	IsUserExist(ctx context.Context, username string) (bool, error)
}
