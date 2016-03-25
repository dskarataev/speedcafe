package speedcafe

import "speedcafe/handlers"

const (
	URLHelloWorld = "/hello/world"
)

func (this *SpeedCafe) addRoutes() {
	this.Engine.GET(URLHelloWorld, handlers.HelloWorld)
}
