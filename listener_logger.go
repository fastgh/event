package event

type ListenerLoggerT struct {
	logr  TopicLogger
	lsner string
}

type ListenerLogger = *ListenerLoggerT

func NewListenerLogger(lsner string, logr TopicLogger) ListenerLogger {
	return &ListenerLoggerT{
		logr:  logr,
		lsner: lsner,
	}
}

func (me ListenerLogger) Listener() string { return me.lsner }

func (me ListenerLogger) Log(enm LogEnum) {
	me.logr.Log(enm, me.lsner)
}

func (me ListenerLogger) LogEvent(enm LogEnum, evnt Event) {
	me.logr.LogEvent(enm, me.lsner, evnt)
}

func (me ListenerLogger) LogErr(enm LogEnum, err any) {
	me.logr.LogErr(enm, me.lsner, err)
}

func (me ListenerLogger) LogEventErr(enm LogEnum, evnt Event, err any) {
	me.logr.LogEventErr(enm, me.lsner, evnt, err)
}
