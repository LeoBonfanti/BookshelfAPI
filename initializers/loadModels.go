package initializers

import (
	"BookShelfAPI/controller"
	"BookShelfAPI/db"
	"BookShelfAPI/repository"
	"BookShelfAPI/usecase"
)

func UserInitializer() *controller.UserController {

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)

	return UserController
}

func LoginInitializer() *controller.LoginController {

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	LoginRepository := repository.NewLoginRepository(dbConnection)
	LoginUseCase := usecase.NewLoginUseCase(LoginRepository)
	LoginController := controller.NewLoginController(LoginUseCase)

	return LoginController
}
