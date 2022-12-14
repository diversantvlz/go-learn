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
	wg := &sync.WaitGroup{}
	queue := make(chan Task, len(tasks))
	for _, task := range tasks {
		queue <- task
	}

	var errCnt int32
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
		out:
			for {
				if atomic.LoadInt32(&errCnt) >= int32(m) {
					break
				}

				select {
				case task := <-queue:
					if task() != nil {
						atomic.AddInt32(&errCnt, 1)
					}
				default:
					break out
				}
			}
		}(wg)
	}

	wg.Wait()
	close(queue)

	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
