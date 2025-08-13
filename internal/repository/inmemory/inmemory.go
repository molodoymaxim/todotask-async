package inmemory

import (
	"errors"
	"github.com/molodoymaxim/todotask-async.git/internal/types"
	"sync"
	"sync/atomic"
)

type TaskRepository interface {
	GetAll(statusFilter string) ([]types.Task, error)
	GetByID(id int) (types.Task, error)
	Create(task types.Task) (types.Task, error)
}

type InMemoryTaskRepository struct {
	tasks  map[int]types.Task
	mu     sync.RWMutex
	nextID atomic.Int32
}

func NewInMemoryTaskRepository() TaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[int]types.Task),
	}
}

func (r *InMemoryTaskRepository) GetAll(statusFilter string) ([]types.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []types.Task
	for _, task := range r.tasks {
		if statusFilter == "" || task.Status == statusFilter {
			result = append(result, task)
		}
	}
	return result, nil
}

func (r *InMemoryTaskRepository) GetByID(id int) (types.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return types.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepository) Create(task types.Task) (types.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int(r.nextID.Add(1))
	task.ID = id
	r.tasks[id] = task
	return task, nil
}
