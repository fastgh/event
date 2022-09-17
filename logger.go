package event

type LogMsg int

const (
	LogMsgSubOk LogMsg = iota
	LogMsgUnSubOk

	LogMsgRegisterTopicBegin
	LogMsgRegisterTopicOk

	LogMsgCloseHubBegin
	LogMsgCloseHubOk
)

type LogEvent int

const (
	LogEventPubBegin LogEvent = iota
	LogEventPubOk

	LogEventSendBegin
	LogEventSendOk

	LogEventHandleBegin
	LogEventHandleOk

	LogEventTopicCloseBegin
	LogEventTopicCloseOk

	LogEventListenerCloseBegin
	LogEventListenerCloseOk
)

type LogErr int

const (
	LogErrSubFailed LogErr = iota
	LogErrUnSubFailed

	LogErrPubFailed
	LogErrSendFailed

	LogErrHandleFailed
)

type Logger interface {
	LogInfo(typ LogMsg, hub string, topic string, lisner string)
	LogEvent(typ LogEvent, lisner string, evnt Event)
	LogErr(typ LogErr, hub string, topic string, lisner string, err any)
	LogEventErr(typ LogErr, hub string, topic string, lisner string, evnt Event, err any)
}
