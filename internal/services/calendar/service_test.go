package calendar

import (
	"errors"
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/model"
	"testing"
)

const (
	correctUser       = "user"
	correctHash       = "$2a$08$SLcWO3xeVzk0Z6.lNreRhuRf7YSeXT9RlfrVDBf96coxKbZYfobW6"
	correctPassword   = "testpassword"
	correctToken      = "correct_token"
	incorrectUser     = "incorrect"
	incorrectPassword = "incorrect_password"
	incorrectHash     = "$2a$08$SLcWO3xeVzk0Z6.lNreRhuRf7YSeXT9RlfrVDBf96coxKbZYfoeee"
	tokenFailsUser    = "token_generating_fails"
)

type RepositoryMock struct {
}

func (r RepositoryMock) CreateUser(login, password, timezone string) error {
	return nil
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := model.Auth{
				Login:    tt.login,
				Password: tt.password,
			}

			token, err := service.Login(user, jwt)

			if token != tt.token {
				t.Errorf("token is wrong, got %q want %q", token, tt.token)
			}

			if !errors.Is(err, tt.error) {
				t.Errorf("error is worng, got '%T' want '%T'", err, tt.error)
			}
		})
	}
}
