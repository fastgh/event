package event

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
)

type TopicBase interface {
	Name() string
	EventId() uint32
	EventType() reflect.Type

	UnSub(name string) bool
	Close(wait bool)
}

type Topic[K any] interface {
	TopicBase

	Logger() *LoggerAdapter[K]

	Sub(name string, lisner Listener[K], queueSize int) int
	Pub(async bool, evnt K)
}

func NewTopic[K any](name string, eventExample K, logger Logger) Topic[K] {
	return NewTopicImpl(name, eventExample, logger)
}

type TopicImpl[K any] struct {
	name          string
	eventType     reflect.Type
	listenerItems []*ListenerItem[K]
	waitGroup     sync.WaitGroup
	logger        *LoggerAdapter[K]
	eventId       atomic.Uint32
	mutex         sync.RWMutex
}

func NewTopicImpl[K any](name string, example K, logger Logger) *TopicImpl[K] {
	r := &TopicImpl[K]{
		name:          name,
		eventType:     reflect.TypeOf(example),
		listenerItems: []*ListenerItem[K]{},
		waitGroup:     sync.WaitGroup{},
		eventId:       atomic.Uint32{},
		mutex:         sync.RWMutex{},
	}
	r.logger = NewLoggerAdapter[K](r, logger)
	return r
}

func (me *TopicImpl[K]) EventType() reflect.Type { return me.eventType }

func (me *TopicImpl[K]) Name() string { return me.name }

func (me *TopicImpl[K]) Logger() *LoggerAdapter[K] { return me.logger }

func (me *TopicImpl[K]) EventId() uint32 { return me.eventId.Load() }

func (me *TopicImpl[K]) Sub(name string, lisner Listener[K], queueSize int) int {
	if lisner == nil {
		panic(errors.New("listener cannot be nil"))
	}

	lgr := me.logger

	me.mutex.Lock()
	defer me.mutex.Unlock()

	items := me.listenerItems
	for i, existingItem := range items {
		if existingItem.name == name {
			lgr.Info(LogTypeListenerSubFailed, fmt.Sprintf("listener '%s' is duplicated on #%d", name, i))
			return -1
		}
	}

	item := NewListenerItem(len(items), name, lisner, queueSize, lgr)
	item.Start()

	items = append(items, item)
	me.listenerItems = items
	me.waitGroup.Add(1)

	lgr.Info(LogTypeListenerSubOk, fmt.Sprint("added", item))

	return len(items)
}

func (me *TopicImpl[K]) stopItem(item *ListenerItem[K]) {
	item.Stop()
	me.waitGroup.Done()
}

func (me *TopicImpl[K]) UnSub(name string) bool {
	lgr := me.logger

	me.mutex.Lock()
	defer me.mutex.Unlock()

	items := me.listenerItems
	for i, existingItem := range items {
		if existingItem.name == name {
			me.listenerItems = append(items[:i], items[i+1])
			lgr.Info(LogTypeListenerUnSubOk, fmt.Sprint("removed", existingItem))

			me.stopItem(existingItem)
			return true
		}
	}

	lgr.Info(LogTypeListenerUnSubFailed, fmt.Sprintf("listener %v is not found", name))
	return false
}

func (me *TopicImpl[K]) Pub(async bool, evnt K) {
	if async {
		go me.pub(evnt)
	} else {
		me.pub(evnt)
	}
}

func (me *TopicImpl[K]) pub(evnt K) {
	me.eventId.Add(1)

	lgr := me.logger
	lgr.Info(LogTypePubBegin, fmt.Sprint("begin to publish event:", evnt))

	defer func() {
		if p := recover(); p != nil {
			lgr.Error(LogTypePubFailed, p, "publish routine got panic")
		}
	}()

	data := &QueueData[K]{
		closed: false,
		event:  evnt,
	}

	me.mutex.RLock()
	defer me.mutex.RUnlock()

	for _, itm := range me.listenerItems {
		itm.SendEvent(data)
	}
}

func (me *TopicImpl[K]) Close(wait bool) {
	lgr := me.logger
	lgr.Info(LogTypePubBegin, "begin to close")

	me.mutex.RLock()
	defer me.mutex.RUnlock()

	for _, itm := range me.listenerItems {
		me.stopItem(itm)
	}

	if wait {
		me.waitGroup.Wait()
	}

	lgr.Info(LogTypePubBegin, "successfully close")
}
