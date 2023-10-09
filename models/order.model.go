package models

import (
	"time"
)

type Status string

type Order struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	Name      string     `gorm:"type:varchar(100);not null"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

type CreateOrderInput struct {
	Name   string `json:"name" validate:"required"`
	UserID uint   `json:"userId" validate:"required"`
}

type OrderResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name,omitempty"`
	Status    Status    `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterOrderRecord(order *Order) OrderResponse {
	return OrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		CreatedAt: *order.CreatedAt,
		UpdatedAt: *order.UpdatedAt,
	}
}
