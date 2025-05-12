package main

import (
	loggerconfig "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/config/logger"
	"github.com/Alexander-s-Digital-Marketplace/payment-service/internal/database"
	walletmodel "github.com/Alexander-s-Digital-Marketplace/payment-service/internal/models/wallet_model"
	"github.com/sirupsen/logrus"
)

func main() {
	loggerconfig.Init()

	var db database.DataBase
	db.InitDB()
	defer db.CloseDB()

	var wallet walletmodel.Wallet
	err := wallet.MigrateToDB(db)
	if err != nil {
		logrus.Errorln("Error migrate Wallet model :")
	}

	db.CloseDB()
}
