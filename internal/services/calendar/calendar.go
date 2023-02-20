package calendar

import (
	"github.com/Selepok/calendar/internal/model"
	"github.com/Selepok/calendar/internal/services/user"
	"log"
)

type Repository interface {
	Create(event model.Event) (int32, error)
}

// Service holds calendar business logic and works with repository
type Service struct {
	repo Repository
	user user.Repository
}

type Calendar interface {
	CreateEvent(event model.Event, login string) error
}

func NewEventsService(repo Repository, user user.Repository) *Service {
	return &Service{repo: repo, user: user}
}

func (s *Service) CreateEvent(event model.Event, login string) error {
	id, err := s.user.GetUserIdByLogin(login)
	if err != nil {
		return err
	}
	event.UserId = id
	eventId, err := s.repo.Create(event)
	if err != nil {
		return err
	}
	log.Println(eventId)
	return nil
}
