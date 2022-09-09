package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Id       int    `json:"id" gorm:"index"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageURL string `json:"image_url"`
}
