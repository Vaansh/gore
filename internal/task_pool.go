package internal

import (
	"fmt"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/platform"
	"github.com/Vaansh/gore/internal/publisher"
	"github.com/Vaansh/gore/internal/subscriber"
	"sync"
)

type TaskPool struct {
	Tasks map[string]*Task
}

func NewTaskPool() *TaskPool {
	return &TaskPool{
		Tasks: make(map[string]*Task),
	}
}

func (tm *TaskPool) RunAll() {
	var wg sync.WaitGroup
	for _, task := range tm.Tasks {
		wg.Add(1)
		go func(t *Task) {
			defer wg.Done()
			t.Run()
		}(task)
	}
	wg.Wait()
}

func (tm *TaskPool) AddTask(publisherIds []string, sources []platform.Name, subscriberId string, destination platform.Name, metadata model.MetaData, repository database.UserRepository) error {
	taskID := string(destination) + subscriberId
	if _, exists := tm.Tasks[taskID]; exists {
		return fmt.Errorf("worker with id %s already exists", taskID)
	}

	if len(publisherIds) != len(sources) {
		return fmt.Errorf("received %d publisherIds and %d platforms", len(publisherIds), len(sources))
	}

	prods := make([]publisher.Publisher, len(publisherIds))
	for i, id := range publisherIds {
		switch sources[i] {
		case platform.YOUTUBE:
			prods[i] = publisher.NewYoutubePublisher(id)
		default:
			return fmt.Errorf("supported.go not found %s for %s", sources[i], id)
		}
	}

	if destination == platform.INSTAGRAM {
		consumer := subscriber.NewInstagramSubscriber(subscriberId, metadata, repository)
		task := NewTask(taskID, prods, consumer)
		tm.Tasks[task.Id] = task
	}

	return nil
}

func (tm *TaskPool) EditTask(taskID string, publishers []publisher.Publisher, subscriber subscriber.Subscriber) error {
	task, ok := tm.Tasks[taskID]
	if !ok {
		return fmt.Errorf("task %s not found", taskID)
	}
	task.Publishers = publishers
	task.Subscriber = subscriber
	return nil
}

func (tm *TaskPool) DeleteTask(taskID string) {
	delete(tm.Tasks, taskID)
}
