package platform

type Publisher interface {
	PublishTo(c chan<- string)
	GetPublisherID() string
}

type Subscriber interface {
	SubscribeTo(c <-chan string)
	GetSubscriberID() string
}
