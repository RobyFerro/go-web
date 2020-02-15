package service

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework"
	"github.com/go-redis/redis/v7"
)

// Connect to Redis
func ConnectRedis(conf *gwf.Conf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		gwf.ProcessError(err)
	}

	return client
}
