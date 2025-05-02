package api

import (
	changepassword "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/change_password"
	getallroles "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/get_all_roles"
	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/login"
	refreshtoken "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/refresh_token"
	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/registration"
	resetpassword "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/reset_password"
	validaccesstoken "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/valid_access_token"
	verefyresetcode "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/handlers/verefy_reset_code"
	"github.com/gin-gonic/gin"
)

type DefaultAPI struct {
}

// Post /register
// New user's registration
func (api *DefaultAPI) Register(c *gin.Context) {

	code, message := registration.RegistrationHandle(c)

	c.JSON(code, gin.H{"message": message})
}

// Post /login
// User's login
func (api *DefaultAPI) Login(c *gin.Context) {

	code, accessToken, refreshToken, message := login.LoginHandle(c)

	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})

}

// Post /refresh-token
// Refresh access token
func (api *DefaultAPI) RefreshToken(c *gin.Context) {

	code, accessToken, refreshToken, message := refreshtoken.RefreshTokenHandle(c)
	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})
}

// Post /change-password
// Change user's password
func (api *DefaultAPI) ChangePassword(c *gin.Context) {

	code, accessToken, refreshToken, message := changepassword.ChangePasswordHandle(c)

	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})
}

// Post /reset-password
// Reset user's password
func (api *DefaultAPI) ResetPassword(c *gin.Context) {

	code, message := resetpassword.ResetPasswordHandle(c)

	c.JSON(code, gin.H{"message": message})
}

// Post /verefy-reset-code
// Verefy code for reset user's password
func (api *DefaultAPI) VerefyRecetCode(c *gin.Context) {

	code, accessToken, refreshToken, message := verefyresetcode.VerefyResetCodeHandle(c)

	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})
}

func (api *DefaultAPI) GetAllRoles(c *gin.Context) {

	code, roles := getallroles.GetAllRoles(c)

	c.JSON(code, roles)
}

func (api *DefaultAPI) ValidAccessToken(c *gin.Context) {

	code, claim := validaccesstoken.ValidAccessToken(c)

	c.JSON(code, claim)
}
