package user

import (
	"context"
	"rent-a-car/pkg/passwd"
	"rent-a-car/pkg/repository/entadp"
)

var _ InterfaceUser = (*User)(nil)

type User struct {
	pw   passwd.Interface
	repo entadp.RepositoryInterface
}

func New(pw passwd.Interface, repo entadp.RepositoryInterface) *User {
	return &User{
		pw:   pw,
		repo: repo,
	}
}

func (u *User) CreateUser(ctx context.Context, c *CreateUserModel) error {
	return nil
}

func (u *User) DeleteUser(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetUserByID(ctx context.Context, id int) (*GetUserModel, error) {
	//TODO implement me
	panic("implement me")
}
