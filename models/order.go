package models

type Order struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	CustomerID uint    `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	CreatedAt  int64   `json:"created_at"`
}
