package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
)

func GetSaldo(user_id int) float64 {
	var balance entities.Balance
	querySelect := "SELECT u.total_balance FROM balances u WHERE u.user_id =  ?"
	config.DB.QueryRow(querySelect, user_id).Scan(&balance.TotalBalance)
	return balance.TotalBalance
}

func ChangeSaldo(user_id int, nominal float64) bool {
	queryUpdate := "UPDATE balances SET total_balance = total_balance + ? WHERE user_id = ?"
	_, errInsert := config.DB.Exec(queryUpdate, nominal, user_id)
	if errInsert != nil {
		fmt.Println(errInsert.Error())
		return false
	}
	return true
}
