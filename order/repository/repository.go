package repository

import (
	"kode-golang-section7/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(order *domain.Order) (*domain.Order, error) {
	err := r.db.Create(order).Error
	if err != nil {
		return &domain.Order{}, err
	}
	return order, nil
}

func (r *repository) FindAll() ([]*domain.Order, error) {
	var orders []*domain.Order
	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (r *repository) Update(order *domain.Order) (*domain.Order, error) {
	err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository) Delete(order *domain.Order) (bool, error) {
	if err := r.db.Delete(order).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindByID(id int) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Where("order_id = ?", id).Find(&order).Error
	if err != nil {
		return &order, err
	}
	return &order, nil
}
