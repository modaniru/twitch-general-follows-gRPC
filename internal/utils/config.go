package utils

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	defaultPort = "8080"
)

// Load .env file if exist
func LoadEnvIfExist() {
	// Если у нас есть .env файл, загружаем переменные с него, иначе они должны быть поставлены другим способом
	// Не создаю абстракцию конфигурации, чтобы не хранить secrets in memory
	_ = godotenv.Load()
}

// Return port from .env file else returning 8080
func GetPort() string {
	p := os.Getenv("PORT")
	if len(p) == 0 {
		return defaultPort
	}
	return p
}
