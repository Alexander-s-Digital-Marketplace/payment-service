package main

import (
	loggerconfig "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/config/logger"
	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/database"
	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models"
	"github.com/sirupsen/logrus"
)

func main() {
	loggerconfig.Init()

	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	var wallet models.Wallet
	err := wallet.MigrateToDB(db)
	if err != nil {
		logrus.Errorln("Error migrate Wallet model :")
	}

	var order models.Order
	err = order.MigrateToDB(db)
	if err != nil {
		logrus.Errorln("Error migrate Order model :")
	}

	db.CloseDB()
}
