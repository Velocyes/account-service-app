package entities

import "gorm.io/gorm"

type HistoryBalance struct {
	gorm.Model
	userId        int
	balanceTypeId int
	total         int
}
