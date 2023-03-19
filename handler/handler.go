package handler

import (
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

func NewHandler(user user.UserUsecase) *Handler {
	return &Handler{user}
}

func (h *Handler) LoginHandler(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
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

	arg := domain.CreateUserRequest{
		Username: req.Username,
		FullName: req.FullName,
		Password: req.Password,
	}

	err := h.userUsecase.Register(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "")
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
