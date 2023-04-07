package main

type Task struct {
	ID        string
	Producers []Producer
	Consumer  Consumer
}

func NewTask(Id string, producers []Producer, consumer Consumer) *Task {
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
