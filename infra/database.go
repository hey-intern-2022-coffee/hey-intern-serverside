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
		Name: "test",
		Price: 100,
		ImageURL: "none",
	}
	db.Create(&product)

	onlineStock := entity.OnlineStock{
		StockQuantity: 10,
		ProductId: product.ID,
	}
	db.Create(&onlineStock)

	purchase := entity.Purchase{
		Name: "Purchase",
		Address: "Address",
		MailAddress: "Mail",
	}
	db.Create(&purchase)

	purchasesProducts := entity.PurchasesProducts {
		ProductID: product.ID,
		PurchaseID: purchase.ID,
	}
	db.Create(&purchasesProducts)

	return db, nil
}