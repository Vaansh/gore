package domain

import (
	"fmt"
	"github.com/Vaansh/gore"
	"github.com/Vaansh/gore/internal/gcloud"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/publisher"
	"github.com/Vaansh/gore/internal/repository"
	"github.com/Vaansh/gore/internal/subscriber"
	"reflect"
	"sync"
)

type Task struct {
	Id         string
	Publishers []publisher.Publisher
	Subscriber subscriber.Subscriber
	Quit       chan struct{}
}

func NewTask(publisherIds []string, sources []gore.Platform, subscriberId string, destination gore.Platform, metadata model.MetaData, repository repository.UserRepository) *Task {
	id := string(destination) + subscriberId
	if len(publisherIds) != len(sources) {
		return nil
	}

	publishers := make([]publisher.Publisher, len(publisherIds))
	for i, id := range publisherIds {
		switch sources[i] {
		case gore.YOUTUBE:
			publishers[i] = publisher.NewYoutubePublisher(id)
		default:
			return nil
		}
	}

	var consumer subscriber.Subscriber = nil
	if destination == gore.INSTAGRAM {
		consumer = subscriber.NewInstagramSubscriber(subscriberId, metadata, repository)
	}

	if consumer == nil {
		return nil
	}

	return &Task{
		Id:         id,
		Publishers: publishers,
		Subscriber: consumer,
		Quit:       make(chan struct{}),
	}
}

func (t *Task) Run(stop chan struct{}) {
	c := make(chan model.Post)

	var wg sync.WaitGroup
	wg.Add(len(t.Publishers))

	for _, p := range t.Publishers {
		gcloud.LogInfo(fmt.Sprintf("%s (id: %s) has started publishing for task %s", reflect.TypeOf(p), p.GetPublisherId(), t.Id))
		go func(publisher publisher.Publisher) {
			defer wg.Done()
			publisher.PublishTo(c, t.Quit)
		}(p)
	}

	// Stop the goroutine execution
	go func() {
		select {
		case <-stop:
			gcloud.LogInfo(fmt.Sprintf("Task %s was force quit", t.Id))
			close(t.Quit)
			return
		}
	}()

	go func() {
		gcloud.LogInfo(fmt.Sprintf("Subscriber has subscribed to task %s", t.Id))
		t.Subscriber.SubscribeTo(c)
		close(t.Quit)
	}()

	wg.Wait()
}
