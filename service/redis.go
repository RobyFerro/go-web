package service

import (
	"fmt"
	"github.com/RobyFerro/go-web/app"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/gommon/log"
)

// ConnectRedis connect to Redis
func ConnectRedis(conf *app.Conf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Error(err)
	}

	return client
}
