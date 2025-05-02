package main

import (
	"log"

	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	query := []string{
		`DROP TABLE user_accounts CASCADE;`,
		`DROP TABLE roles CASCADE;`,
		`DROP TABLE reset_codes CASCADE;`,
	}

	for _, stmt := range query {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing drop: ", stmt, err)
		}
	}

	log.Println("All table is droped")

	db.CloseDB()
}
