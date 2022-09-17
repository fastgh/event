package event

type HubLoggerT struct {
	hub  string
	logr Logger
}

type HubLogger = *HubLoggerT

func NewHubLogger(hub string, logr Logger) HubLogger {
	return &HubLoggerT{
		hub:  hub,
		logr: logr,
	}
}

func (me HubLogger) Hub() string {
	return me.hub
}

func (me HubLogger) Log(enm LogEnum, topic string, lsner string) {
	if me.logr != nil {
		me.logr.Log(enm, me.hub, topic, lsner)
	}
}

func (me HubLogger) LogErr(enm LogEnum, topic string, lsner string, err any) {
	if me.logr != nil {
		me.logr.LogErr(enm, me.hub, topic, lsner, err)
	}
}

func (me HubLogger) LogEvent(enm LogEnum, lsner string, evnt Event) {
	if me.logr != nil {
		me.logr.LogEvent(enm, lsner, evnt)
	}
}

func (me HubLogger) LogEventErr(enm LogEnum, topic string, lsner string, evnt Event, err any) {
	me.logr.LogEventErr(enm, me.hub, topic, lsner, evnt, err)
}
