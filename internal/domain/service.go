package domain

import (
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/model"
	"sync"
)

type TaskService struct {
	Tasks       map[string]*Task
	StopChanMap map[string]chan struct{} // Map to store quit channels for each task
	mutex       sync.Mutex
	db          *sql.DB
}

func NewTaskService() (*TaskService, error) {
	db, err := database.InitDb()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &TaskService{
		Tasks:       make(map[string]*Task),
		StopChanMap: make(map[string]chan struct{}),
		db:          db,
	}, nil
}

func (s *TaskService) RunTask(publisherIds []string, publisherPlatforms []go_pubsub.Name, subscriberId string, subscriberPlatform go_pubsub.Name, metaData model.MetaData) error {
	taskID := subscriberPlatform.String() + subscriberId

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.Tasks[taskID]; ok {
		return fmt.Errorf("task already running for the given subscriber")
	}

	stop := make(chan struct{})
	task := NewTask(publisherIds, publisherPlatforms, subscriberId, subscriberPlatform, metaData, *database.NewUserRepository(s.db, subscriberId, subscriberPlatform))
	if task == nil {
		return fmt.Errorf("invalid task configuration received")
	}

	s.Tasks[taskID] = task
	s.StopChanMap[taskID] = stop

	go func() {
		task.Run(stop)
		s.mutex.Lock()
		defer s.mutex.Unlock()
		delete(s.Tasks, taskID)
		delete(s.StopChanMap, taskID)
	}()

	return nil
}

func (s *TaskService) StopTask(subscriberID, subscriberPlatform string) error {
	taskID := subscriberPlatform + subscriberID

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if stop, ok := s.StopChanMap[taskID]; ok {
		close(stop)
		delete(s.Tasks, taskID)
		delete(s.StopChanMap, taskID)
		return nil
	}

	return fmt.Errorf("task not found for the given subscriber (maybe it finished running?)")
}
