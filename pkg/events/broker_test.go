package events

import (
	"testing"
	"time"
)

func Test_BrokerSubscribe(t *testing.T) {
	b := NewBroker()
	topic := "test"
	handler := func(message *Message) error {
		return nil
	}
	if err := b.Subscribe(topic, handler); err != nil {
		t.Error("Expected nil, got", err)
	}
}

func Test_BrokerPublish(t *testing.T) {
	b := NewBroker()
	topic := "test"
	count := 0
	handler := func(message *Message) error {
		count += 1
		return nil
	}
	if err := b.Subscribe(topic, handler); err != nil {
		t.Error("Expected nil, got", err)
	}
	message := &Message{
		Topic:     topic,
		CreatedAt: time.Now(),
		Payload:   "foo",
		Type:      "bar",
	}
	if err := b.Publish(message); err != nil {
		t.Error("Expected nil, got", err)
	}
	if count != 1 {
		t.Error("Expected 1, got", count)
	}
}
