package event

import (
	"fmt"
	"reflect"
	"sync"
)

type HubT struct {
	name   string
	mx     sync.RWMutex
	topics map[string]TopicBase
	logr   HubLogger
}

type Hub = *HubT

func (me Hub) Logger() HubLogger { return me.logr }

func (me Hub) Name() string { return me.name }

func (me Hub) RegisterTopic(topic TopicBase) {
	logr := me.logr

	me.mx.Lock()
	defer me.mx.Unlock()

	nm := topic.Name()
	logr.Log(RegisterTopicBegin, nm, "")

	if _, has := me.topics[nm]; has {
		panic(fmt.Errorf("duplicated topic '%s'", nm))
	}

	me.topics[nm] = topic

	logr.Log(RegisterTopicOk, nm, "")
}

func (me Hub) HasTopic(name string) bool {
	me.mx.RLock()
	defer me.mx.RUnlock()

	_, has := me.topics[name]
	return has
}

func (me Hub) GetTopic(name string, evntExample any) TopicBase {
	me.mx.RLock()
	defer me.mx.RUnlock()

	r, has := me.topics[name]
	if has {
		expTyp := reflect.TypeOf(evntExample)
		actualTyp := r.EventType()
		if expTyp != actualTyp {
			panic(fmt.Errorf("expected event type is %v, but got %v", expTyp, actualTyp))
		}
	}

	return r
}

func (me Hub) Close(wait bool) {
	logr := me.logr

	me.mx.RLock()
	defer me.mx.RUnlock()

	logr.Log(CloseHubBegin, "", "")

	for _, tp := range me.topics {
		tp.Close(wait)
	}

	logr.Log(CloseHubOk, "", "")
}
