package user

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/Selepok/calendar/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateUser(login, password, timezone string) error
	GetUserHashedPassword(login string) (id int, hashedPassword string, err error)
	Update(user model.User) error
	GetUserIdByLogin(login string) (id int, err error)
}

// Service holds calendar business logic and works with repository
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(user model.CreateUser) error {
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

func (s *Service) Login(credentials model.Login, jwt auth.TokenAuthentication) (token string, err error) {
	id, hashedPassword, err := s.repo.GetUserHashedPassword(credentials.Login)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		return token, errors2.IncorrectPassword(credentials.Login)
	}

	token, err = jwt.GenerateToken(id)
	if err != nil {
		return token, errors2.GenerateTokenIssue{}
	}

	return
}

//
//func (s *Service) GetUserByLogin(login string) error {
//	if err := s.repo.Update(user.Login, user.TimeZone); err != nil {
//		return &errors2.InternalServerError{}
//	}
//	return nil
//}

func (s *Service) Update(user model.User) (err error) {
	userId, err := s.repo.GetUserIdByLogin(user.Login)
	if err != nil {
		return
	}
	if userId != user.Id {
		return &errors2.AccessForbidden{}
	}
	if err = s.repo.Update(user); err != nil {
		return &errors2.InternalServerError{}
	}
	return nil
}
