package repositories

import (
	"dbo/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll(page, limit int) ([]models.Customer, int64, error)
	FindByID(id uint) (models.Customer, error)
	Create(customer models.Customer) (models.Customer, error)
	Update(customer models.Customer) (models.Customer, error)
	Delete(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) FindAll(page, limit int) ([]models.Customer, int64, error) {
	var customers []models.Customer
	var total int64

	offset := (page - 1) * limit
	if err := r.db.Limit(limit).Offset(offset).Find(&customers).Error; err != nil {
		return nil, 0, err
	}
	r.db.Model(&models.Customer{}).Count(&total)
	return customers, total, nil
}

func (r *customerRepository) FindByID(id uint) (models.Customer, error) {
	var customer models.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customerRepository) Create(customer models.Customer) (models.Customer, error) {
	if err := r.db.Create(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customerRepository) Update(customer models.Customer) (models.Customer, error) {
	if err := r.db.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customerRepository) Delete(id uint) error {
	if err := r.db.Delete(&models.Customer{}, id).Error; err != nil {
		return err
	}
	return nil
}
