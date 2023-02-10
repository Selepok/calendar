package http

import (
	"encoding/json"
	"github.com/Selepok/calendar/internal/middleware/auth"
	"net/http"
	"os"
	"strconv"
)

type Validator interface {
	Validate(interface{}) error
}

type Service interface {
	CreateUser(credentials Credentials) error
	Login(credentials Credentials) bool
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

	if !s.calendar.Login(*credentials) {
		w.WriteHeader(http.StatusUnauthorized)
	}

	expirationMinutes, err := strconv.ParseInt(os.Getenv("TOKEN_EXPIRATION_TIME_IN_MINUTES"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jwt := auth.JwtWrapper{
		SecretKey:         os.Getenv("SECRET_KEY"),
		ExpirationMinutes: expirationMinutes,
	}
	token, err := jwt.GenerateToken(credentials.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
	w.WriteHeader(http.StatusOK)
	return
}

func (s *Server) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good!"))
	w.WriteHeader(http.StatusOK)
	return
}
