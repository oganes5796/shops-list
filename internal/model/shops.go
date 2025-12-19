package model

import "time"

type Shop struct {
	ID        int64     `json:"-" db:"id"`
	Info      ShopInfo  `json:"info" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ShopInfo struct {
	Title         string `json:"title" db:"title"`
	Address       string `json:"address" db:"address"`
	OperatingMode string `json:"operating_mode" db:"operating_mode"`
}
