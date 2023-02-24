package calendar

import (
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/model"
	"github.com/Selepok/calendar/internal/services/user"
	"log"
)

type Repository interface {
	Create(event model.Event) (int, error)
	Get(id int) (event model.Event, err error)
}

// Service holds calendar business logic and works with repository
type Service struct {
	repo Repository
	user user.Repository
}

type Calendar interface {
	CreateEvent(event model.Event) error
	GetEvent(eventId int, userId int) (model.Event, error)
}

func NewEventsService(repo Repository, user user.Repository) *Service {
	return &Service{repo: repo, user: user}
}

func (s *Service) CreateEvent(event model.Event) error {
	eventId, err := s.repo.Create(event)
	if err != nil {
		return err
	}
	log.Println(eventId)
	return nil
}

func (s *Service) GetEvent(eventId int, userId int) (event model.Event, err error) {
	event, err = s.repo.Get(eventId)
	if err != nil {
		return
	}
	if event.UserId != userId {
		return model.Event{}, &errors2.AccessForbidden{}
	}
	
	return
}
