package http

import (
	"net/http"
)

// will hold http routes and will registrate them

func NewRouter(server *Server) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(server.HandlerA))

	return router
}
