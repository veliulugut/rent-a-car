package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"rent-a-car/pkg/passwd/bcrypt"
	"rent-a-car/pkg/repository/entadp"
	"testing"
)

func TestUser_CreateUser(t *testing.T) {
	bc := bcrypt.NewBcrypt("31veli31", 10)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tt := []struct {
		username       string
		email          string
		password       string
		confirmPass    string
		expectedErrMsg string
		mockBehavior   func(mock *entadp.MockUserRepositoryInterface)
	}{
		{
			username:       "Veli123",
			email:          "test@mail.com",
			password:       "123123",
			confirmPass:    "123123",
			expectedErrMsg: "",
			mockBehavior: func(mock *entadp.MockUserRepositoryInterface) {
				mock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
		},

		{
			username:       "absd",
			email:          "abc@gmail.com",
			password:       "2312345",
			confirmPass:    "2312345",
			expectedErrMsg: "username already exist",
			mockBehavior: func(mock *entadp.MockUserRepositoryInterface) {
				mock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	repo := entadp.NewMockRepositoryInterface(ctrl)
	userRepo := entadp.NewMockUserRepositoryInterface(ctrl)

	for _, tc := range tt {
		t.Run(tc.username, func(t *testing.T) {
			tc.mockBehavior(userRepo)
			repo.EXPECT().User().Return(userRepo)
			userRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

			user := New(bc, repo)

			err := user.CreateUser(context.Background(), &CreateUserModel{
				Password:        tc.password,
				Email:           tc.email,
				UserName:        tc.username,
				ConfirmPassword: tc.confirmPass,
			})

			if err != nil {
				if tc.expectedErrMsg != "" {
					if err.Error() != tc.expectedErrMsg {
						t.Errorf("expected error: %q, but got: %q", tc.expectedErrMsg, err.Error())
					}
				} else {
					t.Errorf("Error not expected, but error occurred: %v", err)
				}
			}
		})
	}
}
