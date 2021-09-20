package domain

import (
	"time"
)

type Order struct {
	OrderID      int       `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderdAt     time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:order_id;references:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}

type OrderRepository interface {
	Create(order *Order) (*Order, error)
	Update(order *Order) (*Order, error)
	FindAll() ([]*Order, error)
	FindByID(id int) (*Order, error)
	Delete(order *Order) (bool, error)
}

type OrderService interface {
	Create(order *Order) (*Order, error)
	Update(order *Order) (*Order, error)
	FindAll() ([]*Order, error)
	Delete(ID int) (bool, error)
}
