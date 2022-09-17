package event

import "fmt"

type Listener[K any] interface {
	On(evnt K)
}

type ListenerFn[K any] func(evnt K)

func (fn ListenerFn[K]) On(evnt K) {
	fn(evnt)
}

type QueueData[K any] struct {
	event  K
	closed bool
}

type ListenerItem[K any] struct {
	index    int
	name     string
	desc     string
	listener Listener[K]
	queue    chan *QueueData[K]
	logger   *LoggerAdapter[K]
}

func NewListenerItem[K any](index int, name string, listener Listener[K], queueSize int, logger *LoggerAdapter[K]) *ListenerItem[K] {
	return &ListenerItem[K]{
		index:    index,
		name:     name,
		desc:     fmt.Sprintf("listener '%s'(#%d)", name, index),
		listener: listener,
		queue:    make(chan *QueueData[K], queueSize),
		logger:   logger,
	}
}

func (me *ListenerItem[K]) String() string { return me.desc }

func (me *ListenerItem[K]) Stop() {
	me.queue <- &QueueData[K]{closed: true}
}

func (me *ListenerItem[K]) Start() {
	go func() {
		for data := range me.queue {
			if data.closed {
				break
			}

			me.onEvent(data)
		}

		me.logger.Info(LogTypeListenerQueueClosed, fmt.Sprint("closed queue for", me))
	}()
}

func (me *ListenerItem[K]) onEvent(data *QueueData[K]) {
	lgr := me.logger

	defer func() {
		if p := recover(); p != nil {
			lgr.Error(LogTypeListenerFailed, p, fmt.Sprint("failed to send event to", me))
		}
	}()

	lgr.Info(LogTypeListenerBegin, fmt.Sprint("begin to send event to", me))
	(me.listener).On(data.event)
	lgr.Info(LogTypeListenerOk, fmt.Sprint("successfully send event to ", me))
}

func (me *ListenerItem[K]) SendEvent(data *QueueData[K]) {
	lgr := me.logger

	defer func() {
		if p := recover(); p != nil {
			lgr.Error(LogTypeListenerFailed, p, fmt.Sprint("failed to send event to", me))
		}
	}()

	lgr.Info(LogTypeListenerBegin, fmt.Sprint("begin to send event to", me))
	me.queue <- data
	lgr.Info(LogTypeListenerOk, fmt.Sprint("successfully send event to", me))
}
