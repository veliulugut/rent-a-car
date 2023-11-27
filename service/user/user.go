package user

import (
	"context"
)

var _ InterfaceUser = (*User)(nil)

type User struct {
}

func (u User) CreateUser(ctx context.Context, c *CreateUserModel) error {
	//TODO implement me
	panic("implement me")
}

func (u User) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u User) UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error {
	//TODO implement me
	panic("implement me")
}

func (u User) GetUserByID(ctx context.Context, id int) (*GetUserModel, error) {
	//TODO implement me
	panic("implement me")
}
