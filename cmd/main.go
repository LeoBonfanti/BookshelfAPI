package main

import (
	"BookShelfAPI/controller"
	"BookShelfAPI/db"
	"BookShelfAPI/repository"
	"BookShelfAPI/routes"
	"BookShelfAPI/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)

	server := gin.Default()
	routes.SetupRoutes(server, UserController)

	server.Run(":8080")
}
