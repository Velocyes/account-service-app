package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"account-service-app/helpers"
	"fmt"
	"strconv"
	"time"
)

func TopUp() bool {
	helpers.ClearCmd()
	fmt.Println("========= Top-Up =========")

	nominal := 0
	fmt.Printf("Enter your Top-Up nominal : ")
	fmt.Scanln(&nominal)

	if nominal < 100 {
		fmt.Println("Minimum Top-Up is 100")
		return false
	}

	//Menambahkan saldo top-up ke user
	if ChangeSaldo(int(LoggedInUser.ID), float64(nominal)) {
		//Menambahkan riwayat saldo top-up ke history
		_, err := config.DB.Exec("INSERT INTO history_balances(user_id, balance_type_id ,total) VALUES (?,?,?)", LoggedInUser.ID, 1, nominal)
		if err != nil {
			fmt.Println("Top-up failed")
		}
		fmt.Println("Top-up success!")
		fmt.Println()
	}

	return true
}

func HistoryTopups() {
	helpers.ClearCmd()
	fmt.Println("======= History Top-Up ========")

	rows, err := config.DB.Query("SELECT hb.total, hb.created_at FROM history_balances hb WHERE hb.user_id = ? AND hb.balance_type_id = 1 order by hb.id DESC", LoggedInUser.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Top-up history :")
	for rows.Next() {
		historyBalance := entities.HistoryBalance{}
		rows.Scan(
			&historyBalance.Total,
			&historyBalance.CreatedAt,
		)

		fmt.Println(strconv.Itoa(historyBalance.Total) + " | " + historyBalance.CreatedAt.Format(time.RFC1123))
	}
	fmt.Println()
}
