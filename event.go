package event

type EventId int64

type EventT struct {
	id    EventId
	hub   string
	topic string
	cloz  bool
	dat   any
}

type Event = *EventT

func NewDataEvent(id EventId, hub string, topic string, dat any) Event {
	return &EventT{
		id:    id,
		hub:   hub,
		topic: topic,
		dat:   dat,
		cloz:  false,
	}
}

func NewCloseEvent(id EventId, hub string, topic string) Event {
	return &EventT{
		id:    id,
		hub:   hub,
		topic: topic,
		dat:   nil,
		cloz:  true,
	}
}

func (me Event) IsClose() bool {
	return me.cloz
}

func (me Event) Id() EventId {
	return me.id
}

func (me Event) Data() any {
	return me.dat
}

func (me Event) Hub() string {
	return me.hub
}

func (me Event) Topic() string {
	return me.topic
}
