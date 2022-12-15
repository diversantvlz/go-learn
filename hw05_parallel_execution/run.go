package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errCnt int32
	wg := &sync.WaitGroup{}
	wg.Add(n + 1)
	queue := make(chan Task)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for _, task := range tasks {
			if atomic.LoadInt32(&errCnt) >= int32(m) {
				break
			}
			queue <- task
		}
		close(queue)
	}(wg)

	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for task := range queue {
				if task() != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}(wg)
	}

	wg.Wait()

	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
