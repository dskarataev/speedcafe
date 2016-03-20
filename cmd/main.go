package main

import "speedcafe"

func main() {
	app := speedcafe.NewApp()
	app.Init()
	app.Run(":8888")
}
