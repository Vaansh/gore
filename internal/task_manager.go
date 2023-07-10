package internal

import (
	"fmt"
	"github.com/Vaansh/gore/internal/publisher"
	"github.com/Vaansh/gore/internal/subscriber"
	"sync"
)

type TaskManager struct {
	Tasks map[string]*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[string]*Task),
	}
}

func (tm *TaskManager) RunAll() {
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

func (tm *TaskManager) AddTask(producerIDs []string, sources []PlatformName, consumerID string, destination PlatformName) error {
	taskID := string(destination) + consumerID
	if _, exists := tm.Tasks[taskID]; exists {
		return fmt.Errorf("task with ID %s already exists", taskID)
	}

	if len(producerIDs) != len(sources) {
		return fmt.Errorf("received %d producerIds and %d platforms", len(producerIDs), len(sources))
	}

	prods := make([]publisher.Publisher, len(producerIDs))
	for i, id := range producerIDs {
		switch sources[i] {
		case YOUTUBE:
			prods[i] = *publisher.NewYoutubePublisher(id)
		default:
			return fmt.Errorf("platform not found %s for %s", sources[i], id)
		}
	}

	consumer := subscriber.NewInstagramSubscriber(consumerID)
	task := NewTask(taskID, prods, consumer)
	tm.Tasks[task.ID] = task
	return nil
}

func (tm *TaskManager) EditTask(taskID string, producers []publisher.Publisher, consumer subscriber.Subscriber) error {
	task, ok := tm.Tasks[taskID]
	if !ok {
		return fmt.Errorf("task %s not found", taskID)
	}
	task.Producers = producers
	task.Subscriber = consumer
	return nil
}

func (tm *TaskManager) DeleteTask(taskID string) {
	delete(tm.Tasks, taskID)
}
