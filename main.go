package main

import (
	_ "event-booking/database"
	"event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// database.InitConnection()
	server := gin.Default()

	routes.RegisterEndpoints(server)

	server.Run(":8080")
}
