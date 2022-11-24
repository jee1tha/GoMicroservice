package database

import (
	"GoMicroservice/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to database", Instance.Name())

}

func Migrate() {
	Instance.AutoMigrate(&entities.User{})
	log.Println("Database Migration Completed")
}
