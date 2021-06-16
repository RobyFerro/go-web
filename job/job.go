package job

import (
	"encoding/json"
	"fmt"
	"github.com/RobyFerro/go-web/app"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/service"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"reflect"
)

type Param struct {
	Name    string
	Payload string
	Type    string
}

// Job structure will be used to handle every jobs.
// Every method (except "Schedule" and "Execute") accept only one string parameter.
// This param will be decoded into a specific structure (manually defined into your job).
// Redis and MySQL is mandatory to use this feature.
// Redis will be used to handle every queue and MySQL to store failed jobs.
type Job struct {
	Name       string
	MethodName string
	Params     Param
	Queue      string
}

// Schedule specific job
func (j *Job) Schedule(queueName string, redis *redis.Client) {
	jobStr, err := json.Marshal(j)
	if err != nil {
		log.Error(err)
	}

	queue := fmt.Sprintf("queue:%s", queueName)
	redis.RPush(queue, jobStr)
}

// Execute specific job
func (j *Job) Execute() {
	r := reflect.ValueOf(Job{})
	method := r.MethodByName(j.MethodName)
	result := method.Call([]reflect.Value{reflect.ValueOf(j.Params.Payload)})

	// Put job on failed table
	// Beware of "too many connections.
	if result[1].Interface() != nil {
		err := result[1].Interface().(error)

		cf, _ := app.Configuration()
		r := service.ConnectDB(cf)

		jobStr, errMarshal := json.Marshal(j)
		if errMarshal != nil {
			log.Error(errMarshal)
		}

		j.PutOnFailed(j.Queue, string(jobStr), r, err)
		fmt.Printf("ERROR on job %s\n", j.Name)
		return
	}

	fmt.Printf("Job %s executed\n", j.Name)
}

// PutOnFailed inserts failed job in failed_job table
func (Job) PutOnFailed(queue, payload string, db *gorm.DB, err error) {
	failed := model.FailedJob{
		Payload:   payload,
		Queue:     queue,
		Exception: err.Error(),
	}

	if err := db.Save(&failed).Error; err != nil {
		log.Error(err)
	}
}
