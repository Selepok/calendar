package calendar

import (
	//"github.com/Selepok/calendar/internal/repository/postgre"
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
	s.repo.CreateUser(
		credentials.Login,
		string(hashedPassword),
		credentials.Timezone,
	)
	return nil
}

func (s *Service) Login(credentials http.Credentials) bool {
	hashedPassword, err := s.repo.GetUserHashedPassword(credentials.Login)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		return false
	}

	return true
}
