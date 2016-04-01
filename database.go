package speedcafe

import (
	"gopkg.in/pg.v4"
)

func (this *SpeedCafe) initDatabase() {
	this.DB = pg.Connect(&pg.Options{
		User: this.Config.Database.User,
		Database: this.Config.Database.Name,
		Password: this.Config.Database.Password,
		Addr: this.Config.Database.Host,
	})
}
