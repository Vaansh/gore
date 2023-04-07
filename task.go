package main

type Task struct {
	ID        string
	Producers []Publisher
	Consumer  Subscriber
}

func NewTask(Id string, producers []Publisher, consumer Subscriber) *Task {
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
