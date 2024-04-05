package main

import "github.com/tech-pratheesh/learning-with-golang/worker"

func main() {
	w := worker.NewWorker(
		worker.WithJobs([]string{"job #1"}),
		worker.WithRetryFailure(func() string {
			return "failure updated"
		}),
		worker.WithFlag(),
		worker.WithFlagDirect,
	)

	w.Run()
}
