package console

import (
	"fmt"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"log"
)

type QueueFailed struct {
	Signature   string
	Description string
	Args        string
}

func (c *QueueFailed) Register() {
	c.Signature = "queue:failed"
	c.Description = "Move failed jobs into a queue"
}

// Run jobs in Redis queue
func (c *QueueFailed) Run(db *gorm.DB, r *redis.Client) {
	var failed []model.FailedJob

	if err := db.Find(&failed).Error; err != nil {
		log.Fatal(err)
	}

	for _, f := range failed {
		queue := fmt.Sprintf("queue:%s", f.Queue)
		r.RPush(queue, f.Payload)

		removeFailedJob(&f, db)
	}
}

// Remove failed job from SQL failed_job table
func removeFailedJob(job *model.FailedJob, db *gorm.DB) {
	if err := db.Unscoped().Delete(&job).Error; err != nil {
		log.Fatal(err)
	}
}
