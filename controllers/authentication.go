package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
	_ "log"

)

var LoggedInUser entities.User

func Register() {
	var user entities.User
	var form entities.User
	fmt.Printf("Masukan nama : ")
	fmt.Scanln(&user.Name)
	fmt.Printf("Masukan nomor telepon : ")
	fmt.Scanln(&user.PhoneNumber)
	fmt.Printf("Masukan password : ")
	fmt.Scanln(&user.Password)

	querySelect := "SELECT u.phone_number FROM users u WHERE u.phone_number = ?"
	error := config.DB.QueryRow(querySelect, user.PhoneNumber).Scan(&form.PhoneNumber)
	if error != nil {
		queryInsert := "INSERT INTO users (name, phone_number, password) VALUES (?, ?, ?)"
		_, errInsert := config.DB.Exec(queryInsert, user.Name, user.PhoneNumber, user.Password)
		if errInsert != nil {
			fmt.Println(errInsert.Error())
		}
	} else {
		fmt.Println("Data sudah ada")
	}

	
}

func Login() {
	var user entities.User
	phoneNumber, password := "", ""
	fmt.Printf("Masukan nomor telepon : ")
	fmt.Scanln(&phoneNumber)
	fmt.Printf("Masukan password : ")
	fmt.Scanln(&password)
	querySelect := "SELECT u.id, u.name, u.phone_number FROM users u WHERE u.phone_number = ? AND password = ?"
	error := config.DB.QueryRow(querySelect, phoneNumber, password).Scan(&user.ID, &user.Name, &user.PhoneNumber)
	if error != nil {
		fmt.Println("Data tidak ada")
		LoggedInUser = entities.User{}
	} else {
		LoggedInUser = user
	}
	fmt.Println(LoggedInUser)
}