package entities

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserId       int     `json:"user_id"`
	User         *User   `gorm:"foreignKey:UserId"`
	TotalBalance float64 `json:"total_balance"`
}
