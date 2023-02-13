package menu

import (
	"account-service-app/controllers"
	"fmt"
)

func MenuAuth() {
	choice := -1
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	for choice != 0 {
		fmt.Scanln(&choice)
		if choice == 1{
			controllers.Register()
		} else if choice == 2{
			controllers.Login()
		}
	}
}

func MainMenu() {
	fmt.Println("1. Read Account")
	fmt.Println("2. Update Account")
	fmt.Println("3. Delete Account")
	fmt.Println("4. Top Up")
	fmt.Println("5. Transfer")
	fmt.Println("6. History Top Up")
	fmt.Println("7. History Transfer")
	fmt.Println("8. Profile Another User")
	fmt.Println("0. Exit")
}
