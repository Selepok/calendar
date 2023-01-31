package http

import (
	"net/http"
)

type Validator interface {
	Validate(interface{}) error
}

type Service interface {
}

type Server struct {
	valid    Validator
	calendar Service
}

func NewServer(valid Validator, calendar Service) *Server {
	return &Server{valid: valid, calendar: calendar}
}

func (s *Server) HandlerA(w http.ResponseWriter, r *http.Request) {
	// TODO: unmarshall
	// TODO: validate
	// TODO call service
}
