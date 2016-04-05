package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/go-pg/migrations.v4"

	"speedcafe/config"
	"speedcafe/database"
)

const verbose = true

func main() {
	flag.Parse()

	config := config.NewConfig()
	db := database.NewConnection(config.Database)

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if verbose {
		if newVersion != oldVersion {
			fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
		} else {
			fmt.Printf("version is %d\n", oldVersion)
		}
	}
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}