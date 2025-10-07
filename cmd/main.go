package main

import (
	"BookShelfAPI/controller"
	"BookShelfAPI/db"
	"BookShelfAPI/repository"
	"BookShelfAPI/routes"
	"BookShelfAPI/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("ERROR loading .env")
	}

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
