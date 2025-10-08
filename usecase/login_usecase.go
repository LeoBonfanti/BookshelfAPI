package usecase

import (
	"BookShelfAPI/auth"
	"BookShelfAPI/repository"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	repository repository.LoginRepository
}

func NewLoginUseCase(repo repository.LoginRepository) LoginUseCase {
	return LoginUseCase{
		repository: repo,
	}
}

func (lu *LoginUseCase) Login(email string, password string) (string, error) {
	user, err := lu.repository.GetByEmail(email)

	if err != nil {
		return "", fmt.Errorf("USER NOT FOUND")
	}

	user.Password = strings.TrimSpace(user.Password)
	password = strings.TrimSpace(password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", fmt.Errorf("INVALID PASSWORD")
	}

	token, err := auth.GenerateKey(user.ID)

	if err != nil {
		return "", fmt.Errorf("ERROR GENERATING TOKEN")
	}

	return token, nil
}
