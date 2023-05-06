package model

import (
	"bot/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(postgres.Open(helpers.Env.PostgresUri), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return
	}
	log.Println("🔥 Database Connected 🔥")
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
		return
	}
	log.Println("🔥 Database Migrated 🔥")
	DB = db
	return
}
