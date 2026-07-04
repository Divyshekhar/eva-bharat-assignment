package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ticket-system.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database" + err.Error())
		return
	}
	log.Print("Database Connected Successfully")

}
