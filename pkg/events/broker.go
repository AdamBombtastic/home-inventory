package events

type HandlerFunc func(message *Message) error

type Broker interface {
	Publish(message *Message) error
	Subscribe(topic string, handler HandlerFunc) error
}

type broker struct {
	subscriptions map[string][]HandlerFunc
}

func NewBroker() Broker {
	return &broker{
		subscriptions: make(map[string][]HandlerFunc),
	}
}

func (b *broker) Publish(message *Message) error {
	for k, v := range b.subscriptions {
		if k == message.Topic {
			for _, handler := range v {
				if err := handler(message); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (b *broker) Subscribe(topic string, handler HandlerFunc) error {
	b.subscriptions[topic] = append(b.subscriptions[topic], handler)
	return nil
}
