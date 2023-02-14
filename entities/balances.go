package entities

import (
	"time"
)

type Balance struct {
	ID           uint
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UserId       int       `json:"user_id"`
	TotalBalance float64   `json:"total_balance"`
}
