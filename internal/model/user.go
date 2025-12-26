package model

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Info      UserInfo  `json:"info" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserInfo struct {
	Username string `json:"username" db:"username"`
	Role     Role   `json:"role" db:"role"`
}
