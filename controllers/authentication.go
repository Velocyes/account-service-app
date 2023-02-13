package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
	_ "log"
)

var LoggedInUser entities.User

func Login() bool {
	user, phoneNumber, password := entities.User{}, "", ""
	fmt.Printf("Masukan nomor telepon : ")
	fmt.Scanln(&phoneNumber)
	fmt.Printf("Masukan password : ")
	fmt.Scanln(&password)
	fmt.Println()

	querySelectUser := "SELECT u.id, u.name, u.phone_number FROM users u WHERE u.phone_number = ? AND password = ?"
	error := config.DB.QueryRow(querySelectUser, phoneNumber, password).Scan(&user.ID, &user.Name, &user.PhoneNumber)
	if error != nil {
		fmt.Println("Credential is wrong or user doesn't exists")
		LoggedInUser = entities.User{}
	} else {
		LoggedInUser = user
		return true
	}
	return false
}

func Register() bool {
	user, phoneNumber := entities.User{}, ""
	fmt.Printf("Masukan nama : ")
	fmt.Scanln(&user.Name)
	fmt.Printf("Masukan nomor telepon : ")
	fmt.Scanln(&user.PhoneNumber)
	fmt.Printf("Masukan password : ")
	fmt.Scanln(&user.Password)
	fmt.Println()

	queryCreateUser := "SELECT u.phone_number FROM users u WHERE u.phone_number = ?"
	error := config.DB.QueryRow(queryCreateUser, user.PhoneNumber).Scan(&phoneNumber)
	if error != nil {
		_, errInsert := config.DB.Exec("INSERT INTO users (name, phone_number, password) VALUES (?, ?, ?)", user.Name, user.PhoneNumber, user.Password)
		if errInsert != nil {
			fmt.Println(errInsert.Error())
			return false
		} else {
			LoggedInUser = user
			fmt.Println("Register success")
		}
		return true
	} else {
		fmt.Println("Credential is already used")
	}
	return false
}

func Logout() {
	LoggedInUser = entities.User{}
}