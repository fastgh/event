package event

type LogEnum int64

const (
	HubCloseBegin LogEnum = iota
	HubCloseOk

	ListenerSubOk
	ListenerSubErr
	ListenerUnsubOk
	ListenerUnsubErr

	ListenerCloseBegin
	ListenerCloseOk

	TopicRegisterBegin
	TopicRegisterOk

	TopicCloseBegin
	TopicCloseOk

	EventPubBegin
	EventPubOk
	EventPubErr

	EventSendBegin
	EventSendOk
	EventSendErr

	EventHandleBegin
	EventHandleOk
	EventHandleErr
)
