package user

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/Selepok/calendar/internal/model"
	"github.com/Selepok/calendar/internal/server/http"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(login, password, timezone string) error
	GetUserHashedPassword(login string) (hashedPassword string, err error)
}

// Service holds calendar business logic and works with repository
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(credentials http.Credentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 8)
	if err != nil {
		return err
	}
	return s.repo.CreateUser(
		credentials.Login,
		string(hashedPassword),
		credentials.Timezone,
	)
}

func (s *Service) Login(credentials model.Auth, jwt auth.Auth) (token string, err error) {
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