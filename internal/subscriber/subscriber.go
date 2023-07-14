package subscriber

import (
	"github.com/Vaansh/gore/internal/model"
)

type Subscriber interface {
	SubscribeTo(c <-chan model.Post)
	GetSubscriberId() string
}
