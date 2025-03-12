package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// n - количество горутин
	// m - максимальное количество ошибок

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	tasksChan := make(chan Task)
	errorsChan := make(chan struct{}, m)
	quitChan := make(chan bool, n)

	defer close(tasksChan)
	defer close(errorsChan)
	defer close(quitChan)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go Worker(&wg, tasksChan, errorsChan, quitChan, m)
	}

	var err error
	for index, task := range tasks {
		if n+m == index {
			if len(errorsChan) == m {
				err = ErrErrorsLimitExceeded
				break
			}
		}
		tasksChan <- task
	}

	for i := 0; i < n; i++ {
		quitChan <- true
	}

	wg.Wait()

	if err != nil {
		return err
	}

	return nil
}

func Worker(wg *sync.WaitGroup, tasksChan <-chan Task, errorsChan chan<- struct{}, quitChan <-chan bool, m int) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-tasksChan:
			if !ok {
				return
			}
			err := task()
			if err != nil {
				if len(errorsChan) == m {
					return
				}
				errorsChan <- struct{}{}
			}
		case <-quitChan:
			return
		}
	}
}
