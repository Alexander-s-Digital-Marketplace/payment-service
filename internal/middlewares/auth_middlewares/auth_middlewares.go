package authmiddlewares

import (
	"net/http"

	validaccesstokenfuncclient "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/services/valid_access_token/valid_access_token_func_client"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
		}
		tokenString = tokenString[len("Bearer "):]

		code, id, role := validaccesstokenfuncclient.ValidAccessToken(tokenString)
		if code == 200 {
			c.Set("id", id)
			c.Set("role", role)
			c.Next()
		} else {
			c.Abort()
		}
	}
}
