package login

import (
	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// User's login
func LoginHandle(c *gin.Context) (int, string, string, string) {
	var userFront useraccount.UserAccount
	var userDB useraccount.UserAccount

	err := userFront.DecodeFromContext(c)
	if err != nil {
		logrus.Errorln("Empty field")
		return 400, "", "", "Empty field"
	}

	userFront.SetPasswordHash(userFront.Password)
	userDB.Email = userFront.Email

	err = userDB.GetFromTableByEmail()
	if err != nil {
		logrus.Errorln("Incorrect login")
		return 403, "", "", "Incorrect login or password"
	}

	if userDB.Password != userFront.Password {
		logrus.Errorln("Incorrect password")
		return 403, "", "", "Incorrect login or password"
	}

	codeA, accessToken, errAT := jwt.GenerateAccessToken(userDB)
	if codeA != 200 {
		return codeA, "", "", errAT
	}

	codeR, refreshToken, errRT := jwt.GenerateRefreshToken(userDB)
	if codeR != 200 {
		return codeR, "", "", errRT
	}

	logrus.Infoln("Authorization is successful. User: ", userDB.Id, userDB.Email)

	return 200, accessToken, refreshToken, "Authorization is successful"
}
