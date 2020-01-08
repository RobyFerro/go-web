package job

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"ikdev/go-web/exception"
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
	method.Call([]reflect.Value{reflect.ValueOf(j.Params.Payload)})
}
