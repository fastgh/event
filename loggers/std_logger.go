package loggers

import (
	"log"

	"github.com/fastgh/event"
)

type StdLoggerT struct {
	target *log.Logger
}

type StdLogger = *StdLoggerT

func NewStdLogger(target *log.Logger) StdLogger {
	return &StdLoggerT{
		target: target,
	}
}

func NewDefaultStdLogger(target *log.Logger) event.LevelFilteringLogger {
	return event.NewLevelFilteringLogger(event.LogLevelInfo, &StdLoggerT{
		target: target,
	})
}

func NewDefaultGlobalStdLogger() event.LevelFilteringLogger {
	return NewDefaultStdLogger(log.Default())
}

func (me StdLogger) Target() *log.Logger { return me.target }

func (me StdLogger) LogDebug(enm event.LogEnum, hub string, topic string, lsner string) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s --> %v", event.LogLevelDebug, hub, topic, lsner, enm)
}

func (me StdLogger) LogInfo(enm event.LogEnum, hub string, topic string, lsner string) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s --> %v", event.LogLevelInfo, hub, topic, lsner, enm)
}

func (me StdLogger) LogError(enm event.LogEnum, hub string, topic string, lsner string, err any) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s, error=%v --> %v", event.LogLevelError, hub, topic, lsner, err, enm)
}

func (me StdLogger) LogEventDebug(enm event.LogEnum, lsner string, evnt event.Event) {
	me.target.Printf("<%v> event=%v, listener=%s --> %v", event.LogLevelDebug, evnt, lsner, enm)
}

func (me StdLogger) LogEventInfo(enm event.LogEnum, lsner string, evnt event.Event) {
	me.target.Printf("<%v> event=%v, listener=%s --> %v", event.LogLevelInfo, evnt, lsner, enm)
}

func (me StdLogger) LogEventError(enm event.LogEnum, lsner string, evnt event.Event, err any) {
	me.target.Printf("<%v> event=%v, listener=%s, error=%v --> %v", event.LogLevelError, evnt, lsner, err, enm)
}
