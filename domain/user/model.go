package domain

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"fullname"`
	HashedPassword string `json:"password"`
}

type FetchUserRequest struct {
	Page  int
	Limit int
}

type GetUserRequest struct {
	Username string
	Password string
}

type CreateUserRequest struct {
	Username string
	FullName string
	Password string
}
