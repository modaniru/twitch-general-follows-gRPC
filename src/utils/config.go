package utils

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Load ("./" or "./configuration").config.yaml and .env files
func LoadConfig(filePath, fileType string) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	viper.AddConfigPath(filePath)
	viper.AddConfigPath("./")
	viper.SetConfigType(fileType)
	return viper.ReadInConfig()
}
