package controllers

import (
	"account-service-app/config"
	"account-service-app/entities"
	"fmt"
	_ "log"

	"golang.org/x/crypto/bcrypt"
)

var LoggedInUser entities.User

func Login() bool {
	user, phoneNumber, password := entities.User{}, "", ""
	fmt.Printf("Masukan nomor telepon : ")
	fmt.Scanln(&phoneNumber)
	fmt.Printf("Masukan password : ")
	fmt.Scanln(&password)
	fmt.Println()

	querySelectUser := "SELECT u.id, u.name, u.phone_number,u.password FROM users u WHERE u.phone_number = ?"
	error := config.DB.QueryRow(querySelectUser, phoneNumber).Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password)
	if error != nil {
		LoggedInUser = entities.User{}
		fmt.Println("Credential is wrong or user doesn't exists")
		fmt.Println()
	} else {
		//check if password is true
		if CheckPasswordHash(password, user.Password) {
			LoggedInUser = user
			return true
		}

		LoggedInUser = entities.User{}
		fmt.Println("Credential is wrong or user doesn't exists")
		fmt.Println()
	}
	return false
}

func Register() bool {
	user, phoneNumber := entities.User{}, ""
	fmt.Printf("Insert your name : ")
	fmt.Scanln(&user.Name)
	fmt.Printf("Insert your phone number : ")
	fmt.Scanln(&user.PhoneNumber)
	fmt.Printf("Insert password : ")
	fmt.Scanln(&user.Password)
	fmt.Println()

	flag := true
	if user.Name == "" {
		fmt.Printf("Name,")
		flag = false
	}

	if user.PhoneNumber == "" {
		fmt.Printf("Phone Number,")
		flag = false
	}

	if user.Password == "" {
		fmt.Printf("Password,")
		flag = false
	}

	if !flag {
		fmt.Println(" Cannot be null.")
		return flag
	}

	queryCreateUser := "SELECT u.phone_number FROM users u WHERE u.phone_number = ?"
	error := config.DB.QueryRow(queryCreateUser, user.PhoneNumber).Scan(&phoneNumber)
	if error != nil {
		user.Password, _ = HashPassword(user.Password)
		result, errInsert := config.DB.Exec("INSERT INTO users (name, phone_number, password) VALUES (?, ?, ?)", user.Name, user.PhoneNumber, user.Password)
		if errInsert != nil {
			fmt.Println(errInsert.Error())
			return false
		} else {
			id, _ := result.LastInsertId()
			user.ID = uint(id)
			LoggedInUser = user
		}

		//Insert Initial Balance
		saldo := 0
		fmt.Printf("Masukan saldo awal user : ")
		fmt.Scanln(&saldo)
		user_id, _ := result.LastInsertId()
		config.DB.Exec("INSERT INTO balances (user_id, total_balance) VALUES (?, ?)", user_id, saldo)

		fmt.Println("Register user success")
		fmt.Println()
		return true
	} else {
		fmt.Println("Credential is aslready used")
	}

	return false
}

func Logout() {
	LoggedInUser = entities.User{}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
