package main

import (
	"fmt"
	"gopkg.in/go-pg/migrations.v4"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table cafe")
		_, err := db.Exec(`CREATE TABLE cafe()`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table cafe...")
		_, err := db.Exec(`DROP TABLE cafe`)
		return err
	})
}