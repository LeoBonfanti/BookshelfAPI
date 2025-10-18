package routes

import (
	"BookShelfAPI/auth"
	"BookShelfAPI/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouteUsers(router *gin.Engine, userController *controller.UserController) {

	protected := router.Group("/")
	protected.Use(auth.ValidateTokenMiddleware())

	users := protected.Group("/users")
	{
		users.GET("/", userController.GetUsers)
		users.POST("/", userController.CreateUser)
		users.GET("/:userId", userController.GetUserById)
		users.PUT("/:userId", userController.EditUser)
		users.DELETE("/:userId", userController.DeleteUser)
	}
}

func SetupRouteLogin(router *gin.Engine, loginController *controller.LoginController) {

	login := router.Group("/login")
	{
		login.POST("/", loginController.Login)
	}
}
