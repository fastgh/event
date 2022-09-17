package event

type Listener[K any] func(evnt K)

type EventListener[K any] struct {
	name   string
	lisner Listener[K]
	q      chan Event
	logr   ListenerLogger
}

func NewEventListener[K any](name string, lsner Listener[K], qSize int, topicLogr TopicLogger) *EventListener[K] {
	return &EventListener[K]{
		name:   name,
		lisner: lsner,
		q:      make(chan Event, qSize),
		logr:   NewListenerLogger(name, topicLogr),
	}
}

func (me *EventListener[K]) Stop(stopEvnt Event) {
	me.logr.LogEvent(LogEventListenerCloseBegin, stopEvnt)
	me.q <- stopEvnt
}

func (me *EventListener[K]) Start() {
	go func() {
		for evnt := range me.q {
			if evnt.IsClose() {
				me.logr.LogEvent(LogEventListenerCloseOk, evnt)
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
			logr.LogEventErr(LogErrHandleFailed, evnt, p)
		}
	}()

	logr.LogEvent(LogEventHandleBegin, evnt)

	var dat K = evnt.dat.(K)
	me.lisner(dat)

	logr.LogEvent(LogEventHandleOk, evnt)
}

func (me *EventListener[K]) SendEvent(evnt Event) {
	logr := me.logr

	defer func() {
		if p := recover(); p != nil {
			logr.LogEventErr(LogErrSendFailed, evnt, p)
		}
	}()

	logr.LogEvent(LogEventSendBegin, evnt)
	me.q <- evnt
	logr.LogEvent(LogEventSendOk, evnt)
}
