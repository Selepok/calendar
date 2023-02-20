package http

import (
	"github.com/Selepok/calendar/internal/middleware/auth"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(server *Server) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/signup", http.HandlerFunc(server.CreateUser)).Methods(http.MethodPost)
	router.Handle("/login", http.HandlerFunc(server.Login)).Methods(http.MethodPost)
	router.Handle("/api/user", auth.ValidateToken(http.HandlerFunc(server.UpdateUser))).Methods(http.MethodPut)
	router.Handle("/api/events", auth.ValidateToken(http.HandlerFunc(server.CreateEvent))).Methods(http.MethodPost)

	return router
}
