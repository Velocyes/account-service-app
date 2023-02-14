package entities

import "time"

type BalanceType struct {
	ID          uint
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	BalanceType string    `json:"balance_type"`
}
