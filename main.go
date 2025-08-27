package main

import (
	"AuthInGo/app"
)

func main() {
	appConfig := app.NewConfig()
	application := app.NewApplication(appConfig)
	application.Start()
}
