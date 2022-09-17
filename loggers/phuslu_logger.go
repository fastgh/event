package loggers

import (
	"fmt"
	"os"

	"github.com/fastgh/event"
	phuslu "github.com/phuslu/log"
)

type PhusluLoggerT struct {
	target *phuslu.Logger
}

type PhusluLogger = *PhusluLoggerT

func NewPhusluLogger(target *phuslu.Logger) PhusluLogger {
	return &PhusluLoggerT{
		target: target,
	}
}

func NewDefaultPhusluLogger() PhusluLogger {
	return NewPhusluLogger(
		&phuslu.Logger{
			Level:      phuslu.InfoLevel,
			Caller:     0,
			TimeField:  "",
			TimeFormat: "",
			Writer:     &phuslu.IOWriter{Writer: os.Stdout},
		},
	)
}

func (me PhusluLogger) Target() *phuslu.Logger { return me.target }

func (me PhusluLogger) LogDebug(enm event.LogEnum, hub string, topic string, lsner string) {
	me.target.Debug().Str("hub", hub).Str("topic", topic).Str("listener", lsner).Msg(enm.String())
}

func (me PhusluLogger) LogInfo(enm event.LogEnum, hub string, topic string, lsner string) {
	me.target.Info().Str("hub", hub).Str("topic", topic).Str("listener", lsner).Msg(enm.String())
}

func (me PhusluLogger) LogError(enm event.LogEnum, hub string, topic string, lsner string, err any) {
	entry := me.target.Error().Str("hub", hub).Str("topic", topic).Str("listener", lsner)
	PhusluMarshalAnyError(entry, err)
	entry.Msg(enm.String())
}

func (me PhusluLogger) LogEventDebug(enm event.LogEnum, lsner string, evnt event.Event) {
	me.target.Debug().Object("event", evnt).Str("listener", lsner).Msg(enm.String())
}

func (me PhusluLogger) LogEventInfo(enm event.LogEnum, lsner string, evnt event.Event) {
	me.target.Info().Object("event", evnt).Str("listener", lsner).Msg(enm.String())
}

func (me PhusluLogger) LogEventError(enm event.LogEnum, lsner string, evnt event.Event, err any) {
	entry := me.target.Error().Object("event", evnt).Str("listener", lsner)
	PhusluMarshalAnyError(entry, err)
	entry.Msg(enm.String())
}

func PhusluLogLevel(level event.LogLevel) phuslu.Level {
	switch level {
	case event.LogLevelAnyway:
		return phuslu.TraceLevel

	case event.LogLevelDebug:
		return phuslu.DebugLevel
	case event.LogLevelInfo:
		return phuslu.InfoLevel
	case event.LogLevelError:
		return phuslu.ErrorLevel

	case event.LogLevelSilient:
		return phuslu.ErrorLevel

	default:
		return phuslu.ErrorLevel
	}
}

func PhusluMarshalAnyError(entry *phuslu.Entry, anyErr any) {
	if anyErr == nil {
		return
	}

	if err, is := anyErr.(error); is {
		entry.Err(err)
		return
	}

	if str, is := anyErr.(string); is {
		entry.Str("err", str)
		return
	}

	entry.Str("err", fmt.Sprintf("%v", anyErr))
}
