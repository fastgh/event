package main

import (
	"fmt"

	fgevent "github.com/fastgh/go-event"
	"github.com/fastgh/go-event/loggers/std"
)

type MyEvent struct {
	Name string
}

func main() {
	myHub := fgevent.NewHub("default", std.NewDefaultGlobalStdLogger())

	myTopic := fgevent.CreateTopic(myHub, "myTopic", MyEvent{})

	myTopic.Sub("listener1", func(e MyEvent) {
		fmt.Println("listener1 - got event from", e)
	}, 0)

	myTopic.Sub("listener2", func(e MyEvent) {
		fmt.Println("listener2 - got event from", e)
	}, 0)

	myTopic.Pub(fgevent.PubModeAuto, MyEvent{"fastgh"})

	myHub.Close(true)
}
