package benchmark

import (
	"fmt"
	"testing"

	"github.com/fastgh/event"
)

type BenchmarkEvent struct {
	Name    string
	Address string
}

func BenchmarkRepeat(b *testing.B) {
	myHub := event.NewHub("benchmark", nil)
	myTopic := event.CreateTopic(myHub, "myTopic", BenchmarkEvent{})

	for i := 0; i < 100; i++ {
		myTopic.Sub(fmt.Sprintf("listener %d", i), func(e BenchmarkEvent) {
			// do nothing
		}, 0)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		myTopic.Pub(false, BenchmarkEvent{Name: "fastgh", Address: "github.com/fastgh/event"})
	}

	myHub.Close(true)
}
