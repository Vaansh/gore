package lifecycle

import (
	"pubsub/internal/publisher"
	"pubsub/internal/subscriber"
)

type Task struct {
	ID        string
	Producers []publisher.Publisher
	Consumer  subscriber.Subscriber
}

func NewTask(Id string, producers []publisher.Publisher, consumer subscriber.Subscriber) *Task {
	return &Task{
		ID:        Id,
		Producers: producers,
		Consumer:  consumer,
	}
}

func (t *Task) Run() {
	c := make(chan string)
	for _, p := range t.Producers {
		go p.PublishTo(c)
	}
	go t.Consumer.SubscribeTo(c)
}
