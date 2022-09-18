package std

import (
	"log"

	fgevent "github.com/fastgh/go-event"
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

func NewDefaultStdLogger(target *log.Logger) fgevent.LevelFilteringLogger {
	return fgevent.NewLevelFilteringLogger(fgevent.LogLevelInfo, &StdLoggerT{
		target: target,
	})
}

func NewDefaultGlobalStdLogger() fgevent.LevelFilteringLogger {
	return NewDefaultStdLogger(log.Default())
}

func (me StdLogger) Target() *log.Logger { return me.target }

func (me StdLogger) LogDebug(enm fgevent.LogEnum, hub string, topic string, lsner string) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s --> %v", fgevent.LogLevelDebug, hub, topic, lsner, enm)
}

func (me StdLogger) LogInfo(enm fgevent.LogEnum, hub string, topic string, lsner string) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s --> %v", fgevent.LogLevelInfo, hub, topic, lsner, enm)
}

func (me StdLogger) LogError(enm fgevent.LogEnum, hub string, topic string, lsner string, err any) {
	me.target.Printf("<%v> hub=%s, topic=%s, listener=%s, error=%v --> %v", fgevent.LogLevelError, hub, topic, lsner, err, enm)
}

func (me StdLogger) LogEventDebug(enm fgevent.LogEnum, lsner string, evnt fgevent.Event) {
	me.target.Printf("<%v> event=%v, listener=%s --> %v", fgevent.LogLevelDebug, evnt, lsner, enm)
}

func (me StdLogger) LogEventInfo(enm fgevent.LogEnum, lsner string, evnt fgevent.Event) {
	me.target.Printf("<%v> event=%v, listener=%s --> %v", fgevent.LogLevelInfo, evnt, lsner, enm)
}

func (me StdLogger) LogEventError(enm fgevent.LogEnum, lsner string, evnt fgevent.Event, err any) {
	me.target.Printf("<%v> event=%v, listener=%s, error=%v --> %v", fgevent.LogLevelError, evnt, lsner, err, enm)
}
