package main

import (
	"fmt"

	"github.com/fastgh/event"
)

type MyEvent struct {
	Name string
}

func main() {
	myHub := event.NewHub("", nil)

	myTopic := event.CreateTopic(myHub, "myEvent", MyEvent{}, nil)

	myTopic.Sub("listener1", event.ListenerFn[MyEvent](func(e MyEvent) {
		fmt.Println("listener1 - got event from", e)
	}), 0)

	myTopic.Sub("listener2", event.ListenerFn[MyEvent](func(e MyEvent) {
		fmt.Println("listener2 - got event from", e)
	}), 0)

	myTopic.Pub(false, MyEvent{"fastgh"})

	myHub.Close(true)
}
