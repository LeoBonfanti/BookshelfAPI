package repository

import (
	"BookShelfAPI/model"
	"database/sql"
)

type LoginRepository struct {
	connection *sql.DB
}

func NewLoginRepository(connection *sql.DB) LoginRepository {
	return LoginRepository{
		connection: connection,
	}
}

func (r *LoginRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	row := r.connection.QueryRow("SELECT id, user_name, email, password FROM users WHERE email = $1", email)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password)

	return user, err
}
