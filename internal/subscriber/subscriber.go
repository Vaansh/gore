package subscriber

import "github.com/Vaansh/gore/internal"

type Subscriber interface {
	SubscribeTo(c <-chan string, platform internal.PlatformName)
	GetSubscriberID() string
}
