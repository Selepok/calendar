package http

import (
	"encoding/json"
	"github.com/Selepok/calendar/internal/model"
	"github.com/Selepok/calendar/internal/response"
	"github.com/Selepok/calendar/internal/services/validator"
	"io"
	"net/http"
)

type Validator interface {
	Validate(io.Reader, validator.Ok) error
}

type Service interface {
	CreateUser(credentials Credentials) error
	Login(model.Auth) (string, error)
}

type Server struct {
	valid    Validator
	calendar Service
}

type Credentials struct {
	Password string `json:"password", db:"password"`
	Login    string `json:"login", db:"login"`
	Timezone string `json:"timezone", db:"timezone"`
}

type Login struct {
	Password string `json:"password", db:"password"`
	Login    string `json:"login", db:"login"`
}

func NewServer(valid Validator, calendar Service) *Server {
	return &Server{valid: valid, calendar: calendar}
}

func (s *Server) HandlerA(w http.ResponseWriter, r *http.Request) {
	// TODO: unmarshall
	// TODO: validate
	// TODO call service
	credentials := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.calendar.CreateUser(*credentials)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	user := model.Auth{}
	err := s.valid.Validate(r.Body, &user)

	if err != nil {
		response.Respond(w, http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	token, err := s.calendar.Login(user)
	if err != nil {
		response.Respond(w, http.StatusUnauthorized, response.Error{Error: err.Error()})
		return
	}

	response.Respond(w, http.StatusOK, response.Token{Token: token})
	return
}

func (s *Server) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good!"))
	w.WriteHeader(http.StatusOK)
	return
}
