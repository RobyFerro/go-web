package job

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"ikdev/go-web/config"
	"ikdev/go-web/database"
	"ikdev/go-web/exception"
	"ikdev/go-web/helper"
	"reflect"
)

type Param struct {
	Name    string
	Payload string
	Type    string
}

// This structure will be used to handle every jobs.
// Every method (except "Schedule" and "Execute") accept only one string parameter.
// This param will be decoded into a specific structure (manually defined into your job)
type Job struct {
	Name       string
	MethodName string
	Params     Param
}

// Schedule specific job
func (j *Job) Schedule(queueName string, redis *redis.Client) {
	jobStr, err := json.Marshal(j)
	if err != nil {
		exception.ProcessError(err)
	}

	queue := fmt.Sprintf("queue:%s", queueName)
	redis.RPush(queue, jobStr)
}

// Execute specific job
func (j *Job) Execute() {
	r := reflect.ValueOf(Job{})
	method := r.MethodByName(j.MethodName)
	result := method.Call([]reflect.Value{reflect.ValueOf(j.Params.Payload)})
	err := result[1].Interface().(error)

	if err != nil {
		r := database.ConnectRedis(config.Configuration())
		j.PutOnFailed(j.Params.Payload, r)
		helper.Log(err.Error())
	}
}

// Insert failed job in failed queue
func (Job) PutOnFailed(payload string, r *redis.Client) {
	r.RPush("queue:failed", payload)
}
