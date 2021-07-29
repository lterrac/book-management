package rest

import (
	"book-management/pkg/server/pkg/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BookServer is the REST server for the Book and Collections API
type BookServer struct {
	server http.Server
	db     db.Handler
}

// NewBookServer returns a new BookServer
func NewBookServer() *BookServer {
	b := &BookServer{}

	dbOpts := db.NewDBOptions()

	handler, err := db.NewMySQLHandler(dbOpts)

	if err != nil {
		log.Fatal(err)
	}

	b.db = handler

	b.setupRESTSHandlers()
	return b
}

// Start setup the DB connection and http handlers
func (s *BookServer) Start() error {
	fmt.Printf("starting server on %v", s.server.Addr)
	return s.server.ListenAndServe()
}

// setupRESTSHandlers setup REST handlers
func (s *BookServer) setupRESTSHandlers() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/books", s.handleBookModifications).
		Methods(http.MethodPost, http.MethodPut)

	subrouter.NewRoute().Subrouter().Queries("filter", "{filter:[0-9|a-z|_|-]*}").
		Methods(http.MethodDelete, http.MethodGet).
		HandlerFunc(s.handleBookRetrieval)
	s.server = http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
}
