package main

import (
	"flag"
	"fmt"

	"gopkg.in/go-pg/migrations.v4"

	"speedcafe/config"
	"speedcafe/database"
)

const verbose = true

func main() {
	flag.Parse()

	config := config.NewConfig()
	config.Init("../../etc/app.ini")

	db := database.NewConnection(config.Database)

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		panic(err.Error())
	}
	if verbose {
		if newVersion != oldVersion {
			fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
		} else {
			fmt.Printf("version is %d\n", oldVersion)
		}
	}
}
