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

func Benchmark_10_listeners_and_queueSize_is_100(b *testing.B) {
	doBenchmark(b, 10, 100)
}

func doBenchmark(b *testing.B, amountOfSub int, queueSize uint32) {
	myHub := event.NewHub("benchmark", nil)
	myTopic := event.CreateTopic(myHub, "myTopic", BenchmarkEvent{})

	for i := 0; i < amountOfSub; i++ {
		myTopic.Sub(fmt.Sprintf("sub %d", i), func(e BenchmarkEvent) {
			// do nothing
		}, queueSize)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		myTopic.Pub(false, BenchmarkEvent{Name: "fastgh", Address: "github.com/fastgh/event"})
	}

	myHub.Close(true)
}
