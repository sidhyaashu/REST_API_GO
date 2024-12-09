package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http" 

	"github.com/gorilla/mux" 
)

var errNameRequired = errors.New("name is Required")
var errProjectIdRequired = errors.New("project ID is Required")
var errUserIDRequired = errors.New("user ID is Required")

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s} 
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks",  WithJWTAuth(s.HandleCreateTask, s.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}", WithJWTAuth(s.HandleGetTask, s.store)).Methods("GET")
}

func (s *TasksService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	body,err := io.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w,http.StatusBadRequest,ErrorResponse{
			Error: "Invalid Request Payload!",
		})
		return
	}

	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body,&task)
	if err != nil {
		WriteJSON(w,http.StatusBadRequest,ErrorResponse{
			Error: "Invalid Request Payload",
		})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJSON(w,http.StatusBadRequest,ErrorResponse{
			Error: err.Error(),
		})
		return 
	}

	t,err := s.store.CreateTask(task)
	if err != nil {
		WriteJSON(w,http.StatusInternalServerError,ErrorResponse{
			Error: "Error Creating Task",
		})
		return
	}

	WriteJSON(w,http.StatusCreated,t)


}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}
	if task.ProjectID == 0 {
		return errProjectIdRequired
	}
	if task.AssignedToID == 0 {
		return errUserIDRequired
	}

	return nil

}

func (s *TasksService) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == ""{
		WriteJSON(w,http.StatusBadRequest,ErrorResponse{
			Error: "id is Required",
		})
		return
	}

	t,err := s.store.GetTask(id)
	if err != nil {
		WriteJSON(w,http.StatusInternalServerError,ErrorResponse{
			Error: "task not found",
		})
		return
	}

	WriteJSON(w,http.StatusOK,t)
}
