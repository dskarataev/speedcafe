package main

import (
	"speedcafe"

	"os"
)

func main() {
	// read environment variables
	port := os.Getenv("PORT")
	configPath := os.Getenv("CONFIG_PATH")

	if port == "" {
		panic("Port must be set in PORT environment variable")
	}

	if configPath == "" {
		panic("Config path must be set in CONFIG_PATH environment variable")
	}

	app := speedcafe.NewApp()

	app.SetPort(port)
	app.SetConfigPath(configPath)

	err := app.Init()
	if err != nil {
		panic("Cannot init the app, error: " + err.Error())
	}

	app.Run()
}
