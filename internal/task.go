package internal

import (
	"pubsub/internal/consumers"
	"pubsub/internal/producers"
)

type Task struct {
	ID        string
	Producers []producers.Producer
	Consumer  consumers.Consumer
}

func NewTask(Id string, producers []producers.Producer, consumer consumers.Consumer) *Task {
	return &Task{
		ID:        Id,
		Producers: producers,
		Consumer:  consumer,
	}
}

func (t *Task) Run() {
	c := make(chan string)
	for _, p := range t.Producers {
		go p.ProduceOn(c)
	}
	go t.Consumer.ConsumeOn(c)
}
