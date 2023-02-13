package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
)

func Register() {
	var user entities.User
	fmt.Println("Masukan nama : ")
	fmt.Scanln(&user.Name)
	fmt.Println("Masukan nomor telepon")
	fmt.Scanln(&user.PhoneNumber)
	fmt.Println("Masukan password : ")
	fmt.Scanln(&user.Password)
	checkPhoneNumber := config.DB.Raw("SELECT u.phone_number FROM USERS u WHERE u.phone_number = ?", user.PhoneNumber).Scan(&user)
	fmt.Println(checkPhoneNumber)
}

func Login() {

}