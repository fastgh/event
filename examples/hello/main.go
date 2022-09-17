package main

import (
	"fmt"

	"github.com/fastgh/event"
	"github.com/fastgh/event/loggers"
)

type MyEvent struct {
	Name string
}

func main() {
	myHub := event.NewHub("default", loggers.NewDefaultGlobalStdLogger())

	myTopic := event.CreateTopic(myHub, "myTopic", MyEvent{})

	myTopic.Sub("listener1", func(e MyEvent) {
		fmt.Println("listener1 - got event from", e)
	}, 0)

	myTopic.Sub("listener2", func(e MyEvent) {
		fmt.Println("listener2 - got event from", e)
	}, 0)

	myTopic.Pub(false, MyEvent{"fastgh"})

	myHub.Close(true)
}
