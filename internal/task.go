package internal

import (
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/publisher"
	"github.com/Vaansh/gore/internal/subscriber"
	"sync"
)

type Task struct {
	Id         string
	Publishers []publisher.Publisher
	Subscriber subscriber.Subscriber
}

func NewTask(id string, publishers []publisher.Publisher, subscriber subscriber.Subscriber) *Task {
	return &Task{
		Id:         id,
		Publishers: publishers,
		Subscriber: subscriber,
	}
}

func (t *Task) Run() {
	c := make(chan model.Post)
	quit := make(chan struct{})
	
	var wg sync.WaitGroup
	wg.Add(len(t.Publishers))

	for _, p := range t.Publishers {
		go func(publisher publisher.Publisher) {
			defer wg.Done()
			publisher.PublishTo(c, quit)
		}(p)
	}

	go func() {
		t.Subscriber.SubscribeTo(c)
		close(quit)
	}()

	wg.Wait()
}
