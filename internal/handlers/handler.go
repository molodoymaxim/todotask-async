package handlers

import (
	"encoding/json"
	"errors"
	"github.com/molodoymaxim/todotask-async.git/internal/service"
	"github.com/molodoymaxim/todotask-async.git/internal/types"
	"net/http"
	"strconv"
	"strings"
)

type TaskHandler interface {
	GetTasks(w http.ResponseWriter, r *http.Request)
	GetTaskByID(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	srv service.TaskService
}

func NewTaskHandler(srv service.TaskService) TaskHandler {
	return &taskHandler{srv: srv}
}

func (h *taskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	statusFilter := r.URL.Query().Get("status")
	tasks, err := h.srv.GetAll(statusFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *taskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	idStr := strings.Trim(path, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	task, err := h.srv.GetByID(id)
	if err != nil {
		if errors.Is(err, errors.New("task not found")) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task types.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdTask, err := h.srv.Create(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}
