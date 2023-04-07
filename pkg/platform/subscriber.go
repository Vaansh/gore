package platform

type Subscriber interface {
	SubscribeTo(c <-chan string)
	GetSubscriberID() string
}
