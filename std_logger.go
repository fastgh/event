package event

type StdLoggerT struct {
}

type StdLogger = *StdLoggerT

func (me StdLogger) Log(enm LogEnum, hub string, topic string, lsner string) {

}

func (me StdLogger) LogErr(enm LogEnum, hub string, topic string, lsner string, err any) {

}

func (me StdLogger) LogEvent(enm LogEnum, lsner string, evnt Event) {

}

func (me StdLogger) LogEventErr(enm LogEnum, hub string, topic string, lsner string, evnt Event, err any) {

}
