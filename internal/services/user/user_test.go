package user

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/Selepok/calendar/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:generate mockgen -destination=mock_repository.go -source=user.go -package=user
const (
	correctUserId      = 1
	correctUser        = "user"
	correctHash        = "$2a$08$.UdKqUQmZqdPm61PtDvKTuukViGD9Xn6Od1wK0RFkNnJwrfXL5IE."
	correctPassword    = "testpassword"
	correctToken       = "correct_token"
	correctTimezone    = "Europe/Kyiv"
	incorrectUserId    = 2
	incorrectUser      = "incorrect"
	incorrectPassword  = "incorrect_password"
	incorrectHash      = "$2a$08$nnqKRvypmsZkzE4B3JQDB.5ajutXXQOVq73jDdLPgwN/Oq4RTrrrr"
	tokenFailsUserId   = 3
	tokenFailsUser     = "token_generating_fails"
	errorUserId        = 0
	forbiddenUserId    = 4
	forbiddenLogin     = "forbidden"
	internalErrorLogin = "internal"
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
			name:     "CreateUser create success",
			password: correctPassword,
			login:    correctUser,
			timezone: correctTimezone,
			error:    nil,
		},
		{
			name:     "CreateUser create error",
			password: incorrectPassword,
			login:    incorrectUser,
			timezone: "",
			error:    errors2.UserCreationIssue{},
		},
	}
	assertion := assert.New(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := model.CreateUser{Login: tt.login, Password: tt.password, TimeZone: tt.timezone}

			err := service.CreateUser(user)
			assertion.Equalf(tt.error, err, "Test case: %s", tt.name)
		})
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	repositoryMock.EXPECT().GetUserHashedPassword(correctUser).Return(correctUserId, correctHash, nil)
	repositoryMock.EXPECT().GetUserHashedPassword(tokenFailsUser).Return(tokenFailsUserId, correctHash, nil)
	repositoryMock.EXPECT().GetUserHashedPassword(incorrectUser).Return(incorrectUserId, incorrectHash, nil)
	repositoryMock.EXPECT().GetUserHashedPassword("").Return(incorrectUserId, "", errors2.NoUserFound(""))

	service := Service{repositoryMock}

	tokenAuthMock := auth.NewMockTokenAuthentication(ctrl)
	tokenAuthMock.EXPECT().GenerateToken(correctUserId).Return(correctToken, nil)
	tokenAuthMock.EXPECT().GenerateToken(tokenFailsUserId).Return("", errors2.GenerateTokenIssue{})

	tests := []struct {
		name     string
		login    string
		password string
		error    error
		token    string
	}{
		{
			name:     "CreateUser not found",
			login:    "",
			password: "",
			error:    errors2.NoUserFound(""),
			token:    "",
		},
		{
			name:     "CreateUser with incorrect password",
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
			user := model.Login{
				Login:    tt.login,
				Password: tt.password,
			}

			token, err := service.Login(user, tokenAuthMock)

			assertion.Equalf(tt.token, token, "Test case: %s", tt.name)
			assertion.Equalf(tt.error, err, "Test case: %s", tt.name)
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	errorUser := &model.User{Login: incorrectUser}
	forbiddenUser := &model.User{Id: correctUserId, Login: forbiddenLogin}
	internalErrorUser := &model.User{Id: correctUserId, Login: internalErrorLogin}
	successUser := &model.User{Id: correctUserId, Login: correctUser}
	repositoryMock := NewMockRepository(ctrl)
	repositoryMock.EXPECT().GetUserIdByLogin(errorUser.Login).Return(errorUserId, errors2.NoUserFound(errorUser.Login))
	repositoryMock.EXPECT().GetUserIdByLogin(forbiddenUser.Login).Return(forbiddenUserId, nil)
	repositoryMock.EXPECT().GetUserIdByLogin(internalErrorUser.Login).Return(internalErrorUser.Id, nil)
	repositoryMock.EXPECT().GetUserIdByLogin(successUser.Login).Return(successUser.Id, nil)
	repositoryMock.EXPECT().Update(*successUser).Return(nil)
	repositoryMock.EXPECT().Update(*internalErrorUser).Return(&errors2.AccessForbidden{})
	assertion := assert.New(t)

	service := Service{repositoryMock}
	tests := []struct {
		name     string
		user     *model.User
		timezone string
		error    error
	}{
		{
			name:  "User not found",
			user:  errorUser,
			error: errors2.NoUserFound(errorUser.Login),
		},
		{
			name:  "Access forbidden",
			user:  forbiddenUser,
			error: &errors2.AccessForbidden{},
		},
		{
			name:  "Internal Server error",
			user:  internalErrorUser,
			error: &errors2.InternalServerError{},
		},
		{
			name:  "Success update",
			user:  successUser,
			error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.Update(*tt.user)
			assertion.Equalf(tt.error, err, "Test case: %s", tt.name)
		})
	}
}
