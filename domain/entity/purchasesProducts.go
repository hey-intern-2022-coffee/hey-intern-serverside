package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PurchasesProducts struct {
	gorm.Model
	ID         int `json:"id" gorm:"primaryKey"`
	ProductID  int `json:"product_id"`
	PurchaseID int `json:"purchase_id"`
}
