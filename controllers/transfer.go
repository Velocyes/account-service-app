package controllers

import (
	"account-service-app/config"
	"fmt"
)

func Transfer() bool {
	phoneNumber, nominal := "", 0
	fmt.Println("Enter the recipient's phone number")
	fmt.Scanln(&phoneNumber)
	//cek ada atau tidak
	CheckUserExist, user := CheckUserExist(phoneNumber)
	if !CheckUserExist {
		fmt.Println("Recipient's data not found")
		return false
	}

	fmt.Println("Transfer to ", user.Name)
	fmt.Println("Enter your Transfer nominal")
	fmt.Scanln(&nominal)
	//cek jika nominal tersisa lebih besar dari yang di transfers
	if saldo := GetSaldo(int(LoggedInUser.ID)); saldo < float64(nominal) {
		fmt.Println("Your balance is insufficient")
		return false
	}

	if nominal < 100 {
		fmt.Println("Minimum Transfer is 100")
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

	fmt.Println("Transfer Success!")

	return true
}

type HasilHistoryTransfer struct {
	UserPenerima string
	Tipe         string
	Total        string
	Date         string
}

func HistoryTransfer() {
	rows, err := config.DB.Query("select u.name,b.balance_type,h.total,h.created_at from history_balances h,users u,balance_types b where h.balance_type_id = b.id AND h.user_id_to = u.id AND user_id = ? order by h.id DESC", LoggedInUser.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Transfer history :")
	for rows.Next() {
		var response HasilHistoryTransfer
		rows.Scan(
			&response.UserPenerima,
			&response.Tipe,
			&response.Total,
			&response.Date,
		)

		fmt.Println(response.Total + " | Recipient : " + response.UserPenerima + ", " + response.Date)
	}
}
