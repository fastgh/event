package event

import "reflect"

type LogType int

const (
	LogTypeListenerSubOk LogType = iota
	LogTypeListenerSubFailed
	LogTypeListenerUnSubOk
	LogTypeListenerUnSubFailed

	LogTypePubBegin
	LogTypePubOk
	LogTypePubFailed

	LogTypeListenerBegin
	LogTypeListenerOk
	LogTypeListenerFailed
	LogTypeListenerQueueClosed
)

type Logger interface {
	logInfo(name string, typ LogType, eventId uint32, msg string)
	logError(name string, typ LogType, eventId uint32, err any, msg string)
}

type Listener[K any] interface {
	On(evnt K)
}

type ListenerFn[K any] func(evnt K)

func (fn ListenerFn[K]) On(evnt K) {
	fn(evnt)
}

type Topic[K any] interface {
	Name() string
	Logger() *LoggerAdapter[K]
	EventId() uint32
	EventType() reflect.Type

	Sub(name string, lisner Listener[K], queueSize int) int
	UnSub(lisner Listener[K]) bool
	Pub(async bool, evnt K)
}

func NewTopic[K any](name string, eventExample K, logger Logger) Topic[K] {
	return NewTopicImpl(name, eventExample, logger)
}
