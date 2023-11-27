package user

import "time"

type CreateUserModel struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
}

type GetUserModel struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

type UpdateUserModel struct {
	Email    string    `json:"email"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
	UpdateAt time.Time `json:"update_at"`
}
