package internal

import (
	platform2 "github.com/Vaansh/gore/internal/platform"
)

type PlatformName string

const (
	PLATFORM PlatformName = "PF"
)

type Task struct {
	ID         string
	Producers  []platform2.Publisher
	Subscriber platform2.Subscriber
}

func NewTask(Id string, producers []platform2.Publisher, subscriber platform2.Subscriber) *Task {
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
	go t.Subscriber.SubscribeTo(c)
}
