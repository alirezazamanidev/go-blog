package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string 
	DbPort string
	DbHost string
	DbUser string
	DbPass string
	DbName string
	RedisHost string
	RedisPort string
}

func Load() (*Config){

	err:=godotenv.Load();
	if err !=nil {
		log.Fatal(err)
	}
	config:=Config{
		AppPort: os.Getenv("APP_PORT"),
		DbPort: os.Getenv("DB_PORT"),
		DbPass: os.Getenv("DB_PASS"),
		DbHost: os.Getenv("DB_HOST"),
		DbUser: os.Getenv("DB_USER"),
		DbName: os.Getenv("DB_NAME"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}

	return &config
}