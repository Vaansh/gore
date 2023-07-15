package domain

import (
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore"
	"github.com/Vaansh/gore/internal/gcloud"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/repository"
	"sync"
)

// Task service is a singleton

var (
	once                sync.Once
	taskServiceInstance *TaskService
)

type TaskService struct {
	Tasks       map[string]*Task
	StopChanMap map[string]chan struct{} // Map to store quit channels for each task
	db          *sql.DB
}

func NewTaskService() *TaskService {
	once.Do(func() {
		db, err := gcloud.InitDatabase()
		if err != nil {
			gcloud.LogFatal("unable to connect to repository")
		}

		taskServiceInstance = &TaskService{
			Tasks:       make(map[string]*Task),
			StopChanMap: make(map[string]chan struct{}),
			db:          db,
		}
	})
	return taskServiceInstance
}

func (s *TaskService) RunTask(publisherIds []string, publisherPlatforms []gore.Platform, subscriberId string, subscriberPlatform gore.Platform, metaData model.MetaData) error {
	taskID := subscriberPlatform.String() + subscriberId

	if _, ok := s.Tasks[taskID]; ok {
		return fmt.Errorf("task already running for the given subscriber")
	}

	repo, err := repository.NewPostgresUserRepository(s.db, subscriberId, subscriberPlatform)
	if err != nil {
		return err
	}

	stop := make(chan struct{})

	task := NewTask(publisherIds, publisherPlatforms, subscriberId, subscriberPlatform, metaData, repo)
	if task == nil {
		return fmt.Errorf("invalid task configuration received")
	}

	s.Tasks[taskID] = task
	s.StopChanMap[taskID] = stop

	go func() {
		task.Run(stop)
		delete(s.Tasks, taskID)
		delete(s.StopChanMap, taskID)
	}()

	return nil
}

func (s *TaskService) StopTask(subscriberID, subscriberPlatform string) error {
	taskID := subscriberPlatform + subscriberID

	if stop, ok := s.StopChanMap[taskID]; ok {
		close(stop)
		delete(s.Tasks, taskID)
		delete(s.StopChanMap, taskID)
		return nil
	}

	return fmt.Errorf("task not found for the given subscriber (maybe it finished running?)")
}
