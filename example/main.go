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

	myTopic := event.CreateTopic[MyEvent](myHub, "myEvent", MyEvent{}, nil)

	myTopic.Sub("listener1", event.ListenerFn[MyEvent](func(e MyEvent) {
		fmt.Printf("got event: %v", e)
	}), 0)

	myTopic.Pub(false, MyEvent{"fastgh"})
}
