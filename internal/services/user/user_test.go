package user

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/Selepok/calendar/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	correctUser       = "user"
	correctHash       = "$2a$08$.UdKqUQmZqdPm61PtDvKTuukViGD9Xn6Od1wK0RFkNnJwrfXL5IE."
	correctPassword   = "testpassword"
	correctToken      = "correct_token"
	correctTimezone   = "Europe/Kyiv"
	incorrectUser     = "incorrect"
	incorrectPassword = "incorrect_password"
	incorrectHash     = "$2a$08$nnqKRvypmsZkzE4B3JQDB.5ajutXXQOVq73jDdLPgwN/Oq4RTrrrr"
	tokenFailsUser    = "token_generating_fails"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	repositoryMock.EXPECT().CreateUser(correctUser, gomock.Any(), correctTimezone).Return(nil)
	repositoryMock.EXPECT().CreateUser(incorrectUser, gomock.Any(), "").Return(errors2.UserCreationIssue{})

	service := Service{repositoryMock}

	tests := []struct {
		name     string
		password string
		login    string
		timezone string
		error    error
	}{
		{
			name:     "User create success",
			password: correctPassword,
			login:    correctUser,
			timezone: correctTimezone,
			error:    nil,
		},
		{
			name:     "User create error",
			password: incorrectPassword,
			login:    incorrectUser,
			timezone: "",
			error:    errors2.UserCreationIssue{},
		},
	}
	assertion := assert.New(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := model.User{Login: tt.login, Password: tt.password, TimeZone: tt.timezone}

			err := service.CreateUser(user)
			assertion.Equalf(err, tt.error, "Test case: %s", tt.name)
		})
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	repositoryMock.EXPECT().GetUserHashedPassword(correctUser).Return(correctHash, nil)
	repositoryMock.EXPECT().GetUserHashedPassword(tokenFailsUser).Return(correctHash, nil)
	repositoryMock.EXPECT().GetUserHashedPassword(incorrectUser).Return(incorrectHash, nil)
	repositoryMock.EXPECT().GetUserHashedPassword("").Return("", errors2.NoUserFound(""))

	service := Service{repositoryMock}

	tokenAuthMock := auth.NewMockTokenAuthentication(ctrl)
	tokenAuthMock.EXPECT().GenerateToken(correctUser).Return(correctToken, nil)
	tokenAuthMock.EXPECT().GenerateToken(tokenFailsUser).Return("", errors2.GenerateTokenIssue{})

	tests := []struct {
		name     string
		login    string
		password string
		error    error
		token    string
	}{
		{
			name:     "User not found.",
			login:    "",
			password: "",
			error:    errors2.NoUserFound(""),
			token:    "",
		},
		{
			name:     "User with incorrect password",
			login:    incorrectUser,
			password: incorrectPassword,
			error:    errors2.IncorrectPassword(incorrectUser),
			token:    "",
		},
		{
			name:     "Token issue",
			login:    tokenFailsUser,
			password: correctPassword,
			error:    errors2.GenerateTokenIssue{},
			token:    "",
		},
		{
			name:     "Success login",
			login:    correctUser,
			password: correctPassword,
			error:    nil,
			token:    correctToken,
		},
	}
	assertion := assert.New(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := model.Auth{
				Login:    tt.login,
				Password: tt.password,
			}

			token, err := service.Login(user, tokenAuthMock)

			assertion.Equalf(token, tt.token, "Test case: %s", tt.name)
			assertion.Equalf(err, tt.error, "Test case: %s", tt.name)
		})
	}
}
