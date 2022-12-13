package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type Queue struct {
	sync.RWMutex
	tasks []Task
	index int
}

func (q *Queue) get() Task {
	q.RLock()
	defer q.RUnlock()

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
	wg.Add(n)
	queue := Queue{tasks: tasks}

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				if m == 0 {
					break
				}

				task := queue.get()
				if nil == task {
					break
				}

				if task() != nil && m > 0 {
					m--
				}
			}
		}()
	}

	wg.Wait()

	if m == 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
