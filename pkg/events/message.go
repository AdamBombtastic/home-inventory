package events

import "time"

type Message struct {
	CreatedAt time.Time
	Topic     string
	Payload   interface{}
	Type      string
}
