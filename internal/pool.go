package internal

import (
	"fmt"
	"github.com/Vaansh/gore/internal/platform"
	"github.com/Vaansh/gore/internal/publisher"
	"github.com/Vaansh/gore/internal/subscriber"
	"sync"
)

type Pool struct {
	Tasks map[string]*Task
}

func NewPool() *Pool {
	return &Pool{
		Tasks: make(map[string]*Task),
	}
}

func (tm *Pool) RunAll() {
	var wg sync.WaitGroup
	for _, task := range tm.Tasks {
		wg.Add(1)
		go func(t *Task) {
			defer wg.Done()
			t.Run()
			select {}
		}(task)
	}
	wg.Wait()
}

func (tm *Pool) AddTask(producerIDs []string, sources []platform.Name, consumerID string, destination platform.Name) error {
	taskID := string(destination) + consumerID
	if _, exists := tm.Tasks[taskID]; exists {
		return fmt.Errorf("worker with ID %s already exists", taskID)
	}

	if len(producerIDs) != len(sources) {
		return fmt.Errorf("received %d producerIds and %d platforms", len(producerIDs), len(sources))
	}

	prods := make([]publisher.Publisher, len(producerIDs))
	for i, id := range producerIDs {
		switch sources[i] {
		case platform.YOUTUBE:
			prods[i] = *publisher.NewYoutubePublisher(id)
		default:
			return fmt.Errorf("supported.go not found %s for %s", sources[i], id)
		}
	}

	if destination == platform.INSTAGRAM {
		consumer := subscriber.NewInstagramSubscriber(consumerID)
		task := NewTask(taskID, prods, consumer)
		tm.Tasks[task.ID] = task
	}

	return nil
}

func (tm *Pool) EditTask(taskID string, producers []publisher.Publisher, consumer subscriber.Subscriber) error {
	task, ok := tm.Tasks[taskID]
	if !ok {
		return fmt.Errorf("worker %s not found", taskID)
	}
	task.Producers = producers
	task.Subscriber = consumer
	return nil
}

func (tm *Pool) DeleteTask(taskID string) {
	delete(tm.Tasks, taskID)
}
