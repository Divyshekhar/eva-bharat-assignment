package config

import (
	"log"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Tickets{},
	)
	if err != nil {
		log.Fatal("Could not migrate the models to database" + err.Error())
		return
	}
	log.Print("Database Migrated Successfully")
}
