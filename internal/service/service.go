package service

import (
	"fmt"
	"github.com/molodoymaxim/todotask-async.git/internal/repository"
	"github.com/molodoymaxim/todotask-async.git/internal/types"
)

type TaskService interface {
	GetAll(statusFilter string) ([]types.Task, error)
	GetByID(id int) (types.Task, error)
	Create(task types.Task) (types.Task, error)
}

type taskService struct {
	repo    repository.Repositories
	logChan chan<- string
}

func NewTaskService(repo repository.Repositories, logChan chan<- string) TaskService {
	return &taskService{repo: repo, logChan: logChan}
}

func (s *taskService) GetAll(statusFilter string) ([]types.Task, error) {
	tasks, err := s.repo.Task.GetAll(statusFilter)
	if err != nil {
		return nil, err
	}
	s.logAsync(fmt.Sprintf("Action: GetAllTasks, Filter: %s, Count: %d", statusFilter, len(tasks)))
	return tasks, nil
}

func (s *taskService) GetByID(id int) (types.Task, error) {
	task, err := s.repo.Task.GetByID(id)
	if err != nil {
		return types.Task{}, err
	}
	s.logAsync(fmt.Sprintf("Action: GetTaskByID, ID: %d", id))
	return task, nil
}

func (s *taskService) Create(task types.Task) (types.Task, error) {
	createdTask, err := s.repo.Task.Create(task)
	if err != nil {
		return types.Task{}, err
	}
	s.logAsync(fmt.Sprintf("Action: CreateTask, ID: %d, Title: %s, Status: %s", createdTask.ID, createdTask.Title, createdTask.Status))
	return createdTask, nil
}

func (s *taskService) logAsync(msg string) {
	go func() {
		s.logChan <- msg
	}()
}
