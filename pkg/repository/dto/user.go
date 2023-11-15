package dto

import "time"

/*
field.String("first_name"),

	field.String("last_name"),
	field.String("username"),
	field.String("email"),
	field.String("phone_number"),
	field.String("password"),
	field.Time("created_at"),
	field.Time("updated_at"),
*/
type User struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type UserUpdate struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	UpdateAt    time.Time `json:"update_at"`
	PhoneNumber string    `json:"phone_number"`
}
