package usecase

import (
	"BookShelfAPI/model"
	"BookShelfAPI/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (uu *UserUseCase) GetUsers() ([]model.User, error) {
	return uu.repository.GetUsers()
}

func (uu *UserUseCase) CreateUser(user model.User) (model.User, error) {

	userId, err := uu.repository.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}

	user.ID = userId

	return user, nil
}

func (uu *UserUseCase) DeleteUser(id_user int) (int, error) {

	deleted, err := uu.repository.DeleteUser(id_user)

	if err != nil {
		return 0, err
	}

	if deleted == 0 {
		return 0, nil
	}

	return 1, nil
}

func (uu *UserUseCase) GetUserById(id_user int) (*model.User, error) {

	user, err := uu.repository.GetUserById(id_user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UserUseCase) EditUser(editedUser model.User) (*model.User, error) {

	user, err := uu.repository.EditUser(editedUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}
