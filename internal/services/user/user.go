package user

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/Selepok/calendar/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(login, password, timezone string) error
	GetUserHashedPassword(login string) (hashedPassword string, err error)
	Update(login, timezone string) error
}

// Service holds calendar business logic and works with repository
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	return s.repo.CreateUser(
		user.Login,
		string(hashedPassword),
		user.TimeZone,
	)
}

func (s *Service) Login(credentials model.Auth, jwt auth.TokenAuthentication) (token string, err error) {
	hashedPassword, err := s.repo.GetUserHashedPassword(credentials.Login)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		return token, errors2.IncorrectPassword(credentials.Login)
	}

	token, err = jwt.GenerateToken(credentials.Login)
	if err != nil {
		return token, errors2.GenerateTokenIssue{}
	}

	return
}
