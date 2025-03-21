package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	// n - количество горутин
	// m - максимальное количество ошибок

	if m <= 0 || n <= 0 {
		return ErrErrorsLimitExceeded
	}
	var wg sync.WaitGroup

	// Создаем канал для заданий
	tasksChannel := make(chan Task)

	// Считаем количество ошибок
	var errorsCount int32

	// Обрабатываем задания
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Worker(tasksChannel, &wg, &errorsCount)
	}

	// Пишем задания в канал
	for index, task := range tasks {
		if n+m == index && int(atomic.LoadInt32(&errorsCount)) >= m {
			break
		}
		tasksChannel <- task
	}

	// Закроем канал, чтобы горутины-воркеры могли завершиться
	close(tasksChannel)

	wg.Wait()

	if int(atomic.LoadInt32(&errorsCount)) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func Worker(tasksChannel <-chan Task, wg *sync.WaitGroup, errorsCount *int32) {
	defer wg.Done()
	for task := range tasksChannel {
		err := task()
		if err != nil {
			atomic.AddInt32(errorsCount, 1)
		}
	}
}
