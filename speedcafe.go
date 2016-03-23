package speedcafe

import (
	"github.com/gin-gonic/gin"
	"github.com/dskarataev/speedcafe/interfaces"
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