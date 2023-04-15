package database

import (
	"fmt"
	"go-final-task/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := "host=containers-us-west-144.railway.app user=postgres password=iPc8czscJAGMJxwXki1W dbname=railway port=5493 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Success connect to DB")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
