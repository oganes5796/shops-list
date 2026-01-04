package model

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusCreated    OrderStatus = "created"
	OrderStatusInProgress OrderStatus = "in_progress"
	OrderStatusSuccess    OrderStatus = "success"
	OrderStatusFail       OrderStatus = "fail"
)

type CartItem struct {
	SKU      string  `json:"sku" db:"sku"`
	Quantity int     `json:"quantity" db:"quantity"`
	Price    float64 `json:"price" db:"price"`
}

type CreateOrderRequest struct {
	UserID int        `json:"user_id" db:"user_id"`
	ShopID int        `json:"shop_id" db:"shop_id"`
	Cart   []CartItem `json:"cart" db:"-"`
}

type UpdateOrderRequest struct {
	Status *OrderStatus `json:"status,omitempty" db:"status"`
}

type OrderInfo struct {
	ID uuid.UUID `json:"id" db:"id"`

	UserID int `json:"user_id" db:"user_id"`
	ShopID int `json:"shop_id" db:"shop_id"`

	Cart []CartItem `json:"cart" db:"-"`

	SummaryVolume int     `json:"summary_volume" db:"summary_volume"`
	Price         float64 `json:"price" db:"price"`

	Status OrderStatus `json:"status" db:"status"`
}

type Order struct {
	Info OrderInfo `json:"info" db:"-"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
