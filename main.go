package main

import (
	"demo/restserver/db"
	"demo/restserver/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
