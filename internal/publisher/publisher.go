package publisher

import (
	"github.com/Vaansh/gore/internal/model"
)

type Publisher interface {
	PublishTo(c chan<- model.Post)
	GetPublisherId() string
}
