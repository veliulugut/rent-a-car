package user

import "context"

type InterfaceUser interface {
	CreateUser(ctx context.Context, c *CreateUserModel) error
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error
	GetUserByID(ctx context.Context, id int) (*GetUserModel, error)
}
