package event

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
)

type TopicImpl[K any] struct {
	name   string
	hub    Hub
	typ    reflect.Type
	lsners []*EventListener[K]
	waitG  sync.WaitGroup
	logr   TopicLogger
	eid    atomic.Int64
	mx     sync.RWMutex
}

func NewTopicImpl[K any](name string, hub Hub, example K, logr HubLogger) *TopicImpl[K] {
	return &TopicImpl[K]{
		name:   name,
		hub:    hub,
		typ:    reflect.TypeOf(example),
		lsners: []*EventListener[K]{},
		waitG:  sync.WaitGroup{},
		eid:    atomic.Int64{},
		mx:     sync.RWMutex{},
		logr:   NewTopicLogger(name, logr),
	}
}

func (me *TopicImpl[K]) EventType() reflect.Type { return me.typ }

func (me *TopicImpl[K]) Name() string { return me.name }

func (me *TopicImpl[K]) Hub() Hub { return me.hub }

func (me *TopicImpl[K]) CurrEventId() EventId { return EventId(me.eid.Load()) }

func (me *TopicImpl[K]) NewEventId() EventId { return EventId(me.eid.Add(1)) }

func (me *TopicImpl[K]) Sub(name string, lsner Listener[K], qSize int) int {
	if lsner == nil {
		panic(errors.New("listener cannot be nil"))
	}

	logr := me.logr

	me.mx.Lock()
	defer me.mx.Unlock()

	evntLsners := me.lsners
	for i, existing := range evntLsners {
		if existing.name == name {
			logr.LogError(ListenerSubErr, name, fmt.Sprintf("duplicated listener on #%d", i))
			return -1
		}
	}

	evntLsner := NewEventListener(name, lsner, qSize, logr)
	evntLsner.Start()

	evntLsners = append(evntLsners, evntLsner)
	me.lsners = evntLsners
	me.waitG.Add(1)

	logr.LogInfo(ListenerSubOk, name)
	return len(evntLsners)
}

func (me *TopicImpl[K]) UnSub(name string) bool {
	logr := me.logr

	me.mx.Lock()
	defer me.mx.Unlock()

	lsners := me.lsners
	for i, existing := range lsners {
		if existing.name == name {
			me.lsners = append(lsners[:i], lsners[i+1])
			logr.LogInfo(ListenerUnsubOk, name)

			stopEvent := NewCloseEvent(me.NewEventId(), me.Hub().Name(), me.name)
			me.stopListener(existing, stopEvent)
			return true
		}
	}

	logr.LogError(ListenerUnsubErr, name, "not found")
	return false
}

func (me *TopicImpl[K]) Pub(async bool, evnt K) {
	if async {
		go me.doPub(evnt)
	} else {
		me.doPub(evnt)
	}
}

func (me *TopicImpl[K]) doPub(evntData K) {
	evnt := NewDataEvent(me.NewEventId(), me.Hub().name, me.name, evntData)

	logr := me.logr

	defer func() {
		if p := recover(); p != nil {
			logr.LogError(EventPubErr, "", p)
		}
	}()

	me.mx.RLock()
	defer me.mx.RUnlock()

	logr.LogEventDebug(EventPubBegin, "", evnt)

	for _, lsner := range me.lsners {
		lsner.SendEvent(evnt)
	}

	logr.LogEventDebug(EventPubOk, "", evnt)
}

func (me *TopicImpl[K]) Close(wait bool) {
	me.mx.RLock()
	defer me.mx.RUnlock()

	stopEvnt := NewCloseEvent(me.NewEventId(), me.Hub().Name(), me.name)

	logr := me.logr
	logr.LogEventDebug(TopicCloseBegin, "", stopEvnt)

	for _, lsner := range me.lsners {
		me.stopListener(lsner, stopEvnt)
	}

	if wait {
		me.waitG.Wait()
	}

	logr.LogEventDebug(TopicCloseOk, "", stopEvnt)
}

func (me *TopicImpl[K]) stopListener(lsner *EventListener[K], stopEvnt Event) {
	lsner.Stop(stopEvnt)
	me.waitG.Done()
}
