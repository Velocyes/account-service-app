package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"account-service-app/helpers"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadAccount() {
	helpers.ClearCmd()
	fmt.Println("========= Account ========")
	fmt.Printf("Nama \t\t: %s\n", LoggedInUser.Name)
	fmt.Printf("Nomor Telepon \t: %s\n", LoggedInUser.PhoneNumber)
	fmt.Printf("Saldo \t\t: %.2f\n", GetSaldo(int(LoggedInUser.ID)))
	fmt.Println()
}

func UpdateAccount() bool {
	helpers.ClearCmd()
	fmt.Println("====== Update Account =======")
	name, phoneNumber, tempPhoneNumber := "", "", ""
	fmt.Printf("Insert your new name : ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		name = text
	}
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

	_, err := strconv.Atoi(phoneNumber)
	if err != nil {
		fmt.Println("Phone number only for integer input.")
		return flag
	}

	var error error
	if phoneNumber == LoggedInUser.PhoneNumber {
		tempPhoneNumber = phoneNumber

		//Make error for update validation
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
			fmt.Println("Update Data Success")
		}
	} else {
		fmt.Println("The Phone Number is Already in Use")
	}
	fmt.Println()

	return true
}

func DeleteAccount() bool {
	helpers.ClearCmd()
	fmt.Println("======= Delete Account =======")
	fmt.Println("Are you sure to delete this account ?")
	fmt.Println("Y -> Yes, N -> No")
	flag := ""
	fmt.Scanln(&flag)

	if flag == "Y" || flag == "y" {
		queryUpdate := "DELETE FROM users WHERE id = ?"
		_, errInsert := config.DB.Exec(queryUpdate, LoggedInUser.ID)
		if errInsert != nil {
			fmt.Println(errInsert.Error())
		} else {
			LoggedInUser = entities.User{}
		}

		return true
	}

	return true
}

func ReadUser() {
	helpers.ClearCmd()
	fmt.Println("========= PROFILE ========")

	var phoneNumber string
	fmt.Printf("Enter the phone number :")
	fmt.Scanln(&phoneNumber)
	CheckUserExist, user := CheckUserExist(phoneNumber)
	if !CheckUserExist {
		fmt.Println("User Not Found")
	} else {
		fmt.Println("Name \t\t:", user.Name)
		fmt.Println("Number Phone \t:", user.PhoneNumber)
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
