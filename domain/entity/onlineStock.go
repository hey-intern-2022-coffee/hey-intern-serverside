package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OnlineStock struct {
	gorm.Model
	Id            int `json:"id" gorm:"index"`
	StockQuantity int `json:"stock_quantity"`
	ProductId     int `json:"product_id"`
}
