package publisher

type Publisher interface {
	PublishTo(c chan<- string)
	GetPublisherID() string
}
