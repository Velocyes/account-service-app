package menu

import (
	"account-service-app/controllers"
	"account-service-app/helpers"
	"fmt"
)

func AuthMenu() {
	choice := -2
	for choice != -1 {
		helpers.ClearCmd()
		fmt.Println("========== AUTH =========")
		fmt.Println("\t1. Login")
		fmt.Println("\t2. Register")
		fmt.Println("\t0. Exit")
		fmt.Println("=========================")
		fmt.Printf("Enter your choice : ")
		fmt.Scanln(&choice)
		if choice <= -1 || choice >= 3 {
			continue
		}
		if choice == 1 {
			if controllers.Login() {
				MainMenu()
			}
		} else if choice == 2 {
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
		helpers.ClearCmd()
		if controllers.LoggedInUser.Name == "" {
			AuthMenu()
		}

		fmt.Println("========== MENU =========")
		fmt.Printf("Welcome, %s\n", controllers.LoggedInUser.Name)
		fmt.Println("_________________________")

		fmt.Println("1. Read Account")
		fmt.Println("2. Update Account")
		fmt.Println("3. Delete Account")
		fmt.Println("4. Top Up")
		fmt.Println("5. Transfer")
		fmt.Println("6. History Top Up")
		fmt.Println("7. History Transfer")
		fmt.Println("8. Profile Another User")
		fmt.Println("0. Exit")
		fmt.Println("=========================")
		fmt.Printf("Enter your choice : ")
		fmt.Scanln(&choice)
		fmt.Println()
		if choice <= -1 || choice >= 9 {
			continue
		}
		if choice == 1 {
			controllers.ReadAccount()
		} else if choice == 2 {
			controllers.UpdateAccount()
		} else if choice == 3 {
			controllers.DeleteAccount()
		} else if choice == 4 {
			controllers.TopUp()
		} else if choice == 5 {
			controllers.Transfer()
		} else if choice == 6 {
			controllers.HistoryTopups()
		} else if choice == 7 {
			controllers.HistoryTransfer()
		} else if choice == 8 {
			controllers.ReadUser()
		} else if choice == 0 {
			controllers.Logout()
			break
		}

		fmt.Printf("\nPress enter key to continue")
		fmt.Scanln()
	}
}
