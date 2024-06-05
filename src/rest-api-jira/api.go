package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// holds our api server store the connection to
type APIServer struct {
	addr  string
	store Store
}

// constructor
func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// Method to utilize router

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Starting the API server at", s.addr)

	// registering our services
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
