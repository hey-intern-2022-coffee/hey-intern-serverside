package entity

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Purchase struct {
	ID                int                 `json:"id" gorm:"primaryKey; autoIncrement"`
	Name              string              `json:"name"`
	Address           string              `json:"address"`
	MailAddress       string              `json:"mail_address"`
	IsAcceptance      bool                `json:"is_acceptance"`
	PurchasesProducts []PurchasesProducts `json:"purchases_products"`
}
