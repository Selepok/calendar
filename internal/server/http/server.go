package http

import (
	"github.com/Selepok/calendar/internal/config"
	errors2 "github.com/Selepok/calendar/internal/errors"
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/Selepok/calendar/internal/model"
	"github.com/Selepok/calendar/internal/response"
	"github.com/Selepok/calendar/internal/services/calendar"
	"github.com/Selepok/calendar/internal/services/validator"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type Validator interface {
	Validate(io.Reader, validator.Ok) error
}

type UserService interface {
	CreateUser(user model.CreateUser) error
	Login(model.Login, auth.TokenAuthentication) (string, error)
	Update(user model.User) error
}

type Server struct {
	valid    Validator
	user     UserService
	calendar calendar.Calendar
	config   config.Application
}

// TODO: add Logger middleware . Good example Zap
// TODO: gorilla mux add middleware once
func NewServer(valid Validator, user UserService, calendar calendar.Calendar, config config.Application) *Server {
	return &Server{valid: valid, user: user, calendar: calendar, config: config}
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.CreateUser{}
	if err := s.valid.Validate(r.Body, &user); err != nil {
		response.Respond(w, http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	if err := s.user.CreateUser(user); err != nil {
		response.Respond(w, http.StatusUnauthorized, response.Error{Error: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, response.Message{Message: "User has been created successfully."})
	return
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	user := model.Login{}
	if err := s.valid.Validate(r.Body, &user); err != nil {
		response.Respond(w, http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	jwt := &auth.JwtWrapper{
		SecretKey:         s.config.SecretKey,
		ExpirationMinutes: s.config.TokenExpirationDuration,
	}

	token, err := s.user.Login(user, jwt)
	if err != nil {
		response.Respond(w, http.StatusUnauthorized, response.Error{Error: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, response.Token{Token: token})
	return
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := s.valid.Validate(r.Body, &user); err != nil {
		response.Respond(w, http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	user.Id = s.getUserIdFromToken(r.Header.Get("Authorization"))

	// TODO: Add 403 error
	if err := s.user.Update(user); err != nil {
		response.Respond(w, http.StatusInternalServerError, response.Error{Error: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, response.Message{Message: "User has been successfully updated."})
	return
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := model.Event{}
	if err := s.valid.Validate(r.Body, &event); err != nil {

		response.Respond(w, http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	event.UserId = s.getUserIdFromToken(r.Header.Get("Authorization"))

	if err := s.calendar.CreateEvent(event); err != nil {
		response.Respond(w, http.StatusInternalServerError, response.Error{Error: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, event)
}

func (s *Server) GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response.Respond(w, http.StatusBadRequest, response.Error{Error: errors2.IncorrectParam(params["id"]).Error()})
		return
	}

	//login := s.getUserIdFromToken(use request
	//r.Header.Get("Authorization"))

	//if err := s.verifyUser(login); err != nil {
	//	err := errors2.AccessForbidden{}
	//	response.Respond(w, http.StatusForbidden, response.Error{Error: err.Error()})
	//	return
	//}
	//user.Id = s.getUserIdFromToken(r.Header.Get("Authorization"))

	event, err := s.calendar.GetEvent(id, s.getUserIdFromToken(r.Header.Get("Authorization")))
	if err == errors2.NoEventFound(id) {
		response.Respond(w, http.StatusNotFound, response.Error{Error: err.Error()})
		return
	}
	if err != nil {
		response.Respond(w, http.StatusInternalServerError, response.Error{Error: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, event)
}

func (s *Server) getUserIdFromToken(token string) (userId int) {
	jwt := &auth.JwtWrapper{
		SecretKey:         s.config.SecretKey,
		ExpirationMinutes: s.config.TokenExpirationDuration,
	}

	userId = jwt.GetUserIdFromToken(token)
	return
}

func (s *Server) verifyUser(login string) error {
	//s.user.
	return nil
}
