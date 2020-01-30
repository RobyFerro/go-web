package database

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"ikdev/go-web/app/config"
	"ikdev/go-web/exception"
)

// Connect to Redis
func ConnectRedis(conf config.Conf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		exception.ProcessError(err)
	}

	return client
}
