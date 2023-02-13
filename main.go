package main

import (
	"account-service-app/config"
	"account-service-app/menu"
)

func main() {
	config.InitEnvironment()
	menu.MenuAuth()

	config.ConnectToDB()
	
}
