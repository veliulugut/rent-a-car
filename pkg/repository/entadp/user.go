package entadp

import (
	"context"
	"fmt"
	"rent-a-car/ent"
	"rent-a-car/pkg/repository/dto"
	"rent-a-car/pkg/repository/entadp/helper"
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
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("entadp user / get user by id :%w", ErrorIsNotFound)
		}
		return nil, fmt.Errorf("entadp user / get user by id :%w", err)
	}

	return helper.DbUserToDTO(dbUser), nil
}

func (u *UserRepository) UpdatedUser(ctx context.Context, id int, c *dto.User) error {
	_, err := u.DbClient.User.UpdateOneID(id).
		SetEmail(c.Email).
		SetPassword(c.Password).
		SetPhoneNumber(c.PhoneNumber).
		SetUsername(c.UserName).
		SetFirstName(c.FirstName).
		SetLastName(c.LastName).
		SetUpdatedAt(c.UpdateAt).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("entadp user / get updated user :%w", ErrorIsNotFound)
		}
		return fmt.Errorf("entadp user / updated user :%w", err)
	}

	return nil
}
