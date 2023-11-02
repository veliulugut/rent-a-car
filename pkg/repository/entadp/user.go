package entadp

import (
	"context"
	"fmt"
	"rent-a-car/ent"
	"rent-a-car/pkg/repository/dto"
	"rent-a-car/pkg/repository/helper"
	"time"
)

var _ UserRepositoryInterface = (*UserRepository)(nil)

func NewUserRepository(DbClient *ent.Client) *UserRepository {
	return &UserRepository{
		DbClient: DbClient,
	}
}

type UserRepository struct {
	DbClient *ent.Client
}

func (u *UserRepository) CreateUser(ctx context.Context, c *dto.User) error {
	_, err := u.DbClient.User.Create().SetFirstName(c.FirstName).SetLastName(c.LastName).SetUsername(c.UserName).
		SetEmail(c.Email).SetPhoneNumber(c.PhoneNumber).SetPassword(c.Password).SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return fmt.Errorf("entadp user / create user :%w", err)
	}

	return nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int) error {
	if err := u.DbClient.User.DeleteOneID(id).Exec(ctx); err != nil {
		return fmt.Errorf("entadp user / delete user :%w", err)
	}

	return nil

}

func (u *UserRepository) GetUserByID(ctx context.Context, id int) (*dto.User, error) {
	var (
		dbUser *ent.User
		err    error
	)

	if dbUser, err = u.DbClient.User.Get(ctx, id); err != nil {
		return nil, fmt.Errorf("entadp user / get user by id :%w", err)
	}

	return helper.DbUserToDTO(dbUser), nil
}
