package job

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"ikdev/go-web/exception"
	"reflect"
	"strconv"
)

type Param struct {
	Name    string
	Payload []byte
	Type    string
}

type Job struct {
	Name       string
	MethodName string
	Params     []Param
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

	var in []reflect.Value
	for _, params := range j.Params {
		// Decode payload byte to Type
		strPayload := string(params.Payload)

		switch params.Type {
		// To be implemented with other types
		case "string":
			in = append(in, reflect.ValueOf(strPayload))
			break
		case "int":
			data, _ := strconv.Atoi(strPayload)
			in = append(in, reflect.ValueOf(data))
			break
		}
		// End decode
	}

	method.Call(in)
}
