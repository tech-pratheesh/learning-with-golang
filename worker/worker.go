package worker

import "fmt"

type Worker struct {
	jobs                []string
	retryCount          int
	retryFailureHandler func() string
	Flag                bool
}

type Option struct {
	Jobs       []string
	RetryCount int
}

type Options func(*Worker)

func WithRetryCount(count int) Options {
	return func(w *Worker) {
		w.retryCount = count
	}
}

func WithRetryFailure(fn func() string) Options {
	return func(w *Worker) {
		w.retryFailureHandler = fn
	}
}

func WithJobs(jobs []string) Options {
	return func(w *Worker) {
		w.jobs = jobs
	}
}

func WithFlag() Options {
	return func(w *Worker) {
		w.Flag = true
	}
}

func WithFlagDirect(w *Worker) {
	w.Flag = true
}
func NewWorker(opt ...Options) *Worker {

	w := &Worker{
		retryCount: 10,
		retryFailureHandler: func() string {
			return "failed"
		},
	}

	// option
	for _, op := range opt {
		op(w)
	}

	return w
}

func (w *Worker) Run() {
	fmt.Println("worker running : ", w.jobs, w.retryCount, w.retryFailureHandler())
}
