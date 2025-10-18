package main

import (
	"pismo/app"
	_ "pismo/docs"
)

// @title Pismo API
// @version 1.0
// @description This is the Pismo API documentation.
// @host localhost:9001
// @BasePath /api/v1
func main() {
	app.Start()
}
