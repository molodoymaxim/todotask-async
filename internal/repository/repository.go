package repository

import (
	"github.com/molodoymaxim/todotask-async.git/internal/repository/inmemory"
)

type Repositories struct {
	Task inmemory.TaskRepository
}

func New() *Repositories {
	return &Repositories{
		Task: inmemory.NewInMemoryTaskRepository(),
	}
}
