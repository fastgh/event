package event

type MsgLogEnum int

const (
	SubOk MsgLogEnum = iota
	UnSubOk

	RegisterTopicBegin
	RegisterTopicOk

	CloseHubBegin
	CloseHubOk
)

type EventLogEnum int

const (
	PubBegin EventLogEnum = iota
	PubOk

	SendBegin
	SendOk

	HandleBegin
	HandleOk

	TopicCloseBegin
	TopicCloseOk

	ListenerCloseBegin
	ListenerCloseOk
)

type ErrLogEnum int

const (
	ErrSubFailed ErrLogEnum = iota
	ErrUnSubFailed

	ErrPubFailed
	ErrSendFailed

	ErrHandleFailed
)

type Logger interface {
	Log(enm MsgLogEnum, hub string, topic string, lisner string)
	LogErr(enm ErrLogEnum, hub string, topic string, lisner string, err any)

	LogEvent(enm EventLogEnum, lisner string, evnt Event)
	LogEventErr(enm ErrLogEnum, hub string, topic string, lisner string, evnt Event, err any)
}
