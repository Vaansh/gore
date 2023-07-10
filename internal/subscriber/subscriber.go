package subscriber

import "github.com/Vaansh/gore/internal/platform"

type Subscriber interface {
	SubscribeTo(c <-chan string, platform platform.Name)
	GetSubscriberID() string
}
