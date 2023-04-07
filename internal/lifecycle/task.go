package lifecycle

import "github.com/Vaansh/gore/internal/platform"

type Task struct {
	ID         string
	Producers  []platform.Publisher
	Subscriber platform.Subscriber
}

func NewTask(Id string, producers []platform.Publisher, subscriber platform.Subscriber) *Task {
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
