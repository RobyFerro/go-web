package command

import (
	"encoding/json"
	"fmt"
	"ikdev/go-web/config"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
	"ikdev/go-web/job"
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
		var j job.Job

		if len(tasks) != 2 {
			continue
		}

		if err := json.Unmarshal([]byte(tasks[1]), &j); err != nil {
			exception.ProcessError(err)
		}

		j.Execute()
	}
}
