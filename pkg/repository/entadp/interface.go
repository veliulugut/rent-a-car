package entadp

import (
	"context"
	"rent-a-car/pkg/repository/dto"
)

type RepositoryInterface interface {
	UserRepository()
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, c *dto.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByID(ctx context.Context, id int) (*dto.User, error)
	UpdatedUser(ctx context.Context, id int, c *dto.UserUpdate) error
}
