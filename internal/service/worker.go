package service

import (
	"context"
	"errors"
	"sync"
)

const (
	DEFAULT_WORKER       = 10
	DEFAULT_QUEUE_BUFFER = 10
)

type workerService struct {
	queue   chan Command
	started bool
	cancel  context.CancelFunc
	mu      *sync.RWMutex
}

var workerServiceInstance *workerService

func InitWorkerService() *workerService {
	workerServiceInstance = &workerService{queue: make(chan Command, DEFAULT_QUEUE_BUFFER), started: false, mu: &sync.RWMutex{}}
	return workerServiceInstance
}

func GetWorkerService() (*workerService, error) {
	if workerServiceInstance == nil {
		return nil, errors.New("Backend: Create the Worker service first")
	}

	return workerServiceInstance, nil
}

func (worker *workerService) Start(ctx context.Context) error {
	defer close(worker.queue)

	errChan := make(chan error)

	ctx, cancelF := context.WithCancel(ctx)

	worker.mu.Lock()
	worker.cancel = cancelF
	worker.mu.Unlock()

	defer close(errChan)

	for i := 0; i < DEFAULT_WORKER; i++ {
		go func() {
			for {
				select {
				case command := <-worker.queue:
					{
						err := command.Run()
						if err != nil {
							errChan <- err
						}
					}
				}
			}
		}()
	}

	worker.mu.Lock()
	worker.started = true
	worker.mu.Unlock()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return errors.New("Context Canceled")
	}

}

func (worker *workerService) Stop() error {
	worker.mu.Lock()
	defer worker.mu.Unlock()
	if !worker.started || worker.cancel == nil {
		return errors.New("Worker Service not started")
	}
	worker.cancel()
	return nil
}

func (worker *workerService) AddToQueue(cmd Command) error {
	worker.mu.Lock()
	defer worker.mu.Unlock()
	if !worker.started || worker.cancel == nil {
		return errors.New("Worker Service not started")
	}

	worker.queue <- cmd
	return nil
}
