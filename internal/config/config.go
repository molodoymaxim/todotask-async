package config

import (
	"github.com/molodoymaxim/todotask-async.git/internal/types"
	"os"
	"strconv"
)

// Получение переменных окружения
func Load() (*types.ConfigApp, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	if port < 1 || port > 65535 {
		return nil, strconv.ErrRange
	}

	return &types.ConfigApp{Port: port}, nil
}
