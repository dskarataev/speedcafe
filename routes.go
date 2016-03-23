package speedcafe

import "github.com/dskarataev/speedcafe/handlers"

const (
	URLHelloWorld = "/hello/world"
)

func (app *Speedcafe) addRoutes() {
	app.Engine.GET(URLHelloWorld, handlers.HelloWorld)
}
