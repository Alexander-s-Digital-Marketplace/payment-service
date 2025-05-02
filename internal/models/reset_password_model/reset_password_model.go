package resetpasswordmodel

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/database"
	account "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResetCode struct {
	Id        int                 `gorm:"primaryKey;autoIncrement"`
	Id_user   int                 `gorm:"not null"`
	Code      int                 `gorm:"not null;type:integer"`
	StartTime string              `gorm:"type:varchar(50)"`
	ExpTime   string              `gorm:"type:varchar(50)"`
	User      account.UserAccount `gorm:"foreignKey:Id_user;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Decode struct from json gin context
func (resetForm *ResetCode) DecodeFromContext(c *gin.Context) error {

	if err := c.ShouldBindJSON(&resetForm); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

// Generate reset code
func (resetForm *ResetCode) GenerateCode() error {
	min := int64(100000)
	max := int64(999999)

	randNum, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		logrus.Error("Error generate reset code: ", err)
		return err
	}

	resetForm.Code = int(randNum.Int64() + min)

	return nil
}

// Init date for code, where param is duration in mitune
func (resetForm *ResetCode) InitDate(duration int) {

	currentTime := time.Now()
	expirationTime := currentTime.Add(time.Duration(duration) * time.Minute)

	resetForm.StartTime = currentTime.Format("2006-01-02 15:04:05")
	resetForm.ExpTime = expirationTime.Format("2006-01-02 15:04:05")
}

// Create new row in code
func (resetForm *ResetCode) AddToTable() int {
	var db database.DataBase
	db.InitDB()

	resetForm.Code = int(resetForm.Code)
	err := db.Connection.Create(&resetForm).Error
	if err != nil {
		db.CloseDB()
		return 503
	}
	db.CloseDB()
	return 0
}

// Get coderecord from table by code
func (resetForm *ResetCode) GetFromTableByCode() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.First(&resetForm, "code = ?", int(resetForm.Code)).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

// Get coderecord from table by user id
func (resetForm *ResetCode) GetFromTableByUserId() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.First(&resetForm, "id_user = ?", resetForm.Id_user).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

// Remove coderecord from table by code
func (resetForm *ResetCode) DeleteFromTableByCode() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.Delete(&resetForm, "code = ?", resetForm.Code).Error
	if err != nil {
		db.CloseDB()
		logrus.Error("Error deleting reset code: ", err)
		return err
	}
	db.CloseDB()
	return nil
}

func (resetForm *ResetCode) ValideCode() (int, string) {
	dateExp, _ := time.Parse("2006-01-02 15:04:05", resetForm.ExpTime)
	dateNow := time.Now()
	if dateExp.Before(dateNow) {
		return 403, "Code is not valide"
	}
	return 200, ""
}

func (resetForm *ResetCode) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&ResetCode{})
	if err != nil {
		logrus.Errorln("Error migrate ResetCode model :")
		return err
	}
	logrus.Infoln("Success migrate ResetCode model :")
	return nil
}
