package middlewares

import (
	"blogo/app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	_, err := utils.ValidateJWT(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
}
