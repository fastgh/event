package event

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type EventId int64

type EventT struct {
	Id    EventId
	Hub   string
	Topic string
	Close bool
	Data  any
}

type Event = *EventT

func NewDataEvent(id EventId, hub string, topic string, dat any) Event {
	return &EventT{
		Id:    id,
		Hub:   hub,
		Topic: topic,
		Data:  dat,
		Close: false,
	}
}

func NewCloseEvent(id EventId, hub string, topic string) Event {
	return &EventT{
		Id:    id,
		Hub:   hub,
		Topic: topic,
		Data:  nil,
		Close: true,
	}
}

func (me Event) String() string {
	bytes, err := json.Marshal(me)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal event"))
	}
	return string(bytes)
}
