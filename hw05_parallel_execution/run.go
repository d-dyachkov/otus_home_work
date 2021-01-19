package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrInvalidWorkersNumber = errors.New("invalid number of workers")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	if err := checkValidParams(n, m); err != nil {
		return err
	}

	taskCh := make(chan Task, len(tasks))

	wg := sync.WaitGroup{}
	var errCount, maxErrorsCount int32 = 0, int32(m)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stopped := false
			for task := range taskCh {
				if err := task(); err != nil {
					stopped = atomic.AddInt32(&errCount, 1) >= maxErrorsCount
				} else {
					stopped = atomic.LoadInt32(&errCount) >= maxErrorsCount
				}
				if stopped {
					break
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCount) >= maxErrorsCount {
			break
		}
		taskCh <- task
	}
	close(taskCh)
	wg.Wait()

	if errCount >= maxErrorsCount {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func checkValidParams(n int, m int) error {
	switch {
	case m <= 0:
		return ErrErrorsLimitExceeded
	case n <= 0:
		return ErrInvalidWorkersNumber
	default:
		return nil
	}
}
