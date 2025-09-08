package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			ctx.Abort()
			return
		}
		authHeader = strings.TrimPrefix(authHeader, "Bearer")
		claims, err := ValidateToken(authHeader)

		if err != nil {
			log.Printf("Token validation error: %v", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
