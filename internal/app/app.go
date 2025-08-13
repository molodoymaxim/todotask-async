package app

import (
	"github.com/molodoymaxim/todotask-async.git/internal/config"
	"github.com/molodoymaxim/todotask-async.git/internal/handlers"
	"github.com/molodoymaxim/todotask-async.git/internal/repository"
	"github.com/molodoymaxim/todotask-async.git/internal/server"
	"github.com/molodoymaxim/todotask-async.git/internal/service"
	"log"
	"sync"
)

func Start() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация репозитория
	repo := repository.New()

	// Инициализация канала для логов и горутины-логгера
	logChan := make(chan string, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range logChan {
			log.Println(msg)
		}
	}()

	// Инициализация сервиса
	srv := service.NewTaskService(*repo, logChan)

	// Инициализация хендлера
	h := handlers.NewTaskHandler(srv)

	// Инициализация сервера
	server := server.NewServer(cfg, h, logChan, &wg)

	// Запуск сервера
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
