package changepassword

import (
	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	changepasswordform "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/change_password_form"
	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func ChangePasswordHandle(c *gin.Context) (int, string, string, string) {

	var user useraccount.UserAccount
	var form changepasswordform.ChangePasswordForm
	var err error

	err = form.DecodeFromContext(c)
	if err != nil {
		return 400, "", "", "No data input"
	}

	form.HashPassword()

	if form.NewPassword == form.OldPassword {
		return 400, "", "", "Passwords mustn't match"
	}

	user.Email = form.Email
	err = user.GetFromTableByEmail()
	if err != nil {
		logger.Errorln("Incorrect login")
		return 403, "", "", "Incorrect login or password"
	}

	if user.Password != form.OldPassword {
		logger.Errorln("Incorrect password")
		return 403, "", "", "Incorrect login or password"
	}

	user.Password = form.NewPassword

	user.UpdateInTable()

	codeA, accessToken, errAT := jwt.GenerateAccessToken(user)
	if codeA != 200 {
		return codeA, "", "", errAT
	}

	codeR, refreshToken, errRT := jwt.GenerateRefreshToken(user)
	if codeR != 200 {
		return codeR, "", "", errRT
	}

	logger.Infoln("Change password for user is successful. User: ", user.Id, user.Email)

	return 200, accessToken, refreshToken, "Change password is successful"
}
