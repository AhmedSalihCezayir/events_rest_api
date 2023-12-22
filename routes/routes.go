package routes

import (
	"example.com/events-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api/v1")

	api.GET("/events", getEvents)
	api.GET("/events/:id", getEventById)
	api.GET("/events/:id/attendees", getEventAttendees)

	authenticated := api.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	
	authenticated.GET("/users", middlewares.AdminCheck, getUsers)
	authenticated.DELETE("/users/:id", middlewares.AdminCheck, deleteUser)

	authenticated.GET("/registered-events", getUserRegisteredEvents)
	authenticated.PUT("/update-user", updateUserInfo)

	api.POST("/signup", signup) // we could've directly added the middleware here as well
	api.POST("/login", login)
}
