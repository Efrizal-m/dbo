package services

import (
	"dbo/models"
	"dbo/repositories"
)

type OrderService interface {
	GetAllOrders(page, limit int) ([]models.Order, int64, error)
	GetOrderByID(id uint) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(id uint) error
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo}
}

func (s *orderService) GetAllOrders(page, limit int) ([]models.Order, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *orderService) GetOrderByID(id uint) (models.Order, error) {
	return s.repo.FindByID(id)
}

func (s *orderService) CreateOrder(order models.Order) (models.Order, error) {
	return s.repo.Create(order)
}

func (s *orderService) UpdateOrder(order models.Order) (models.Order, error) {
	return s.repo.Update(order)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.repo.Delete(id)
}
