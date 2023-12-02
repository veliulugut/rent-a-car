package user

import (
	"context"
	"fmt"
	"rent-a-car/pkg/passwd"
	"rent-a-car/pkg/repository/dto"
	"rent-a-car/pkg/repository/entadp"
	"rent-a-car/pkg/validator"
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
	var err error

	emailVld := validator.IsMailValid{Email: c.Email}
	passVld := validator.IsInRange{
		Text: c.Password,
		Min:  6,
		Max:  20,
	}

	if err = validator.ValidateAll(&emailVld, &passVld); err != nil {
		return fmt.Errorf("service / create user :%w", err)
	}

	var hashedPw string

	if hashedPw, err = u.pw.Generate(c.Password); err != nil {
		return fmt.Errorf("create user / generate :%w", err)
	}

	dtoUser := dto.User{
		UserName: c.UserName,
		Email:    c.Email,
		Password: hashedPw,
		CreateAt: c.CreatedAt,
	}

	userRepo := u.repo.User()
	if err = userRepo.CreateUser(ctx, &dtoUser); err != nil {
		return fmt.Errorf("create user / user :%w", err)
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, id int) error {

	user := u.repo.User()
	if err := user.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user / user:%w", err)
	}

	return nil
}

func (u *User) UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error {
	dtoUser := dto.UserUpdate{
		UserName: c.UserName,
		Email:    c.Email,
		Password: c.Password,
		UpdateAt: c.UpdateAt,
	}

	user := u.repo.User()
	if err := user.UpdatedUser(ctx, id, &dtoUser); err != nil {
		return fmt.Errorf("update user / user :%w", err)
	}

	return nil
}

func (u *User) GetUserByID(ctx context.Context, id int) (*GetUserModel, error) {
	var (
		user *dto.User
		err  error
	)
	usr := u.repo.User()
	if user, err = usr.GetUserByID(ctx, id); err != nil {
		return nil, fmt.Errorf("get user by id / user: %w", err)
	}

	return &GetUserModel{
		Email:    user.Email,
		UserName: user.UserName,
	}, nil
}
