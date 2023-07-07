package internal

import (
	"github.com/Vaansh/gore/internal/publisher"
	platform2 "github.com/Vaansh/gore/internal/subscriber"
)

type PlatformName string

const (
	PLATFORM PlatformName = "PF"
)

type Task struct {
	ID         string
	Producers  []publisher.YoutubePublisher
	Subscriber platform2.Subscriber
}

func NewTask(Id string, producers []publisher.YoutubePublisher, subscriber platform2.Subscriber) *Task {
	return &Task{
		ID:         Id,
		Producers:  producers,
		Subscriber: subscriber,
	}
}

func (t *Task) Run() {
	c := make(chan string)
	for _, p := range t.Producers {
		go p.PublishTo(c)
	}
	go t.Subscriber.SubscribeTo(c, PLATFORM)
}
