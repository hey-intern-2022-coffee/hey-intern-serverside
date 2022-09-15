package infra

import (
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"gorm.io/gorm"
)

type PurchaseRepository struct {
	DB *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{
		DB: db,
	}
}

func (p *PurchaseRepository) Insert(purchase *entity.Purchase) error {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if result := tx.Create(&purchase); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	for _, v := range purchase.PurchasesProducts {
		var stock entity.OnlineStock
		if result := tx.First(&stock, "product_id = ?", v.ProductID); result.Error != nil {
			tx.Rollback()
			return result.Error
		}

		stock.SoldQuantity++
		if result := tx.Save(stock); result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return nil
}

func (p *PurchaseRepository) FindAll() ([]entity.Purchase, error) {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var purchases []entity.Purchase
	if result := tx.Find(&purchases); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	for i, purchase := range purchases {
		id := purchase.ID
		var purchasesProducts []entity.PurchasesProducts
		if result := tx.Where("purchase_id = ?", id).Find(&purchasesProducts); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		purchases[i].PurchasesProducts = purchasesProducts
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	return purchases, nil
}

func (p *PurchaseRepository) ToggleIsAcceptance(id int) (*entity.Purchase, error) {
	var purchase entity.Purchase
	if result := p.DB.First(&purchase, "id = ?", id).Update("is_acceptance", true); result.Error != nil {
		return nil, result.Error
	}
	return &purchase, nil
}

func (p *PurchaseRepository) FindByPurchaseID(id int) (*entity.Purchase, error) {
	tx := p.DB.Begin()

	var purchase entity.Purchase
	if result := tx.First(&purchase, "id = ?", id); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	var purchasesProducts []entity.PurchasesProducts
	if result := tx.Find(&purchasesProducts, "purchase_id = ?", purchase.ID); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	purchase.PurchasesProducts = purchasesProducts
	if result := tx.Commit(); result.Error != nil {
		return nil, result.Error
	}

	return &purchase, nil
}
