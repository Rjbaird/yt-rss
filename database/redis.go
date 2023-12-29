package database

import (
	"log"

	"github.com/bairrya/yt-rss/config"
	"github.com/redis/go-redis/v9"
)

func RedisConnect() (*redis.Client, error) {
	config, err := config.ENV()
	if err != nil {
		log.Printf("Error loading config: %s", err)
		return nil, err
	}
	options, err := redis.ParseURL(config.REDIS_URL)
	if err != nil {
		log.Printf("Error parsing redis url: %s", config.REDIS_URL)
		return nil, err
	}

	client := redis.NewClient(options)

	return client, nil
}
