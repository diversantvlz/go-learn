package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type Queue struct {
	sync.RWMutex
	tasks []Task
	index int
}

func (q *Queue) get() Task {
	q.Lock()
	defer q.Unlock()

	if q.index == len(q.tasks) {
		return nil
	}

	task := q.tasks[q.index]
	q.index++

	return task
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	queue := Queue{tasks: tasks}
	var errCnt int32

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				if atomic.LoadInt32(&errCnt) >= int32(m) {
					break
				}

				task := queue.get()
				if nil == task {
					break
				}

				if task() != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()

		wg.Add(1)
	}

	wg.Wait()

	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
