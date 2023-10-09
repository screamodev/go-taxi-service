package models

import (
	"time"
)

type Driver struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100);not null"`
	Orders    []Order
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

type CreateDriverInput struct {
	Name string `json:"name" validate:"required"`
}

type DriverResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterDriverRecord(driver *Driver) DriverResponse {
	return DriverResponse{
		ID:        driver.ID,
		Name:      driver.Name,
		CreatedAt: *driver.CreatedAt,
		UpdatedAt: *driver.UpdatedAt,
	}
}
