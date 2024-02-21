package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    *uuid.UUID `gorm:"type:uuid;"`
	ProductID *uuid.UUID `gorm:"type:uuid;"`
	Quantity  int        `gorm:"not null;default:1"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

type CartResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
	ProductID uuid.UUID `json:"product_id,omitempty"`
	Quantity  int       `json:"quantity,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CartItemInput struct {
	UserID    *uuid.UUID `gorm:"type:uuid;"`
	ProductID *uuid.UUID `gorm:"type:uuid;"`
	Quantity  int        `gorm:"not null;default:1"`
}

type CartCheckoutInput struct {
	UserID string `json:"user_id"`
	Money  int    `json:"money"`
}

func FilterCartRecord(cart *Cart) CartResponse {
	return CartResponse{
		ID:        *cart.ID,
		UserID:    *cart.UserID,
		ProductID: *cart.ProductID,
		Quantity:  cart.Quantity,
		CreatedAt: *cart.CreatedAt,
		UpdatedAt: *cart.UpdatedAt,
	}
}
