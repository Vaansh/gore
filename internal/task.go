package internal

import (
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/publisher"
	"github.com/Vaansh/gore/internal/subscriber"
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
	for _, p := range t.Publishers {
		go p.PublishTo(c)
	}
	go t.Subscriber.SubscribeTo(c)
}
