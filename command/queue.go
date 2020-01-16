package command

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"ikdev/go-web/database/model"
	"ikdev/go-web/exception"
	"ikdev/go-web/job"
	"runtime"
	"sync"
	"time"
)

// Run jobs in Redis queue
func RunQueue(name string, sc *dig.Container) {
	var rc *redis.Client
	queue := fmt.Sprintf("queue:%s", name)
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

// Retry all failed jobs
// This method will put every failed job present in SQL db into original Redis queue
func RetryFailedQueue(container *dig.Container) {
	var failed []model.FailedJob

	err := container.Invoke(func(db *gorm.DB, r *redis.Client) {
		if err := db.Find(&failed).Error; err != nil {
			exception.ProcessError(err)
		}

		for _, f := range failed {
			queue := fmt.Sprintf("queue:%s", f.Queue)
			r.RPush(queue, f.Payload)

			removeFailedJob(&f, db)
		}
	})

	if err != nil {
		exception.ProcessError(err)
	}

}

// Remove failed job from SQL failed_job table
func removeFailedJob(job *model.FailedJob, db *gorm.DB) {
	if err := db.Unscoped().Delete(&job).Error; err != nil {
		exception.ProcessError(err)
	}
}
