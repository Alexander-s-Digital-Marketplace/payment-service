package email

import (
	"strconv"

	emailconfig "github.com/Alexander-s-Digital-Marketplace/auth-service/.env/email"
	"github.com/sirupsen/logrus"

	"gopkg.in/gomail.v2"
)

func SendEmail(adressTo string, username string, code int) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailconfig.Email) // Адрес отправителя
	m.SetHeader("To", adressTo)            // Адрес получателя
	m.SetHeader("Subject", "СБРОС ПАРОЛЯ Catering Auth Service!")

	message := "Код для пользователя: " + username + ": " + strconv.Itoa(code)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(emailconfig.Host, emailconfig.Port, emailconfig.Email, emailconfig.Password)

	err := d.DialAndSend(m)
	if err != nil {
		logrus.Error("Error send email: ", err)
	}

	return err
}
