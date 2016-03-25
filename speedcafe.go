package speedcafe

import (
	"speedcafe/config"
	"speedcafe/interfaces"
	"speedcafe/services/foursquare"

	"github.com/gin-gonic/gin"
)

type SpeedCafe struct {
	Engine *gin.Engine
	Port   string

	ConfigPath string
	Config     *config.Config

	FoursquareService interfaces.IFoursquareService
}

func NewApp() interfaces.ISpeedCafe {
	return &SpeedCafe{}
}

func (this *SpeedCafe) SetConfigPath(path string) {
	this.ConfigPath = path
}

func (this *SpeedCafe) SetPort(port string) {
	this.Port = port
}

func (this *SpeedCafe) Init() error {
	// config
	err := this.initConfig()

	// services
	this.FoursquareService = foursquare.NewFoursquareService(
		this.Config.FoursquareClient.ID,
		this.Config.FoursquareClient.SecretKey,
	)

	// router and HTTP server
	this.Engine = gin.Default()

	// templates
	this.Engine.LoadHTMLGlob("templates/*")

	// routes
	this.addRoutes()

	return err
}

func (this *SpeedCafe) initConfig() error {
	this.Config = config.NewConfig()
	return this.Config.Init(this.ConfigPath)
}

func (this *SpeedCafe) Run() {
	this.Engine.Run(":" + this.Port)
}
