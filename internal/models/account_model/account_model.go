package accountmodel

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoleEnum string

const (
	ADM RoleEnum = "adm"
	MDR RoleEnum = "mdr"
)

type ProfileTDO struct {
	UserName      string `json:"user_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	AccountInfoId int
}

func (profile *ProfileTDO) DecodeFromContext(c *gin.Context) error {

	if err := c.ShouldBindJSON(&profile); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

type UserAccount struct {
	Id       int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string   `json:"email" gorm:"type:varchar(50)"`
	Password string   `json:"password" gorm:"type:varchar(100)"`
	Role     RoleEnum `json:"role" gorm:"type:varchar(5)"`
}

// Set password as hash
func (user *UserAccount) SetPasswordHash(passwordIN string) {
	hasher := sha256.New()
	hasher.Write([]byte(passwordIN))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	user.Password = hashedPassword
}

// Decode struct from json gin context
func (user *UserAccount) DecodeFromContext(c *gin.Context) error {

	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

// Create new row in user's account
func (user *UserAccount) AddToTable() int {
	var db database.DataBase
	db.InitDB()

	userFind := &UserAccount{}
	userFind.Email = user.Email

	errFind := userFind.GetFromTableByEmail()
	if errFind == nil {
		db.CloseDB()
		logrus.Println("userFind ", userFind)
		return 409
	}

	err := db.Connection.Create(&user).Error
	if err != nil {
		logrus.Println(err)
		logrus.Println(user)
		db.CloseDB()
		return 503
	}
	db.CloseDB()
	return 0
}

// Get user from table by username
func (user *UserAccount) GetFromTableByEmail() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.First(&user, "email = ?", user.Email).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

// Get user from table by id
func (user *UserAccount) GetFromTable() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.First(&user).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

// Update user in table by id
func (user *UserAccount) UpdateInTable() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.Save(&user).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

func (user *UserAccount) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&UserAccount{})
	if err != nil {
		logrus.Errorln("Error migrate UserAccount model :")
		return err
	}
	logrus.Infoln("Success migrate UserAccount model :")
	return nil
}
