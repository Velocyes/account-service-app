package entities

import "gorm.io/gorm"

type BalanceType struct {
	gorm.Model
	BalanceType string
}