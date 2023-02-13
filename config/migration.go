package config

import "account-service-app/entities"

func InitMigration() {
	DB.AutoMigrate(
		&entities.User{},
		&entities.BalanceType{},
		&entities.Balance{},
		&entities.HistoryBalance{},
	)
}
