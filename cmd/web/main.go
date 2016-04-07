package main

import (
	"speedcafe"
)

func main() {
	app := speedcafe.NewApp()
	err := app.Init()
	if err != nil {
		panic("Cannot init the app, error: " + err.Error())
	}
	app.Run()
}
