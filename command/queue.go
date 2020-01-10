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
