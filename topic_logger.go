package event

type TopicLoggerT struct {
	topic string
	logr  HubLogger
}

type TopicLogger = *TopicLoggerT

func NewTopicLogger(topic string, logr HubLogger) TopicLogger {
	return &TopicLoggerT{
		logr:  logr,
		topic: topic,
	}
}

func (me TopicLogger) Topic() string { return me.topic }

func (me TopicLogger) Log(enm LogEnum, lsner string) {
	me.logr.Log(enm, me.topic, lsner)
}

func (me TopicLogger) LogEvent(enm LogEnum, lsner string, evnt Event) {
	me.logr.LogEvent(enm, lsner, evnt)
}

func (me TopicLogger) LogErr(enm LogEnum, lsner string, err any) {
	me.logr.LogErr(enm, me.topic, lsner, err)
}

func (me TopicLogger) LogEventErr(enm LogEnum, lsner string, evnt Event, err any) {
	me.logr.LogEventErr(enm, me.topic, lsner, evnt, err)
}
