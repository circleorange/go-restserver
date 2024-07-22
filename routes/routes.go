package routes

import (
	"demo/restserver/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)

	server.POST("/login", login)
	server.GET("/events", getEvents)
	server.GET("/events/:id")

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // middleware executed before route handlers become active
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}
