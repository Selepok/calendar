package user

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/model"
	"github.com/Selepok/calendar/internal/server/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	correctUser       = "user"
	correctHash       = "$2a$08$SLcWO3xeVzk0Z6.lNreRhuRf7YSeXT9RlfrVDBf96coxKbZYfobW6"
	correctPassword   = "testpassword"
	correctToken      = "correct_token"
	correctTimezone   = "Europe/Kyiv"
	incorrectUser     = "incorrect"
	incorrectPassword = "incorrect_password"
	incorrectHash     = "$2a$08$SLcWO3xeVzk0Z6.lNreRhuRf7YSeXT9RlfrVDBf96coxKbZYfoeee"
	tokenFailsUser    = "token_generating_fails"
)

type RepositoryMock struct {
}

func (r RepositoryMock) CreateUser(login, password, timezone string) error {
	if login == correctUser && timezone == correctTimezone {
		return nil
	}

	return errors2.UserCreationIssue{}
}

func (r RepositoryMock) GetUserHashedPassword(login string) (hashedPassword string, err error) {
	switch login {
	case correctUser, tokenFailsUser:
		return correctHash, nil
	case incorrectUser:
		return incorrectHash, nil
	default:
		return "", errors2.NoUserFound(login)
	}
}

type JwtMock struct {
}

func (jw *JwtMock) GenerateToken(login string) (tokenString string, err error) {
	switch login {
	case correctUser:
		return correctToken, nil
	case tokenFailsUser:
		return tokenString, errors2.GenerateTokenIssue{}
	default:
		return
	}
}
func (jw *JwtMock) ValidateToken(string) error {
	return nil
}

func TestCreateUser(t *testing.T) {
	repository := RepositoryMock{}
	service := Service{repository}

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
			creds := http.Credentials{
				Password: tt.password,
				Login:    tt.login,
				Timezone: tt.timezone,
			}
			err := service.CreateUser(creds)
			assertion.Equalf(err, tt.error, "Test case: %s", tt.name)
		})
	}
}

func TestLogin(t *testing.T) {
	repository := RepositoryMock{}
	service := Service{repository}

	jwt := &JwtMock{}

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

			token, err := service.Login(user, jwt)

			assertion.Equalf(token, tt.token, "Test case: %s", tt.name)
			assertion.Equalf(err, tt.error, "Test case: %s", tt.name)
		})
	}
}
