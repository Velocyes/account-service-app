package entities

import "gorm.io/gorm"

type HistoryBalance struct {
	gorm.Model
	UserId        int          `json:"user_id"`
	User          *User        `gorm:"foreignKey:UserId"`
	UserIdTo      int          `json:"user_id_to"`
	UserTo        *User        `gorm:"foreignKey:UserIdTo"`
	BalanceTypeId int          `json:"balance_type_id"`
	BalanceType   *BalanceType `gorm:"foreignKey:BalanceTypeId"`
	Total         int          `json:"total"`
}
