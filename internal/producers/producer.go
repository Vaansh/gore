package producers

type Producer interface {
	ProduceOn(c chan<- string)
	ProducerID() string
}
