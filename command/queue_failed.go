package command

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"ikdev/go-web/app/kernel"
	"ikdev/go-web/database/model"
	"ikdev/go-web/exception"
)

type QueueFailed struct {
	Signature   string
	Description string
}

func (c *QueueFailed) Register() {
	c.Signature = "queue:failed"
	c.Description = "Move failed jobs into a queue"
}

// Run jobs in Redis queue
func (c *QueueFailed) Run(kernel *kernel.HttpKernel, args string, console map[string]interface{}) {
	var failed []model.FailedJob

	err := kernel.Container.Invoke(func(db *gorm.DB, r *redis.Client) {
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
