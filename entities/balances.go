package entities

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserId int
	totalBalance float64
}