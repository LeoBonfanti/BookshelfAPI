package main

import (
	"BookShelfAPI/initializers"
	"BookShelfAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	initializers.LoadEnvVariables()

	UserController := initializers.UserInitializer()
	LoginController := initializers.LoginInitializer()
	server := gin.Default()

	routes.SetupRouteUsers(server, UserController)
	routes.SetupRouteLogin(server, LoginController)

	server.Run(":8080")
}
