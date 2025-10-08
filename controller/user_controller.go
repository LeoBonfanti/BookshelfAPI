package controller

import (
	"BookShelfAPI/model"
	"BookShelfAPI/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) *UserController {

	return &UserController{
		userUseCase: usecase,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {

	users, err := u.userUseCase.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) CreateUser(ctx *gin.Context) {

	var user model.User
	err := ctx.BindJSON(&user)

	if user.Email == "" || user.Name == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "EMPTY VALUES"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := u.userUseCase.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var userResponse model.UserResponse
	userResponse.Email = insertedUser.Email
	userResponse.Name = insertedUser.Name
	userResponse.ID = insertedUser.ID

	ctx.JSON(http.StatusCreated, userResponse)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("userId")

	if id == "" {
		response := model.Response{
			Message: "id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "id need to be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedUser, err := u.userUseCase.DeleteUser(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if deletedUser == 0 {
		response := model.Response{
			Message: "user not found in database",
		}
		ctx.JSON(http.StatusNotFound, response)
	}

	ctx.JSON(http.StatusNoContent, deletedUser)
}

func (u *UserController) GetUserById(ctx *gin.Context) {

	id := ctx.Param("userId")

	if id == "" {
		response := model.Response{
			Message: "id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "id need to be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.userUseCase.GetUserById(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		response := model.Response{
			Message: "user not found in database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) EditUser(ctx *gin.Context) {

	var user model.User
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	editedUser, err := u.userUseCase.EditUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, editedUser)

}
