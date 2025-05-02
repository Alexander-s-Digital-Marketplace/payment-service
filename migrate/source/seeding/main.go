package main

import (
	"log"

	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	sqlStatements := []string{
		`INSERT INTO user_accounts (email, password, role) VALUES
        ('user1@example.com', 'password123', 'adm'),
        ('mod1@example.com', 'password123', 'mng'),
        ('admin1@example.com', 'password123', 'wtr'),
        ('user2@example.com', 'password456', 'ktn'),
        ('mod2@example.com', 'password456', 'bar');`,
	}

	for _, stmt := range sqlStatements {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing seed: ", stmt, err)
		}
	}

	log.Println("Success seeding")

	db.CloseDB()
}
