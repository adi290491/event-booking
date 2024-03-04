package routes

import (
	"event-booking/gateway"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(gateway.Authenticate)
	authenticated.GET("/events", getEvents)
	authenticated.POST("/events", createEvent)
	authenticated.GET("/events/:id", getEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// user routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
