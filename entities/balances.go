package entities

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserId int
	TotalBalance float64
}