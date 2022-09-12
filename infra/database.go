package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.Purchase{}, &entity.Product{}, &entity.OnlineStock{}, &entity.OfflineStock{}, &entity.PurchasesProducts{})
	return db, nil
}