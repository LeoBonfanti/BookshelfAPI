package model

type UserResponse struct {
	ID    int    `json:"id_user"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
