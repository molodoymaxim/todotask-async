package server

import (
	"context"
	"fmt"
	"github.com/molodoymaxim/todotask-async.git/internal/handlers"
	"github.com/molodoymaxim/todotask-async.git/internal/types"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
	logChan    chan<- string
	wg         *sync.WaitGroup
}

// NewServer создает новый экземпляр Server
func NewServer(cfg *types.ConfigApp, h handlers.TaskHandler, logChan chan<- string, wg *sync.WaitGroup) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", h.GetTasks)
	mux.HandleFunc("GET /tasks/{id}", h.GetTaskByID)
	mux.HandleFunc("POST /tasks", h.CreateTask)

	addr := fmt.Sprintf(":%d", cfg.Port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &Server{
		httpServer: httpServer,
		logChan:    logChan,
		wg:         wg,
	}
}

// Start запускает сервер и обрабатывает graceful shutdown
func (s *Server) Start() error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
		}
	}()
	fmt.Printf("Server started on %s\n", s.httpServer.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	return s.Shutdown()
}

// Shutdown выполняет graceful shutdown сервера
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	close(s.logChan)
	s.wg.Wait()

	fmt.Println("Server exited gracefully")
	return nil
}
