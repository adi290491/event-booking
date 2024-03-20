package main

import (
	"event-booking/configurations"
	"event-booking/database"
	"event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	var cfg configurations.Config
	err := cfg.Init()

	if err != nil {
		panic(err)
	}

	database.InitConnection(&cfg)
	server := gin.Default()

	routes.RegisterEndpoints(server)
	server.Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
