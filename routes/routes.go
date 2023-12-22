package routes

import (
	"example.com/events-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.GET("/events/:id/attendees", getEventAttendees)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	authenticated.GET("/registered-events", getUserRegisteredEvents)

	server.POST("/signup", signup) // we could've directly added the middleware here as well
	server.POST("/login", login)

	server.GET("/users", getUsers)
}
