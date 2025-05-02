package verefyresetcode

import (
	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	resetpasswordcode "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/reset_password_model"
	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func VerefyResetCodeHandle(c *gin.Context) (int, string, string, string) {

	var resetFormFront resetpasswordcode.ResetCode
	var resetFormDB resetpasswordcode.ResetCode
	var user useraccount.UserAccount
	var userFront useraccount.UserAccount
	var err error

	resetFormFront.DecodeFromContext(c)
	userFront = resetFormFront.User
	user = resetFormFront.User

	err = user.GetFromTableByEmail()
	if err != nil {
		return 404, "", "", "User not found"
	}

	resetFormDB.Id_user = user.Id

	err = resetFormDB.GetFromTableByUserId()
	if err != nil {
		logrus.Errorln("No reset request")
		return 400, "", "", "No reset request"
	}

	if resetFormDB.Code != resetFormFront.Code {
		logrus.Errorln("Incorrect reset code")
		return 403, "", "", "Incorrect reset code"
	}

	code, errValide := resetFormDB.ValideCode()
	if code != 200 {
		return code, "", "", errValide
	}

	codeA, accessToken, errAT := jwt.GenerateAccessToken(resetFormDB.User)
	if codeA != 200 {
		return codeA, "", "", errAT
	}

	codeR, refreshToken, errRT := jwt.GenerateRefreshToken(resetFormDB.User)
	if codeR != 200 {
		return codeR, "", "", errRT
	}

	resetFormDB.DeleteFromTableByCode()

	user.SetPasswordHash(userFront.Password)

	user.UpdateInTable()

	return 200, accessToken, refreshToken, "Reset password is success"

}
