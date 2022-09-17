package loggers

import "github.com/fastgh/event"

type LevelFilteringLoggerT struct {
	target event.Logger
	Level  event.LogLevel
}

type LevelFilteringLogger = *LevelFilteringLoggerT

func NewLevelFilteringLogger(level event.LogLevel, target event.Logger) LevelFilteringLogger {
	return &LevelFilteringLoggerT{
		Level:  level,
		target: target,
	}
}

func (me LevelFilteringLogger) Target() event.Logger { return me.target }

func (me LevelFilteringLogger) LogDebug(enm event.LogEnum, hub string, topic string, lsner string) {
	if me.Level <= event.LogLevelDebug {
		me.target.LogDebug(enm, hub, topic, lsner)
	}
}

func (me LevelFilteringLogger) LogInfo(enm event.LogEnum, hub string, topic string, lsner string) {
	if me.Level <= event.LogLevelInfo {
		me.target.LogInfo(enm, hub, topic, lsner)
	}
}

func (me LevelFilteringLogger) LogError(enm event.LogEnum, hub string, topic string, lsner string, err any) {
	if me.Level <= event.LogLevelError {
		me.target.LogError(enm, hub, topic, lsner, err)
	}
}

func (me LevelFilteringLogger) LogEventDebug(enm event.LogEnum, lsner string, evnt event.Event) {
	if me.Level <= event.LogLevelDebug {
		me.target.LogEventDebug(enm, lsner, evnt)
	}
}

func (me LevelFilteringLogger) LogEventInfo(enm event.LogEnum, lsner string, evnt event.Event) {
	if me.Level <= event.LogLevelInfo {
		me.target.LogEventInfo(enm, lsner, evnt)
	}
}

func (me LevelFilteringLogger) LogEventError(enm event.LogEnum, lsner string, evnt event.Event, err any) {
	if me.Level <= event.LogLevelError {
		me.target.LogEventError(enm, lsner, evnt, err)
	}
}
