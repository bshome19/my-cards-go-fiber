package configs

import (
	"github.com/bshome19/my-cards-go-gin-fiber/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func ConnectDB() {
	DB, err := gorm.Open(sqlite.Open("configs/my_cards.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: "+ err.Error())
	}
	DB.AutoMigrate(&models.MyCard{}, &models.User{})
}