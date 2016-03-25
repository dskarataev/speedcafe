package speedcafe

import "speedcafe/handlers"

func (this *SpeedCafe) addRoutes() {
	this.Engine.GET("/", handlers.Index)
	this.Engine.GET("/hello/world", handlers.HelloWorld)
}
