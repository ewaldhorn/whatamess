package main

// --------------------------------------------------------------------------------------- Publisher
type Publisher struct {
	subscriberList []Subscriber
}

func GetNewPublisher() Publisher {
	return Publisher{subscriberList: []Subscriber{}}
}

func (p *Publisher) Register(subs Subscriber) {
	p.subscriberList = append(p.subscriberList, subs)
}

func (p Publisher) NotifyAll(message string) {
	for _, subs := range p.subscriberList {
		subs.ReactToPublisherMessage(message)
	}
}

// -------------------------------------------------------------------------------------- Subscriber
type Subscriber struct {
	subscriberId string
}

func GetNewSubscriber(id string) Subscriber {
	return Subscriber{subscriberId: id}
}

func (s Subscriber) ReactToPublisherMessage(message string) {
	println("We received:", message, "For Sub#:", s.subscriberId)
}

func main() {
	pub := GetNewPublisher()

	pub.Register(GetNewSubscriber("1"))
	pub.Register(GetNewSubscriber("2"))

	pub.NotifyAll("Hello there folks!")
}
