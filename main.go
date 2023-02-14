package main

import (
	"account-service-app/config"
	"account-service-app/helpers"
	"account-service-app/menu"
	"time"
)

func init() {
	config.InitEnvironment()
	config.ConnectToDB()
}

func main() {
	helpers.StartCmd()
	time.Sleep(3 * time.Second)
	defer config.DB.Close()
	if config.DB != nil {
		menu.AuthMenu()
	}
}
