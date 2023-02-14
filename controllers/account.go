package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"errors"
	"fmt"
)

func ReadAccount() {
	fmt.Printf("Nama : %s\n", LoggedInUser.Name)
	fmt.Printf("Nomor Telepon : %s\n", LoggedInUser.PhoneNumber)
	fmt.Printf("Saldo : %.2f\n", GetSaldo(int(LoggedInUser.ID)))
	fmt.Println()
}

func UpdateAccount() bool {
	name, phoneNumber, tempPhoneNumber := "", "", ""
	fmt.Printf("Insert your new name : ")
	fmt.Scanln(&name)
	fmt.Printf("Insert your new phone number : ")
	fmt.Scanln(&phoneNumber)
	fmt.Println()

	flag := true
	if name == "" {
		fmt.Printf("Name,")
		flag = false
	}

	if phoneNumber == "" {
		fmt.Printf("Phone Number,")
		flag = false
	}

	if !flag {
		fmt.Println(" Cannot be null.")
		return flag
	}

	var error error
	if phoneNumber == LoggedInUser.PhoneNumber {
		tempPhoneNumber = phoneNumber

		//make error for update validation
		error = errors.New("dummy error")
	} else {
		querySelectUser := "SELECT u.phone_number FROM users u WHERE u.phone_number = ?"
		error = config.DB.QueryRow(querySelectUser, phoneNumber).Scan(&tempPhoneNumber)
	}

	if error != nil {
		queryUpdateUser := "UPDATE users SET name = ?, phone_number = ? WHERE id = ?"
		_, errInsert := config.DB.Exec(queryUpdateUser, name, phoneNumber, LoggedInUser.ID)
		if errInsert != nil {
			fmt.Println(errInsert.Error())
		} else {
			LoggedInUser.Name = name
			LoggedInUser.PhoneNumber = phoneNumber
			fmt.Println("Update data success")
		}
	} else {
		fmt.Println("The phone number is already in use")
	}
	fmt.Println()

	return true
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

func ReadUser() {
	var phoneNumber string
	fmt.Println("Masukan nomor telepon")
	fmt.Scanln(&phoneNumber)
	CheckUserExist, user := CheckUserExist(phoneNumber)
	if !CheckUserExist {
		fmt.Println("User Not Found")
	} else {
		fmt.Println("--User Profile--")
		fmt.Println("Name :", user.Name)
		fmt.Println("Number Phone :", user.PhoneNumber)
	}
	fmt.Println()
}

func CheckUserExist(phoneNumber string) (bool, entities.User) {
	var user entities.User
	querySelect := "SELECT id,name,phone_number FROM users u WHERE u.phone_number =  ?"
	config.DB.QueryRow(querySelect, phoneNumber).Scan(&user.ID, &user.Name, &user.PhoneNumber)

	if user.Name == "" {
		return false, user
	}

	return true, user
}
