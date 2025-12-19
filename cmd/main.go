package main

import (
"moscow-zoo/console"
"moscow-zoo/di"
)

func main() {
	container := di.NewContainer()
	app := console.NewApp(container.GetZoo())
	app.Run()
}
