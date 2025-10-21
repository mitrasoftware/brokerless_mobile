package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	utils "github.com/mitrasoftware/pureone_backend_go/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Expect header: Authorization: Bearer <token>
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store mobile in context
		c.Set("mobile", claims.Mobile)

		c.Next()
	}
}
