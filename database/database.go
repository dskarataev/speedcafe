package database

import (
	"gopkg.in/pg.v4"
	"speedcafe/config"
)

func NewConnection(config config.Database) *pg.DB {
	return pg.Connect(&pg.Options{
		User: config.User,
		Database: config.Name,
		Password: config.Password,
		Addr: config.Host,
	})
}
