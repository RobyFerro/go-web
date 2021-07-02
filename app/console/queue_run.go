package console

import (
	"encoding/json"
	"fmt"
	foundation "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/job"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/gommon/log"
	"runtime"
	"sync"
	"time"
)

type QueueRun struct {
	Signature   string
	Description string
	Args        string
}

func (c *QueueRun) Register() {
	c.Signature = "queue:run <name>"
	c.Description = "Run a specific queue"
}

func (c *QueueRun) Run() {
	queue := fmt.Sprintf("queue:%s", c.Args)
	cpus := runtime.NumCPU()

	var wg sync.WaitGroup
	wg.Add(cpus)
	_ = foundation.RetrieveServiceContainer().Invoke(func(redis *redis.Client) {
		for i := 0; i < cpus; i++ {
			go worker(queue, redis, &wg)
		}

		wg.Wait()
	})
}

// Execute worker
// This method will schedule one worker for each CPU
func worker(queue string, rc *redis.Client, wg *sync.WaitGroup) {
	for true {
		var j job.Job
		tasks, err := rc.BLPop(5*time.Second, queue).Result()

		if err != nil && err.Error() != "redis: nil" {
			log.Error(err)
			break
		}

		if len(tasks) != 2 {
			continue
		}

		if err := json.Unmarshal([]byte(tasks[1]), &j); err != nil {
			log.Error(err)
		}

		j.Execute()
	}

	wg.Done()
}
