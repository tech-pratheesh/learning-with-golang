# learning-with-golang
A code learning repo



## Struct with pattern
```go
w := worker.NewWorker(
		worker.WithJobs([]string{"job #1"}),
		worker.WithRetryFailure(func() string {
			return "failure updated"
		}),
		worker.WithFlag(),
		worker.WithFlagDirect,
	)

```


## prevent unkeyed literal
```
type Config struct {
	_      struct{}
}
```