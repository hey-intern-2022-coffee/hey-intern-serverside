package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Purchase struct {
	gorm.Model
	ID           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	IsAcceptance bool   `json:"is_acceptance"`
}
