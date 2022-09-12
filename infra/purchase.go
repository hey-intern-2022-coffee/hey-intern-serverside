package infra

import (
	"gorm.io/gorm"

	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
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
		return result.Error
	}

	for _, v := range purchase.PurchasesProducts {
		v.PurchaseID = purchase.ID
		if result := tx.Create(&purchase); result.Error != nil {
			return result.Error
		}
	}

	if result := tx.Commit(); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *PurchaseRepository) ToggleIsAcceptance(id int) (*entity.Purchase, error) {
	var purchase *entity.Purchase
	if result := p.DB.Model(&purchase).Where("id = ?", id).Update("is_acceptance", true); result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}

func (p *PurchaseRepository) FindByProductID(id int) (*entity.Purchase, error) {
	var purchase *entity.Purchase
	if result := p.DB.Model(&purchase).Where("id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}
