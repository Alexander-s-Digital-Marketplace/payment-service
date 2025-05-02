package main

import (
	"log"

	"github.com/Alexander-s-Digital-Marketplace/auth-service/internal/database"
	accountmodel "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	resetpasswordmodel "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/reset_password_model"
	rolemodel "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/role_model"
)

func main() {

	var db database.DataBase
	db.InitDB()

	var account accountmodel.UserAccount
	account.MigrateToDB(db)

	var role rolemodel.Role
	role.MigrateToDB(db)

	var resetCode resetpasswordmodel.ResetCode
	resetCode.MigrateToDB(db)

	sqlStatements := []string{
		`INSERT INTO roles (role, role_string) VALUES
        ('adm', 'Администратор'),
        ('mng', 'Менеджер'),
        ('wtr', 'Официант'),
        ('ktn', 'Кухня'),
        ('bar', 'Бар');`,
	}

	for _, stmt := range sqlStatements {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing seed: ", stmt, err)
		}
	}

	log.Println("Success seeding")

	db.CloseDB()
}
