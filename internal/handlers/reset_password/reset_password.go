package resetpassword

import (
	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	resetpasswordcode "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/reset_password_model"
	notificationservice "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/services/notification_service/notificaton_service_client"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// User's reset password
func ResetPasswordHandle(c *gin.Context) (int, string) {

	var userFront useraccount.UserAccount
	var userDB useraccount.UserAccount
	var err error
	var code int
	var message string

	userFront.DecodeFromContext(c)
	userDB.Email = userFront.Email

	err = userDB.GetFromTableByEmail()
	if err != nil {
		logrus.Errorln("Incorrect login")
		return 403, "Incorrect login or email"
	}

	if userDB.Email != userFront.Email {
		logrus.Errorln("Incorrect email")
		return 403, "Incorrect login or email"
	}

	var resetForm resetpasswordcode.ResetCode

	resetForm.GenerateCode()
	resetForm.InitDate(5)

	code, message = notificationservice.ResetNotif(resetForm, userDB.Email)
	if code != 200 {
		return code, message
	}

	resetForm.User = userDB
	resetForm.AddToTable()

	logrus.Infoln("Send reset email is successful. User: ", userDB.Id, userDB.Email,
		" EMAIL: ", userDB.Email, "code: ", resetForm.Code, "time: ", resetForm.StartTime, " ", resetForm.ExpTime)

	return 200, "Send reset email is successful"
}
