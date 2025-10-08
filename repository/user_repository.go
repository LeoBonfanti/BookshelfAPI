package repository

import (
	"BookShelfAPI/model"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {

	query := "SELECT id, user_name, email FROM users ORDER BY id"
	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Email)
		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {

	var id int
	query, err := ur.connection.Prepare("INSERT INTO users" +
		"(user_name, email, password)" +
		" VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Name, user.Email, hash).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (ur *UserRepository) DeleteUser(id_user int) (int, error) {

	sql := "DELETE FROM users WHERE id = $1"

	result, err := ur.connection.Exec(sql, id_user)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, nil
	}

	return 1, nil
}

func (ur *UserRepository) GetUserById(id_user int) (*model.User, error) {

	query, err := ur.connection.Prepare("SELECT id, user_name, email FROM users WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.User
	err = query.QueryRow(id_user).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &user, nil
}

func (ur *UserRepository) EditUser(editedUser model.User) (*model.User, error) {

	query, err := ur.connection.Prepare("UPDATE users SET user_name = $1, email = $2 WHERE id = $3")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = query.Exec(editedUser.Name, editedUser.Email, editedUser.ID)

	if err != nil {
		return nil, err
	}

	query.Close()
	return &editedUser, nil
}
