package event

import (
	"fmt"
	"reflect"
	"sync"
)

type HubT struct {
	name   string
	mutex  sync.RWMutex
	topics map[string]TopicBase
	logger Logger
}

type Hub = *HubT

func NewHub(name string, logger Logger) Hub {
	return &HubT{
		name:   name,
		mutex:  sync.RWMutex{},
		topics: map[string]TopicBase{},
		logger: logger,
	}
}

func (me Hub) Name() string { return me.name }

func CreateTopic[K any](hub Hub, name string, eventExample K, logger Logger) Topic[K] {
	r := NewTopic(name, eventExample, logger)
	hub.registerTopic(r)
	return r
}

func (me Hub) registerTopic(topic TopicBase) {
	me.mutex.Lock()
	defer me.mutex.Unlock()

	name := topic.Name()
	_, has := me.topics[name]
	if has {
		panic(fmt.Errorf("topic %s already exists", name))
	}

	me.topics[name] = topic
}

func (me Hub) HasTopic(name string) bool {
	me.mutex.RLock()
	defer me.mutex.RUnlock()

	_, has := me.topics[name]
	return has
}

func GetTopic[K any](me Hub, name string, eventExample K) Topic[K] {
	r := me.GetTopic(name, eventExample)
	return r.(Topic[K])
}

func (me Hub) GetTopic(name string, eventExample any) TopicBase {
	me.mutex.RLock()
	defer me.mutex.RUnlock()

	r, has := me.topics[name]
	if has {
		expected := reflect.TypeOf(eventExample)
		actual := r.EventType()
		if expected != actual {
			panic(fmt.Errorf("expected topic event type is %v, however, the actual one is %v", expected, actual))
		}
	}

	return r
}

func (me Hub) Close(wait bool) {
	me.mutex.RLock()
	defer me.mutex.RUnlock()

	for _, topic := range me.topics {
		topic.Close(wait)
	}
}
