package entity

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	ID          int         `json:"id" gorm:"primaryKey; autoIncrement"`
	Name        string      `json:"name"`
	Price       int         `json:"price"`
	ImageURL    string      `json:"image_url"`
	OnlineStock OnlineStock `json:"online_stock"`
}
