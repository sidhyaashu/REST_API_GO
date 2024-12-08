package main

import (
	"net/http" // Importing the net/http package for HTTP request/response handling
	"github.com/gorilla/mux" // Importing the mux package for routing HTTP requests
)

// TasksService struct represents the service layer for handling tasks.
// It contains a 'store' field of type Store, which will be used to interact with the data store (e.g., database).
type TasksService struct {
	store Store // Store is an interface (or type) representing the data store where tasks are managed
}

// NewTasksService is a constructor function that creates and returns a new instance of TasksService.
// It takes a 'Store' type (which implements the data storage logic) as a parameter.
func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s} // Returns a pointer to a new TasksService with the provided store
}

// RegisterRoutes is a method on TasksService that registers the HTTP routes for task-related operations.
// It takes a *mux.Router as a parameter, which will be used to define and handle HTTP routes.
func (s *TasksService) RegisterRoutes(r *mux.Router) {
	// Register a POST route for creating tasks. When the client sends a POST request to "/tasks",
	// it will trigger the HandleCreateTask method.
	r.HandleFunc("/tasks", s.HandleCreateTask).Methods("POST")

	// Register a GET route for retrieving a task by its ID. The ID is extracted from the URL path parameter.
	// When the client sends a GET request to "/tasks/{id}", it will trigger the HandleGetTask method.
	r.HandleFunc("/tasks/{id}", s.HandleGetTask).Methods("GET")
}

// HandleCreateTask is a placeholder method for handling the creation of a task.
// The method will accept an HTTP POST request and write a response using the http.ResponseWriter.
func (s *TasksService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	// This method is currently empty, but it would contain logic to create a task in the store (e.g., database)
	
}

// HandleGetTask is a placeholder method for handling the retrieval of a task.
// The method will accept an HTTP GET request with a task ID in the URL and write the task's details to the response.
func (s *TasksService) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	// This method is currently empty, but it would contain logic to retrieve a task from the store by its ID
}
