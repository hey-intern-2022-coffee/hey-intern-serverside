package entity

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OnlineStock struct {
	ID            int `json:"id" gorm:"primaryKey; autoIncrement"`
	StockQuantity int `json:"stock_quantity"`
	SoldQuantity int `json:"sold_quantity"`
	ProductId     int `json:"product_id"`
}
