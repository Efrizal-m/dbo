package repositories

import (
	"dbo/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll(page, limit int) ([]models.Order, int64, error)
	FindByID(id uint) (models.Order, error)
	Create(order models.Order) (models.Order, error)
	Update(order models.Order) (models.Order, error)
	Delete(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) FindAll(page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit
	if err := r.db.Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	r.db.Model(&models.Order{}).Count(&total)
	return orders, total, nil
}

func (r *orderRepository) FindByID(id uint) (models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (r *orderRepository) Create(order models.Order) (models.Order, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (r *orderRepository) Update(order models.Order) (models.Order, error) {
	if err := r.db.Save(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (r *orderRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}
