package speedcafe

import (
	"speedcafe/config"
	"speedcafe/interfaces"
	"speedcafe/services/foursquare"
	"speedcafe/database"
	"github.com/gin-gonic/gin"
	"gopkg.in/pg.v4"
)

type SpeedCafe struct {
	Engine *gin.Engine

	Config     *config.Config

	FoursquareService interfaces.IFoursquareService

	DB *pg.DB
}

func NewApp() interfaces.ISpeedCafe {
	return &SpeedCafe{}
}

func (this *SpeedCafe) Init() error {
	err := this.initConfig()

	if err != nil {
		return err
	}

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
	
	// db
	this.DB = database.NewConnection(this.Config.Database);

	return err
}

func (this *SpeedCafe) initConfig() error {
	this.Config = config.NewConfig();
	return this.Config.Init()
}

func (this *SpeedCafe) Run() {
	this.Engine.Run(":" + this.Config.HttpPort)
}
