package entity

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OfflineStock struct {
	ID            int `json:"id" gorm:"primaryKey; autoIncrement"`
	StockQuantity int `json:"stock_quantity"`
	ProductId     int `json:"product_id"`
}
