package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
)

func ReadAccount() {
	fmt.Printf("Nama : %s\n", LoggedInUser.Name)
	fmt.Printf("Nomor Telepon : %s\n", LoggedInUser.PhoneNumber)
}

func UpdateAccount() {
	name, phoneNumber := "", ""
	fmt.Printf("Masukan nama baru anda : ")
	fmt.Scanln(&name)
	fmt.Printf("Masukan nomor telepon baru anda : ")
	fmt.Scanln(&phoneNumber)
	queryUpdate := "UPDATE users SET name = ?, phone_number = ? WHERE id = ?"
	_, errInsert := config.DB.Exec(queryUpdate, name, phoneNumber, LoggedInUser.ID)
	if errInsert != nil {
		fmt.Println(errInsert.Error())
	} else {
		LoggedInUser.Name = name
		LoggedInUser.PhoneNumber = phoneNumber
	}
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