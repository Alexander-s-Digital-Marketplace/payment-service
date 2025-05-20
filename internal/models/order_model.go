package models

import (
	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/database"
	"github.com/sirupsen/logrus"
)

type Order struct {
	Id                  int     `gorm:"primaryKey"`
	ContractAddress     string  `gorm:"type:varchar(100);not null"`
	SellerWalletAddress string  `gorm:"type:varchar(100);not null"`
	BuyerWalletAddress  string  `gorm:"type:varchar(100);not null"`
	ProductPrice        float64 `gorm:"not null"`
	TxHash              string  `gorm:"type:varchar(100);not null"`
	IsPaid              bool    `gorm:"type:boolean"`
}

func (order *Order) AddToTable() int {
	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()
	err := db.Connection.Create(&order).Error
	if err != nil {
		logrus.Error("Error add to table: ", err)
		return 503
	}
	return 200
}

func (order *Order) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Order{})
	if err != nil {
		logrus.Errorln("Error migrate Order model :")
		return err
	}
	logrus.Infoln("Success migrate Order model :")
	return nil
}
