package service

import (
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/domain"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/platform"
	"sync"
)

type TaskService struct {
	Tasks       map[string]*domain.Task
	StopChanMap map[string]chan struct{} // Map to store quit channels for each task
	mutex       sync.Mutex
	db          *sql.DB
}

func NewTaskService() (*TaskService, error) {
	db, err := database.ConnectDb()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &TaskService{
		Tasks:       make(map[string]*domain.Task),
		StopChanMap: make(map[string]chan struct{}),
		db:          db,
	}, nil
}

func (s *TaskService) RunTask(channels, subscriberID, igPostTags string) error {
	taskID := platform.INSTAGRAM.String() + subscriberID

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.Tasks[taskID]; ok {
		return fmt.Errorf("task already running for the given subscriber")
	}

	stop := make(chan struct{})
	plats := []platform.Name{platform.YOUTUBE}
	chans := []string{channels}
	task := domain.NewTask(chans, plats, subscriberID, platform.INSTAGRAM, model.MetaData{IgPostTags: igPostTags},
		*database.NewUserRepository(s.db, subscriberID, platform.INSTAGRAM))
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

	return fmt.Errorf("task not found for the given subscriber, maybe it finished running")
}
