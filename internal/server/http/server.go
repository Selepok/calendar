package http

import (
	"github.com/Selepok/calendar/internal/config"
	"github.com/Selepok/calendar/internal/middleware/auth"
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
	CreateUser(user model.User) error
	Login(model.Auth, auth.TokenAuthentication) (string, error)
}

type Server struct {
	valid  Validator
	user   Service
	config config.Application
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

func NewServer(valid Validator, user Service, config config.Application) *Server {
	return &Server{valid: valid, user: user, config: config}
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
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
	user := model.Auth{}
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

func (s *Server) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good!"))
	w.WriteHeader(http.StatusOK)
	return
}
