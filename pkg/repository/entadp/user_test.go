package entadp

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"rent-a-car/ent"
	"rent-a-car/ent/enttest"
	"rent-a-car/pkg/repository/dto"
	"testing"
	"time"
)

func TestUserRepository_CreateUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	users := []dto.User{
		{
			FirstName:   "testVeli",
			LastName:    "Ulugut",
			UserName:    "abcabc",
			Email:       "test@mail.com",
			PhoneNumber: "555 555 555",
			Password:    "testPas",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		},
		{
			FirstName:   "testCasPer",
			LastName:    "Jack",
			UserName:    "abc abcoglu",
			Email:       "tests@mail.com",
			PhoneNumber: "222 222 222",
			Password:    "123123123",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		},
	}

	for i := range users {
		t.Run(fmt.Sprintf("CreateUser_Index:%d", i), func(t *testing.T) {
			if err := userRepo.CreateUser(context.Background(), &users[i]); err != nil {
				t.Error(err)
			}
		})
	}

	t.Run("ContextTimeOut", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*0)
		defer cancel()

		if err := userRepo.CreateUser(ctx, &dto.User{}); err == nil {
			t.Error(err)
		}
	})

	t.Run("ContextCancel", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
		cancel()

		if err := userRepo.CreateUser(ctx, &dto.User{}); err == nil {
			t.Error(err)
		}
	})

	/*t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("DuplicateUser", func(t *testing.T) {
			err := userRepo.CreateUser(context.Background(), &dto.User{
				FirstName:   "testVeli",
				LastName:    "Ulugut",
				UserName:    "asdd",
				Email:       "asd@gmail.com",
				PhoneNumber: "31 31 31 31",
				Password:    "123123123",
				CreateAt:    time.Now(),
				UpdateAt:    time.Now(),
			})
			if err == nil {
				t.Error(err)
			}

		})
	})*/
}
