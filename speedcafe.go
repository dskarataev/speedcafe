package speedcafe

import (
	"speedcafe/interfaces"

	"github.com/gin-gonic/gin"
)

type Speedcafe struct {
	Engine *gin.Engine
}

func NewApp() interfaces.ISpeedCafe {
	return &Speedcafe{}
}

func (app *Speedcafe) Init() {
	// attach router
	app.Engine = gin.Default()

	// add routes
	app.addRoutes()
}

func (app *Speedcafe) Run(port string) {
	app.Engine.Run(":" + port)
}