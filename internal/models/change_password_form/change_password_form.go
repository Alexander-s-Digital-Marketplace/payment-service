package changepasswordform

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ChangePasswordForm struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// Hash all type of password
func (form *ChangePasswordForm) HashPassword() {
	hasher1 := sha256.New()
	hasher1.Write([]byte(form.NewPassword))
	hashedPassword1 := hex.EncodeToString(hasher1.Sum(nil))
	form.NewPassword = hashedPassword1

	hasher2 := sha256.New()
	hasher2.Write([]byte(form.OldPassword))
	hashedPassword2 := hex.EncodeToString(hasher2.Sum(nil))
	form.OldPassword = hashedPassword2
}

// Decode struct from json gin context
func (form *ChangePasswordForm) DecodeFromContext(c *gin.Context) error {

	if err := c.ShouldBindJSON(&form); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}
