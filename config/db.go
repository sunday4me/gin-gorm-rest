package config

import (
	"github.com/sunday4me/gin-gorm-rest/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


var DB *gorm.DB
func Connect () {
	db, err := gorm.Open(postgres.Open("postgres:postgres@localhost:5432/postgres"),&gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
	//defer db.Close()
}