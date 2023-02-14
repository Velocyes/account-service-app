package entities

import "time"

type HistoryBalance struct {
	ID            uint
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     time.Time    `json:"deleted_at"`
	UserId        int          `json:"user_id"`
	User          *User        `gorm:"foreignKey:UserId"`
	UserIdTo      int          `json:"user_id_to"`
	UserTo        *User        `gorm:"foreignKey:UserIdTo"`
	BalanceTypeId int          `json:"balance_type_id"`
	BalanceType   *BalanceType `gorm:"foreignKey:BalanceTypeId"`
	Total         int          `json:"total"`
}
