package service

import (
	"fmt"
	"github.com/RobyFerro/go-web/config"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/gommon/log"
)

// ConnectRedis connect to Redis
func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisConf.Host, config.RedisConf.Port),
		Password: config.RedisConf.Password,
		DB:       1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Error(err)
	}

	return client
}
