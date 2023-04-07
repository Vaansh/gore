package instagram

import (
	"fmt"
	"time"
)

type InstagramSubscriber struct {
	InstagramID string
}

func NewInstagramConsumer(InstagramID string) *InstagramSubscriber {
	return &InstagramSubscriber{InstagramID: InstagramID}
}

func (p *InstagramSubscriber) SubscribeTo(c <-chan string) {
	fmt.Println("Consuming...")
	for link := range c {
		fmt.Printf("New video uploaded: %s\n", link)
		time.Sleep(4 * time.Second)
	}
}

func (p *InstagramSubscriber) GetSubscriberID() string {
	return p.InstagramID
}
