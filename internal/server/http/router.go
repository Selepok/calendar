package http

import (
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/gorilla/mux"
	"net/http"
)

// will hold http routes and will registrate them

func NewRouter(server *Server) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/signup", http.HandlerFunc(server.HandlerA))
	router.Handle("/login", http.HandlerFunc(server.Login))
	router.Handle("/test", auth.MiddlewareAuth(http.HandlerFunc(server.Test)))

	return router
}
