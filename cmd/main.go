package main

import (
	"mini-pos/database"
	"mini-pos/server"
	"mini-pos/util"
)

func main() {
	// load config from config.env
	util.LoadConfig()
	// setup model and create seeder
	database.SetupModels()
	// call endpoint routing (serve on localhost:8000)
	server.RegisterRoutes()
}
