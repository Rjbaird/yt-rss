package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT            string `env:"PORT" envDefault:"3000"`
	REDIS_URL       string `env:"REDIS_URL"`
	YOUTUBE_API_KEY string `env:"YOUTUBE_API_KEY"`
}

func ENV() (*Config, error) {
	godotenv.Load(".env")

	PORT := os.Getenv("PORT")

	REDIS_URL := os.Getenv("REDIS_URL")
	if REDIS_URL == "" {
		log.Fatal("You must set your 'REDIS_URL' environment variable.")
	}

	YOUTUBE_API_KEY := os.Getenv("YOUTUBE_API_KEY")
	if YOUTUBE_API_KEY == "" {
		log.Fatal("You must set your 'YOUTUBE_API_KEY' environment variable.")
	}

	config := Config{PORT: PORT, REDIS_URL: REDIS_URL, YOUTUBE_API_KEY: YOUTUBE_API_KEY}

	return &config, nil
}
