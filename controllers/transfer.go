package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
)

func CheckUserExist(phoneNumber string) (bool, entities.User) {
	var user entities.User
	querySelect := "SELECT * FROM balances u WHERE u.user_id =  ?"
	config.DB.QueryRow(querySelect, phoneNumber).Scan(&user)
	if user.Name != "" {
		return false, user
	}

	return true, user
}

func GetSaldo(user_id int) float64 {
	var balance entities.Balance
	querySelect := "SELECT u.total_balance FROM balances u WHERE u.user_id =  ?"
	config.DB.QueryRow(querySelect, user_id).Scan(&balance.TotalBalance)
	return balance.TotalBalance
}

func ChangeSaldo(user_id int, nominal float64) {
	queryUpdate := "UPDATE balances SET total_balance = total_balance + ?, WHERE user_id = ?"
	_, errInsert := config.DB.Exec(queryUpdate, nominal, LoggedInUser.ID)
	if errInsert != nil {
		fmt.Println(errInsert.Error())
	}
}

func Transfer() bool {
	phoneNumber, nominal := "", 0
	fmt.Println("Masukan nomer telepon penerima")
	fmt.Scanln(&phoneNumber)
	//cek ada atau tidak
	CheckUserExist, user := CheckUserExist(phoneNumber)
	if !CheckUserExist {
		fmt.Println("Data penerima ada")
		return false
	}

	fmt.Println("Masukan nominal")
	fmt.Scanln(&nominal)
	//cek jika nominal tersisa lebih besar dari yang di transfers
	if saldo := GetSaldo(int(LoggedInUser.ID)); saldo < float64(nominal) {
		fmt.Println("Saldo tidak mencukupi")
		return false
	}

	//saldo penerima ditambah
	ChangeSaldo(int(user.ID), float64(nominal))

	//saldo pengirim berkurang
	ChangeSaldo(int(LoggedInUser.ID), float64(nominal)*-1)

	//insert into history
	_, err := config.DB.Exec("INSERT INTO history_balances(user_id,user_id_to,balance_type_id,total) values (?,?,?,?)", LoggedInUser.ID, user.ID, 2, nominal)
	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	fmt.Println("Insert Success!")

	return true
}

type HasilHistoryTransfer struct {
	ID           int
	UserPenerima string
	Tipe         string
	Total        string
	Date         string
}

func HistoryTransfer() {
	rows, err := config.DB.Query("select h.id,u.name,b.balance_type,h.total,h.created_at from history_balances h,users u,balance_types b where h.balance_type_id = b.id AND h.user_id_to = u.id AND user_id = ? order by h.id DESC", LoggedInUser.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var response HasilHistoryTransfer
		rows.Scan(
			&response.ID,
			&response.UserPenerima,
			&response.Tipe,
			&response.Total,
			&response.Date,
		)

		fmt.Println(response)
	}
}
