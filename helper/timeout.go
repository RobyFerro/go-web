package helper

import (
	"context"
	"fmt"
	"time"
)

// JobHandler defines the signature of the function that should be executed.
// by a job
type JobHandler = func([]byte) error

// JobResponseHandler defines the signature of a function that should be used
// to handle one of the possible responses of a job (success, failure, timeout).
type JobResponseHandler = func(error)

// JobResponse represents the response of a long running job.
type JobResponse struct {
	// The error that may be returned by running a job.
	Error error
	// The flag that indicates if the job was run, regardless of the result.
	// If false, this flag indicates a timeout error, which will be reflected by Error.
	Completed bool
}

// Job structure defines a job that should be run with the RunJob function.
type Job struct {
	// The handler being executed by the job.
	Handler JobHandler
	// Arguments of the job.
	Args []byte
	// The handler that could be will be processed if the job is run successfully.
	DidSucceed JobResponseHandler
	// The handler that could be will be processed if the job is run, but returns an error.
	DidFail JobResponseHandler
	// The handler that could be will be processed if the job does not return in less than a certain amount of time.
	DidTimeOut JobResponseHandler
	// Indicates the timeout for the execution of the job handler.
	// If not defined, this value will be set to 10 seconds.
	// If the value is set to something equal or lower to 0, it
	// will be set to 10 seconds.
	Timeout time.Duration
	// The number of times that a job should be executed in case of timeout (not error).
	// If set to 0 or less, the job will be executed just once.
	RetryCount int
}

// RunJob executes the handler of the given job once or more than once.
// If retryCount is 0 or less, the job handler will be run just once,
// otherwise it will be run 1 time plus the value of retryCount (worst case scenario).
func RunJob(job Job) {
	ctx := context.Background()
	result := JobResponse{}
	if job.Timeout <= 0 {
		job.Timeout = 10 * time.Second
	}
	if job.RetryCount < 0 {
		job.RetryCount = 0
	}
	// The job must be run at least once, hence the "<=" operator in this for-loop.
	for i := 0; i <= job.RetryCount; i++ {
		context.WithTimeout(context.Background(), job.Timeout)
		result = runJobWithContext(ctx, job, job.Timeout)
		if result.Completed {
			jobDidComplete(job, result)
			return
		}
		if i >= job.RetryCount {
			runJobResponseHandler(job.DidTimeOut, result.Error)
		}
	}
}

// Core functions.
// Runs a job.
func runJobWithContext(ctx context.Context, job Job, timeout time.Duration) JobResponse {
	// This channel is used to manage the execution timeout.
	channel := make(chan JobResponse)
	// The local context supports a timeout.
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	go func() {
		response := job.Handler(job.Args)
		channel <- JobResponse{Error: response, Completed: true}
	}()
	// This function will return something, i.e. either a response or a timeout error.
	select {
	case <-ctx.Done():
		return JobResponse{Error: fmt.Errorf("The job could not be run in the given time slot (%s)", timeout)}
	case response := <-channel:
		return response
	}
}

// Helper functions.
// Handles the completion of a job.
func jobDidComplete(job Job, response JobResponse) {
	if err := response.Error; err != nil {
		runJobResponseHandler(job.DidFail, err)
		return
	}
	runJobResponseHandler(job.DidSucceed, nil)
}

// Runs a generic JobResponseHandler
func runJobResponseHandler(handler JobResponseHandler, err error) {
	if handler != nil {
		handler(err)
	}
}
