package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domain "github.com/jwilyandi19/simple-auth/domain/user"
	"github.com/jwilyandi19/simple-auth/usecase/user"
)

type Handler struct {
	userUsecase user.UserUsecase
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

var (
	ErrDataExist = fmt.Errorf("data exist")
)

func NewHandler(user user.UserUsecase) *Handler {
	return &Handler{user}
}

func (h *Handler) LoginHandler(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := validate(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	token, err := h.userUsecase.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) SignUpHandler(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := validate(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := domain.CreateUserRequest{
		Username: req.Username,
		FullName: req.FullName,
		Password: req.Password,
	}

	token, err := h.userUsecase.Register(ctx, arg)
	if err == ErrDataExist {
		ctx.JSON(http.StatusConflict, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) UserListHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	arg := domain.FetchUserRequest{
		Page:  page,
		Limit: limit,
	}

	users, err := h.userUsecase.FetchUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	datas := make([]UserResponse, 0)
	for _, user := range users {
		datas = append(datas, UserResponse{
			ID:       user.ID,
			Username: user.Username,
			FullName: user.FullName,
		})
	}

	ctx.JSON(http.StatusOK, datas)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func validate(username string, password string) error {
	if username == "" {
		return errors.New("username required")
	}

	if len(username) < 2 {
		return errors.New("username must more than 2 character")
	}

	if password == "" {
		return errors.New("password required")
	}

	if len(password) < 5 {
		return errors.New("password most more than 5 character")
	}

	return nil
}
