package main

import (
	"log"
	"simple_go_app/config"
	"simple_go_app/controllers"
	"simple_go_app/repositories"
)

func main() {
	config.LoadConfigs()
	config.PrintConfigs()
	repositories.InitDB()
	repositories.ApplyConfigs()
	repositories.InitSchema()
	if config.GetVar(config.SeedDB) == "true" {
		repositories.SeedDB()
	}
	server := controllers.InitializeServer(config.GetVar(config.APIHost), config.GetVar(config.APIPort))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
