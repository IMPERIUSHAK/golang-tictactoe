package models

type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"user_name"`
	PasswordHash string `json:"password_hash"`
}

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
