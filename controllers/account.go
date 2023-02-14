package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
)

func ReadAccount() {
	fmt.Printf("Nama : %s\n", LoggedInUser.Name)
	fmt.Printf("Nomor Telepon : %s\n", LoggedInUser.PhoneNumber)
	fmt.Printf("Saldo : %.2f\n", GetSaldo(int(LoggedInUser.ID)))
	fmt.Println()
}

func UpdateAccount() {
	name, phoneNumber, tempPhoneNumber := "", "", ""
	fmt.Printf("Masukan nama baru anda : ")
	fmt.Scanln(&name)
	fmt.Printf("Masukan nomor telepon baru anda : ")
	fmt.Scanln(&phoneNumber)
	fmt.Println()

	querySelectUser := "SELECT u.phone_number FROM users u WHERE u.phone_number = ?"
	error := config.DB.QueryRow(querySelectUser, phoneNumber).Scan(&tempPhoneNumber)
	if error != nil {
		queryUpdateUser := "UPDATE users SET name = ?, phone_number = ? WHERE id = ?"
		_, errInsert := config.DB.Exec(queryUpdateUser, name, phoneNumber, LoggedInUser.ID)
		if errInsert != nil {
			fmt.Println(errInsert.Error())
		} else {
			LoggedInUser.Name = name
			LoggedInUser.PhoneNumber = phoneNumber
			fmt.Println("Update data berhasil")
		}
	} else {
		fmt.Println("Nomor telepon sudah digunakan")
	}
	fmt.Println()
}

func DeleteAccount() {
	queryUpdate := "DELETE FROM users WHERE id = ?"
	_, errInsert := config.DB.Exec(queryUpdate, LoggedInUser.ID)
	if errInsert != nil {
		fmt.Println(errInsert.Error())
	} else {
		LoggedInUser = entities.User{}
	}
}

func CheckUserExist(phoneNumber string) (bool, entities.User) {
	var user entities.User
	querySelect := "SELECT * FROM balances u WHERE u.user_id =  ?"
	config.DB.QueryRow(querySelect, phoneNumber).Scan(&user)
	if user.Name != "" {
		return false, user
	}

	return true, user
}