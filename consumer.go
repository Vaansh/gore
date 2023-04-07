package main

import (
	"fmt"
	"time"
)

type Consumer interface {
	ConsumeOn(c <-chan string)
	ConsumerID() string
}

type InstagramConsumer struct {
	InstagramID string
}

func NewInstagramConsumer(InstagramID string) *InstagramConsumer {
	return &InstagramConsumer{InstagramID: InstagramID}
}

func (p *InstagramConsumer) ConsumeOn(c <-chan string) {
	fmt.Println("Consuming...")
	for link := range c {
		fmt.Printf("New video uploaded: %s\n", link)
		time.Sleep(4 * time.Second)
	}
}

func (p *InstagramConsumer) ConsumerID() string {
	return p.InstagramID
}
