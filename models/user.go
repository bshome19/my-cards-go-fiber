package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	UserId string `json:"user_id" gorm:"uniqueIndex"`
	Name string `json:"name"`
	HashPassword string `json:"hash_password"`
}