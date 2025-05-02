package authmiddlewares

import (
	validaccesstoken "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/valid_access_token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		code, claims := validaccesstoken.ValidAccessToken(c)
		if code == 200 {
			c.Set("id", claims["id"])
			c.Set("role", claims["role"])
			c.Next()
		} else {
			c.Abort()
		}
	}
}
