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

	var in []reflect.Value
	// Decode payload byte to Type
	for _, params := range j.Params {
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
		case "int64":
			data, _ := strconv.ParseInt(strPayload, 10, 64)
			in = append(in, reflect.ValueOf(data))
			break
		case "bool":
			data, _ := strconv.ParseBool(strPayload)
			in = append(in, reflect.ValueOf(data))
			break
		case "float32":
			data, _ := strconv.ParseFloat(strPayload, 32)
			in = append(in, reflect.ValueOf(data))
			break
		case "float64":
			data, _ := strconv.ParseFloat(strPayload, 64)
			in = append(in, reflect.ValueOf(data))
			break
		case "uint32":
			data, _ := strconv.ParseUint(strPayload, 10, 32)
			in = append(in, reflect.ValueOf(data))
			break
		case "uint64":
			data, _ := strconv.ParseUint(strPayload, 10, 64)
			in = append(in, reflect.ValueOf(data))
			break
		}
	}
	// End decode

	method.Call(in)
}
