package services

import (
	"dbo/models"
	"dbo/repositories"
)

type CustomerService interface {
	GetAllCustomers(page, limit int) ([]models.Customer, int64, error)
	GetCustomerByID(id uint) (models.Customer, error)
	CreateCustomer(customer models.Customer) (models.Customer, error)
	UpdateCustomer(customer models.Customer) (models.Customer, error)
	DeleteCustomer(id uint) error
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{repo}
}

func (s *customerService) GetAllCustomers(page, limit int) ([]models.Customer, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *customerService) GetCustomerByID(id uint) (models.Customer, error) {
	return s.repo.FindByID(id)
}

func (s *customerService) CreateCustomer(customer models.Customer) (models.Customer, error) {
	return s.repo.Create(customer)
}

func (s *customerService) UpdateCustomer(customer models.Customer) (models.Customer, error) {
	return s.repo.Update(customer)
}

func (s *customerService) DeleteCustomer(id uint) error {
	return s.repo.Delete(id)
}
