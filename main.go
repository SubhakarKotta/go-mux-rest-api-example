package main

import (
	"go-mux-rest-api-example/app"
	"go-mux-rest-api-example/config"
)

func main() {
	config := config.GetConfig()
	app := app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
