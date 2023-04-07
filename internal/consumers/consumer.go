package consumers

type Consumer interface {
	ConsumeOn(c <-chan string)
	ConsumerID() string
}
