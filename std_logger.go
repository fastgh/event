package event

import "log"

type StdLoggerT struct {
	target *log.Logger
}

type StdLogger = *StdLoggerT

func NewStdLogger(target *log.Logger) StdLogger {
	return &StdLoggerT{
		target: target,
	}
}

func NewDefaultStdLogger(target *log.Logger) LevelFilteringLogger {
	return NewLevelFilteringLogger(LogLevelInfo, &StdLoggerT{
		target: target,
	})
}

func NewDefaultGlobalStdLogger() LevelFilteringLogger {
	return NewDefaultStdLogger(log.Default())
}

func (me StdLogger) Target() *log.Logger { return me.target }

func (me StdLogger) LogDebug(enm LogEnum, hub string, topic string, lsner string) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s --> %v", LogLevelDebug, hub, topic, lsner, enm)
}

func (me StdLogger) LogInfo(enm LogEnum, hub string, topic string, lsner string) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s --> %v", LogLevelInfo, hub, topic, lsner, enm)
}

func (me StdLogger) LogError(enm LogEnum, hub string, topic string, lsner string, err any) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s, error=%v --> %v", LogLevelError, hub, topic, lsner, err, enm)
}

func (me StdLogger) LogEventDebug(enm LogEnum, lsner string, evnt Event) {
	me.target.Printf("<%v> event=%v, listener=%s --> %v", LogLevelDebug, evnt, lsner, enm)
}

func (me StdLogger) LogEventInfo(enm LogEnum, lsner string, evnt Event) {
	me.target.Printf("<%v> event=%v, listener=%s --> %v", LogLevelInfo, evnt, lsner, enm)
}

func (me StdLogger) LogEventError(enm LogEnum, lsner string, evnt Event, err any) {
	me.target.Printf("<%v> event=%v, listener=%s, error=%v --> %v", LogLevelError, evnt, lsner, err, enm)
}
