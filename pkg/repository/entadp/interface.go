//go:generate mockgen -destination=user_mocks.go -package=entadp rent-a-car/pkg/repository/entadp UserRepositoryInterface
//go:generate mockgen -destination=repository_mocks.go -package=entadp rent-a-car/pkg/repository/entadp RepositoryInterface
package entadp

import (
	"context"
	"rent-a-car/pkg/repository/dto"
)

type RepositoryInterface interface {
	User() UserRepository
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, c *dto.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByID(ctx context.Context, id int) (*dto.User, error)
	UpdatedUser(ctx context.Context, id int, c *dto.UserUpdate) error
}
