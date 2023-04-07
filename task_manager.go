package main

import (
	"fmt"
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

func (tm *TaskManager) AddTask(producerIDs []string, sources []PlatformName, consumerID string, destination PlatformName) error {
	taskID := string(destination) + consumerID
	if _, exists := tm.Tasks[taskID]; exists {
		return fmt.Errorf("task with ID %s already exists", taskID)
	}

	if len(producerIDs) != len(sources) {
		return fmt.Errorf("received %d producerIds and %d platforms", len(producerIDs), len(sources))
	}

	producers := make([]Producer, len(producerIDs))
	for i, id := range producerIDs {
		switch sources[i] {
		case PLATFORM:
			producers[i] = NewYoutubeProducer(id)
		default:
			return fmt.Errorf("platform not found %s for %s", sources[i], id)
		}
	}

	consumer := NewInstagramConsumer(consumerID)
	task := NewTask(taskID, producers, consumer)
	tm.Tasks[task.ID] = task
	return nil
}

func (tm *TaskManager) EditTask(taskID string, producers []Producer, consumer Consumer) error {
	task, ok := tm.Tasks[taskID]
	if !ok {
		return fmt.Errorf("task %s not found", taskID)
	}
	task.Producers = producers
	task.Consumer = consumer
	return nil
}

func (tm *TaskManager) DeleteTask(taskID string) {
	delete(tm.Tasks, taskID)
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
