package utils

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Load ("./" or "./configuration").config.yaml and .env files
func LoadConfig(filePath, fileType string) error {
	viper.AddConfigPath(filePath)
	viper.AddConfigPath("./")
	viper.SetConfigType(fileType)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	// Если у нас есть .env файл, загружаем переменные с него, иначе они должны быть поставлены другим способом
	_ = godotenv.Load()
	return nil
}
