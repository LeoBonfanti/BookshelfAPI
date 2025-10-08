package controller

import (
	"BookShelfAPI/model"
	"BookShelfAPI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginUseCase usecase.LoginUseCase
}

func NewLoginController(usecase usecase.LoginUseCase) *LoginController {

	return &LoginController{
		loginUseCase: usecase,
	}
}

func (u *LoginController) Login(ctx *gin.Context) {

	var loginReq model.LoginReq
	err := ctx.BindJSON(&loginReq)

	if loginReq.Password == "" || loginReq.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "EMAIL AND PASSWORD REQUIRED"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := u.loginUseCase.Login(loginReq.Email, loginReq.Password)

	if err != nil {
		if err.Error() == "USER NOT FOUND" || err.Error() == "INVALID PASSWORD" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": key})
}
