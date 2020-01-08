package command

import (
	"fmt"
	"ikdev/go-web/config"
	"ikdev/go-web/database"
	"time"
)

var conf config.Conf

// Run jobs in Redis queue
func RunQueue(name string) {
	conf = config.Configuration()
	redis := database.ConnectRedis(conf)
	queue := fmt.Sprintf("queue:%s", name)

	for true {
		tasks, _ := redis.BLPop(5*time.Second, queue).Result()

		if len(tasks) != 2 {
			continue
		}

		// TMP - Get content
		fmt.Println(tasks[1])
	}
}
