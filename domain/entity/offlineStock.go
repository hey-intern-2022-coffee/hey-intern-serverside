package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OfflineStock struct {
	gorm.Model
	ID            int `json:"id" gorm:"primaryKey"`
	StockQuantity int `json:"stock_quantity"`
	ProductId     int `json:"product_id"`
}
