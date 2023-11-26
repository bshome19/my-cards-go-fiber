package models

import "gorm.io/gorm"


type MyCard struct {
	gorm.Model
	UserId string `gorm:"index" json:"user_id"`
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV int32 `json:"cvv"`
	NameOnCard string `json:"name_on_card"`
}