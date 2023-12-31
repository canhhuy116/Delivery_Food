package asyncjob

import (
	"golang.org/x/net/context"
	"time"
)

// Job requirement:
// 1. Job can do something in background (handler)
// 2. Job can retry if failed
//   2.1. Job can retry with custom duration
//   2.2. Job can retry with default duration
// 3. Should be stateful
// 4. We should have job manager to manage all jobs

type Job interface {
	// Execute init job
	Execute(ctx context.Context) error
	// Retry retry job
	Retry(ctx context.Context) error
	// State get job state
	State() JobState
	// SetRetryDurations set retry duration
	SetRetryDurations(times []time.Duration)
}

const (
	// set default max time out ( 10s)
	defaultMaxTimeout = time.Second * 10
	// set default max retry time out ( 3 times)
	defaultMaxRetryCount = 3
)

var (
	// set default RetryTime 1s, 5s, 10s
	defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}
)

type JobHandler func(ctx context.Context) error

// JobState define type JobState
type JobState int

const (
	// StateInit iota =  0 (default) and if create new line (the variable in new line has value  = iota ++ )
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

type jobHandler func(ctx context.Context) error

// define jobConfig
type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

// js has type int, but I want to convert it to string, so I use String() with []string have type string
func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type job struct {
	config     jobConfig  // include max timeout and retry time (array)
	handler    JobHandler // function
	state      JobState   // state of JobState
	retryIndex int        // index of retry time
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	return &job{
		config:     jobConfig{MaxTimeout: defaultMaxTimeout, Retries: defaultRetryTime},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
}
func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning
	var err error
	err = j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}
	j.state = StateCompleted
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])
	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}
	j.config.Retries = times
}
