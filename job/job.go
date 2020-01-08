package job

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"ikdev/go-web/exception"
	"reflect"
)

type Job struct {
	Name       string
	MethodName string
	Params     []reflect.Value
}

func (j *Job) Schedule(queueName string, redis *redis.Client) {
	jobStr, err := json.Marshal(j)
	if err != nil {
		exception.ProcessError(err)
	}

	queue := fmt.Sprintf("queue:%s", queueName)
	redis.RPush(queue, jobStr)
}

func (j *Job) Execute() {
	r := reflect.ValueOf(Job{})
	method := r.MethodByName(j.MethodName)
	method.Call(j.Params)
}
