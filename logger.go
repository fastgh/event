package event

type Logger interface {
	Log(enm LogEnum, hub string, topic string, lsner string)
	LogErr(enm LogEnum, hub string, topic string, lsner string, err any)

	LogEvent(enm LogEnum, lsner string, evnt Event)
	LogEventErr(enm LogEnum, hub string, topic string, lsner string, evnt Event, err any)
}
