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

	db.Migrator().DropTable(&entity.Purchase{}, &entity.Product{}, &entity.OnlineStock{}, &entity.OfflineStock{}, &entity.PurchasesProducts{})
	db.AutoMigrate(&entity.Purchase{}, &entity.Product{}, &entity.OnlineStock{}, &entity.OfflineStock{}, &entity.PurchasesProducts{})

	product := entity.Product{
		Name:     "イベントうちわ",
		Price:    1000,
		ImageURL: "https://cdn-novelty.raksul.com/public_images/2f06a2ee-cb1e-4416-b3c9-eef52688162a",
	}
	db.Create(&product)

	onlineStock := entity.OnlineStock{
		StockQuantity: 10,
		DeliveredQuantity: 10,
		ProductId:     product.ID,
	}
	db.Create(&onlineStock)

	purchase := entity.Purchase{
		Name:        "田中 大貴",
		Address:     "	高知県四万十市西土佐中半9-7-7",
		MailAddress: "real@gmail.com",
	}
	db.Create(&purchase)

	purchasesProducts := entity.PurchasesProducts{
		ProductID:  product.ID,
		PurchaseID: purchase.ID,
	}
	db.Create(&purchasesProducts)

	return db, nil
}
