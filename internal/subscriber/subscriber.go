package subscriber

type Subscriber interface {
	SubscribeTo(c <-chan string)
	GetSubscriberID() string
}
