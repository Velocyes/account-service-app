package main

import (
	"account-service-app/config"
	"account-service-app/menu"
)

func init() {
	config.InitEnvironment()
	config.ConnectToDB()
}

func main() {
	if config.DB != nil {
		menu.AuthMenu()
	}
}
