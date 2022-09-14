package infra

import (
	"errors"

	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) Insert(product *entity.Product) error {
	product.OnlineStock.DeliveredQuantity = product.OnlineStock.StockQuantity
	if result := p.DB.Create(&product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *ProductRepository) Update(product *entity.Product) error {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if result := tx.Save(product); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	var stock entity.OnlineStock
	if result := tx.First(&stock, "product_id = ?", product.ID).Update("stock_quantity", product.OnlineStock.StockQuantity); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	product.OnlineStock = stock

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return nil
}

func (p *ProductRepository) Delete(id int) error {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if result := tx.Delete(&entity.OnlineStock{}, "product_id = ?", id); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result := tx.Delete(&entity.Product{}, "id = ?", id); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return nil
}

func (p *ProductRepository) FindAll() ([]entity.Product, error) {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var products []entity.Product
	if result := tx.Find(&products); result.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}

	for i, product := range products {
		id := product.ID
		var stock entity.OnlineStock
		if result := tx.First(&stock, "product_id = ?", id); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		products[i].OnlineStock = stock
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	return products, nil
}

func (p *ProductRepository) FindIdOne(id int) (*entity.Product, error) {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var stock entity.OnlineStock
	if result := tx.First(&stock, "product_id = ?", id); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	var product entity.Product
	if result := tx.First(&product, "id = ?", id); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	product.OnlineStock = stock
	return &product, nil
}

func (p *ProductRepository) PatchPurchase(id int) (*entity.Product, error) {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var stock entity.OnlineStock
	if result := tx.First(&stock, "product_id = ?", id); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if stock.DeliveredQuantity < 1 {
		tx.Rollback()
		return nil, errors.New("stock is 0")
	}

	stock.DeliveredQuantity--
	if result := tx.Model(&stock).Update("delivered_quantity", stock.DeliveredQuantity); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	var product entity.Product
	if result := tx.First(&product, "id = ?", id); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	product.OnlineStock = stock
	return &product, nil
}
