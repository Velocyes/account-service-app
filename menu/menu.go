package menu

import (
	"account-service-app/controllers"
	"fmt"
)

func AuthMenu() {
	choice := -2
	for choice != -1 {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Printf("Masukan pilihan anda : ")
		fmt.Scanln(&choice)
		fmt.Println()
		if choice <= -1 || choice >= 3 {
			continue
		}
		if choice == 1{
			if controllers.Login() {
				MainMenu()
			}
		} else if choice == 2{
			if controllers.Register() {
				MainMenu()
			}
		} else if choice == 0 {
			break
		}
	}
}

func MainMenu() {
	choice := -2
	for choice != -1 {
		if controllers.LoggedInUser.Name == "" {
			AuthMenu()
		}
		fmt.Printf("Welcome, %s\n", controllers.LoggedInUser.Name)
		fmt.Println("1. Read Account")
		fmt.Println("2. Update Account")
		fmt.Println("3. Delete Account")
		fmt.Println("4. Top Up")
		fmt.Println("5. Transfer")
		fmt.Println("6. History Top Up")
		fmt.Println("7. History Transfer")
		fmt.Println("8. Profile Another User")
		fmt.Println("0. Exit")
		fmt.Printf("Masukan pilihan anda : ")
		fmt.Scanln(&choice)
		fmt.Println()
		if choice <= -1 || choice >= 8 {
			continue
		}
		if choice == 1{
			controllers.ReadAccount()
		} else if choice == 2{
			controllers.UpdateAccount()
		} else if choice == 3{
			controllers.DeleteAccount()
		} else if choice == 4{
			
		} else if choice == 5{
			
		} else if choice == 6{
			
		} else if choice == 7{
			
		} else if choice == 8{
			
		} else if choice == 0{
			controllers.Logout()
			break
		}
	}
}
