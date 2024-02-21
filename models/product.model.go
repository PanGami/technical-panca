package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID          *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	Price       float64    `gorm:"not null"`
	Stock       int        `gorm:"not null"`
	CreatedAt   *time.Time `gorm:"not null;default:now()"`
	UpdatedAt   *time.Time `gorm:"not null;default:now()"`
}

type CreateProductInput struct {
	Name        string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
}

// Can be nullable because its for update method
type UpdateProductInput struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}
