package registration

import (
	"net/http"

	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	profileregister "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/services/profile_register_service/profile_register_service_client"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// User's register
func RegistrationHandle(c *gin.Context) (int, string) {
	var err error
	var code int
	var profile useraccount.ProfileTDO
	err = profile.DecodeFromContext(c)
	if err != nil {
		return 400, string(err.Error())
	}
	logrus.Infoln("profile: ", profile)

	var user useraccount.UserAccount
	user.Email = profile.Email
	user.Password = profile.Password
	user.SetPasswordHash(user.Password)

	if user.Email == "" || user.Password == "" {
		logrus.Errorln("Field is empty")
		return 400, "Field is empty"
	}

	code = user.AddToTable()
	if code == 409 {
		logrus.Errorln("With user is already exist")
		return 400, "With user is already exist"
	}
	if code == 503 {
		logrus.Errorln("Not avalible")
		return 503, "Not avalible"
	}
	profile.AccountInfoId = user.Id

	var message string
	code, message = profileregister.ProfileRegister(profile)
	if code != 200 {
		logrus.Errorln(message)
		return code, message
	}
	logrus.Infoln("Registration of new user is successful. User: ", user.Id, user.Email)
	return http.StatusOK, "Registration of new user is successful"

}
