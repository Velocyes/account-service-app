package main

import (
	"account-service-app/config"
	"account-service-app/menu"
)

func main() {
	menu.MenuAuth()

	config.ConnectToDB()
}
