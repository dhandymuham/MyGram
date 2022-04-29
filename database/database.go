package database

import (
	"final/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "koinworks"
	port     = "5432"
	dbname   = "final_project"
	db       *gorm.DB
	err      error
)

func StartDB() {
	database := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatal("Connecting to database error:", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comments{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
