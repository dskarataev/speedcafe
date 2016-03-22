package main

import (
	"speedcafe"
	"os"
)

func main() {
	address := ":" + os.Getenv("PORT")

	app := speedcafe.NewApp()
	app.Init()
	app.Run(address)
}
