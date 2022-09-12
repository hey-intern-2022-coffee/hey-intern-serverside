package entity

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PurchasesProducts struct {
	ID         int `json:"id" gorm:"primaryKey; autoIncrement"`
	ProductID  int `json:"product_id"`
	PurchaseID int `json:"purchase_id"`
}
