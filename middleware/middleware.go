package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jwilyandi19/simple-auth/utils"
)

func AuthorizationMiddleware(ctx *gin.Context) {
	s := ctx.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := utils.ValidateToken(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
