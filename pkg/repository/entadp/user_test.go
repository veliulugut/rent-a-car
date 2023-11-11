package entadp

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
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

func TestUserRepository_GetUserByID(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	tests := []struct {
		name         string
		setup        func() (*ent.User, error)
		id           int
		expectedUser *dto.User
		expectedErr  error
	}{
		{
			name: "pass",
			setup: func() (*ent.User, error) {
				return client.User.Create().
					SetUsername("veli").
					SetFirstName("test").
					SetLastName("ulugut").
					SetEmail("test@mail.com").
					SetPassword("555 555 555").
					SetPhoneNumber("555 555 555").
					SetCreatedAt(time.Now()).
					SetUpdatedAt(time.Now()).
					Save(context.Background())
			},
			id: 1,
			expectedUser: &dto.User{
				ID:          1,
				FirstName:   "test",
				LastName:    "ulugut",
				UserName:    "veli",
				Email:       "test@mail.com",
				PhoneNumber: "555 555 555",
				Password:    "555 555 555",
			},
			expectedErr: nil,
		},
		{
			name:         "user not found",
			setup:        func() (*ent.User, error) { return nil, nil },
			id:           31,
			expectedUser: nil,
			expectedErr:  ErrorIsNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			_, setupErr := tt.setup()
			if setupErr != nil {
				t.Fatalf("Setup error: %v", setupErr)
			}

			// Test
			resultUser, err := userRepo.GetUserByID(context.Background(), tt.id)

			if !reflect.DeepEqual(resultUser, tt.expectedUser) {
				t.Errorf("Expected user: %+v, but got: %+v", tt.expectedUser, resultUser)
			}

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedErr, err)
			}
		})
	}
}

func TestUserRepository_DeleteUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := []dto.User{
		{
			ID:          1,
			FirstName:   "test",
			LastName:    "testUser",
			UserName:    "user",
			Email:       "test@mail.com",
			PhoneNumber: "555 555",
			Password:    "asdfg",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		},
	}

	for i := range user {
		t.Run(fmt.Sprintf("CreateUser_%d", i), func(t *testing.T) {
			if err := userRepo.CreateUser(context.Background(), &user[i]); err != nil {
				t.Error(err)
			}
			t.Run(fmt.Sprintf("DeleteUser_%d", i), func(t *testing.T) {
				if err := userRepo.DeleteUser(context.Background(), 1); err != nil {
					t.Error(err)
				}
			})
		})
	}
}

func TestUserRepository_UpdatedUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := []dto.User{
		{
			ID:          1,
			FirstName:   "veli",
			LastName:    "ulugut",
			UserName:    "testUser",
			Email:       "test@user.com",
			PhoneNumber: "222 222",
			Password:    "123123",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		},
	}

	for i := range user {
		t.Run(fmt.Sprintf("CreateUser_Index:%d", user[i].ID), func(t *testing.T) {
			if err := userRepo.CreateUser(context.Background(), &user[i]); err != nil {
				t.Error(err)
			}
		})

		t.Run(fmt.Sprintf("DeleteUser_Index:%d", user[i].ID), func(t *testing.T) {
			if err := userRepo.DeleteUser(context.Background(), 1); err != nil {
				t.Error(err)
			}
		})
	}
}
