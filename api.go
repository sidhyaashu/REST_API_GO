package main

import (
	"log" // Importing the log package for logging messages
	"net/http" // Importing the net/http package for handling HTTP requests
	"github.com/gorilla/mux" // Importing the Gorilla Mux router for managing HTTP routes
)

// APIServer struct represents the HTTP server that serves the API.
// It holds the server's address and the store interface for data access.
type APIServer struct {
	addr  string // addr holds the address where the API server will listen (e.g., "localhost:8080")
	store Store // store represents the data store interface, used for interacting with the database
}

// NewAPIServer is a constructor function that creates and returns a new instance of APIServer.
// It takes an address string (addr) and a Store (used to interact with the data storage) as parameters.
func NewAPIServer(addr string, store Store) *APIServer {
	// Returning a new APIServer instance with the provided address and store
	return &APIServer{
		addr:  addr, // Set the address for the server
		store: store, // Set the store for interacting with the database
	}
}

// Serve is a method on APIServer that starts the HTTP server and begins handling incoming requests.
// It initializes the router, registers the routes, and starts the HTTP server listening on the given address.
func (s *APIServer) Serve() {
	// Create a new router instance using Gorilla Mux
	router := mux.NewRouter()

	// Create a subrouter that handles routes prefixed with "/api/v1"
	subrouter := router.PathPrefix("/api/v1").Subrouter()


	//Register Our Service
	taskService := NewTasksService(s.store)
	taskService.RegisterRoutes(router)

	// Log a message indicating the server is starting and display the address it's listening on
	log.Println("Starting API at ", s.addr)

	// Start the HTTP server, passing the server's address and the subrouter for API routes
	// Log and terminate the program if the server encounters an error
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
