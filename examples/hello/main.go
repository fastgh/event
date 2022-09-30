package main

import (
	"fmt"

	"github.com/fastgh/go-event"
	"github.com/fastgh/go-event/loggers/std"
)

type MyEvent struct {
	Name string
}

func main() {
	myHub := event.NewHub("default", std.NewDefaultGlobalStdLogger())

	myTopic := event.CreateTopic(myHub, "myTopic", MyEvent{})

	myTopic.SubP("listener1", func(e MyEvent) {
		fmt.Println("listener1 - got event from", e)
	}, 0)

	myTopic.SubP("listener2", func(e MyEvent) {
		fmt.Println("listener2 - got event from", e)
	}, 0)

	myTopic.Pub(event.PubModeAuto, MyEvent{"fastgh"})

	myHub.Close(true)
}
