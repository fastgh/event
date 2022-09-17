package event

type Listener[K any] func(evnt K)

type EventListener[K any] struct {
	name  string
	lsner Listener[K]
	q     chan Event
	logr  ListenerLogger
}

func NewEventListener[K any](name string, lsner Listener[K], qSize int, topicLogr TopicLogger) *EventListener[K] {
	return &EventListener[K]{
		name:  name,
		lsner: lsner,
		q:     make(chan Event, qSize),
		logr:  NewListenerLogger(name, topicLogr),
	}
}

func (me *EventListener[K]) Stop(stopEvnt Event) {
	me.logr.LogEvent(ListenerCloseBegin, stopEvnt)
	me.q <- stopEvnt
}

func (me *EventListener[K]) Start() {
	go func() {
		for evnt := range me.q {
			if evnt.IsClose() {
				me.logr.LogEvent(ListenerCloseOk, evnt)
				break
			}

			me.onEvent(evnt)
		}
	}()
}

func (me *EventListener[K]) onEvent(evnt Event) {
	logr := me.logr

	defer func() {
		if p := recover(); p != nil {
			logr.LogEventErr(EventHandleErr, evnt, p)
		}
	}()

	logr.LogEvent(EventHandleBegin, evnt)

	var dat K = evnt.dat.(K)
	me.lsner(dat)

	logr.LogEvent(EventHandleOk, evnt)
}

func (me *EventListener[K]) SendEvent(evnt Event) {
	logr := me.logr

	defer func() {
		if p := recover(); p != nil {
			logr.LogEventErr(EventSendErr, evnt, p)
		}
	}()

	logr.LogEvent(EventSendBegin, evnt)
	me.q <- evnt
	logr.LogEvent(EventSendOk, evnt)
}
