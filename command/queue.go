package command

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"go.uber.org/dig"
	"ikdev/go-web/exception"
	"ikdev/go-web/job"
	"time"
)

// Run jobs in Redis queue
func RunQueue(name string, sc *dig.Container) {
	var rc *redis.Client
	queue := fmt.Sprintf("queue:%s", name)

	err := sc.Invoke(func(r *redis.Client) {
		rc = r
	})

	if err != nil {
		exception.ProcessError(err)
	}

	for true {
		tasks, _ := rc.BLPop(5*time.Second, queue).Result()
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
