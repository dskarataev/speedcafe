package main

import (
	"speedcafe"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

        if port == "" {
                log.Fatal("$PORT must be set")
        }

	app := speedcafe.NewApp()
	app.Init()
	app.Run(port)
}
