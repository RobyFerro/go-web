package command

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"go.uber.org/dig"
	"ikdev/go-web/exception"
	"ikdev/go-web/job"
	"runtime"
	"sync"
	"time"
)

type QueueRun struct {
	Signature   string
	Description string
}

func (c *QueueRun) Register() {
	c.Signature = "queue:run <name>"
	c.Description = "Run a specific queue"
}

func (c *QueueRun) Run(sc *dig.Container, args string) {
	var rc *redis.Client
	queue := fmt.Sprintf("queue:%s", args)
	cpus := runtime.NumCPU()

	err := sc.Invoke(func(r *redis.Client) {
		rc = r
	})

	if err != nil {
		exception.ProcessError(err)
	}

	var wg sync.WaitGroup
	wg.Add(cpus)
	for i := 0; i < cpus; i++ {
		go worker(queue, rc, &wg)
	}

	wg.Wait()
}

// Execute worker
// This method will schedule one worker for each CPU
func worker(queue string, rc *redis.Client, wg *sync.WaitGroup) {
	for true {
		var j job.Job
		tasks, err := rc.BLPop(5*time.Second, queue).Result()

		if err != nil && err.Error() != "redis: nil" {
			exception.ProcessError(err)
			break
		}

		if len(tasks) != 2 {
			continue
		}

		if err := json.Unmarshal([]byte(tasks[1]), &j); err != nil {
			exception.ProcessError(err)
		}

		j.Execute()
	}

	wg.Done()
}
