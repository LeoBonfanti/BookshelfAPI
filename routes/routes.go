package routes

import (
	"BookShelfAPI/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controller.UserController) {

	users := router.Group("/users")
	{
		users.GET("/", userController.GetUsers)
		users.POST("/", userController.CreateUser)

		users.GET("/:userId", userController.GetUserById)
		users.PUT("/:userId", userController.EditUser)
		users.DELETE("/:userId", userController.DeleteUser)
	}
}
