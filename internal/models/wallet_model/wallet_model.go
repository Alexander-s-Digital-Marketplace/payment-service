package walletmodel

import (
	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/database"
	"github.com/sirupsen/logrus"
)

type Wallet struct {
	Id            int    `gorm:"primaryKey"`
	WalletAddress string `gorm:"type:varchar(100);unique;not null"`
}

func (w *Wallet) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Create(&w).Error
	if err != nil {
		logrus.Error("Error add to table: ", err)
		return 503
	}

	return 200
}

func (w *Wallet) GetFromTableById() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.First(&w).Error
	if err != nil {
		return 503
	}

	return 200
}

func (w *Wallet) DeleteFromTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	err := db.Connection.Delete(&w).Error
	if err != nil {
		logrus.Error("Error deleting wallet: ", err)
		return 503
	}
	return 200
}

func (w *Wallet) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Wallet{})
	if err != nil {
		logrus.Errorln("Error migrate Wallet model :")
		return err
	}
	logrus.Infoln("Success migrate Wallet model :")
	return nil
}
