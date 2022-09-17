package event

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

type LoggerAdapter[K any] struct {
	topic  Topic[K]
	logger Logger
}

func NewLoggerAdapter[K any](topic Topic[K], logger Logger) *LoggerAdapter[K] {
	return &LoggerAdapter[K]{
		topic:  topic,
		logger: logger,
	}
}

func (me *LoggerAdapter[K]) Info(typ LogType, msg string) {
	if me.logger != nil {
		t := me.topic
		me.logger.logInfo(t.Name(), typ, t.EventId(), msg)
	}
}

func (me *LoggerAdapter[K]) Error(typ LogType, err any, msg string) {
	if me.logger != nil {
		t := me.topic
		me.logger.logError(t.Name(), typ, t.EventId(), err, msg)
	}
}
