package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Purchase struct {
	gorm.Model
	Id           int    `json:"id" gorm:"index"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	IsAcceptance bool   `json:"is_acceptance"`
	ProductIds   []int  `json:"product_ids"`
}
