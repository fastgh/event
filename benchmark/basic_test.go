package benchmark

import (
	"fmt"
	"testing"

	fgevent "github.com/fastgh/go-event"
)

type BenchmarkEvent struct {
	Name    string
	Address string
}

func Benchmark_1_listeners_and_queueSize_is_0___serial(b *testing.B) {
	doBenchmark(b, 1, 0, false)
}

func Benchmark_1_listeners_and_queueSize_is_0_parallel(b *testing.B) {
	doBenchmark(b, 1, 1, true)
}

func Benchmark_1_listeners_and_queueSize_is_1___serial(b *testing.B) {
	doBenchmark(b, 1, 1, false)
}

func Benchmark_1_listeners_and_queueSize_is_1_parallel(b *testing.B) {
	doBenchmark(b, 1, 1, false)
}

func Benchmark_1_listeners_and_queueSize_is_2___serial(b *testing.B) {
	doBenchmark(b, 1, 2, false)
}

func Benchmark_1_listeners_and_queueSize_is_2_parallel(b *testing.B) {
	doBenchmark(b, 1, 2, true)
}

func Benchmark_1_listeners_and_queueSize_is_10___serial(b *testing.B) {
	doBenchmark(b, 1, 10, false)
}

func Benchmark_1_listeners_and_queueSize_is_10_parallel(b *testing.B) {
	doBenchmark(b, 1, 10, true)
}

func Benchmark_1_listeners_and_queueSize_is_100___serial(b *testing.B) {
	doBenchmark(b, 1, 100, false)
}

func Benchmark_1_listeners_and_queueSize_is_100_parallel(b *testing.B) {
	doBenchmark(b, 1, 100, true)
}

func Benchmark_2_listeners_and_queueSize_is_100___serial(b *testing.B) {
	doBenchmark(b, 2, 100, false)
}

func Benchmark_2_listeners_and_queueSize_is_100_parallel(b *testing.B) {
	doBenchmark(b, 2, 100, true)
}

func Benchmark_10_listeners_and_queueSize_is_100___serial(b *testing.B) {
	doBenchmark(b, 10, 100, false)
}

func Benchmark_10_listeners_and_queueSize_is_100_parallel(b *testing.B) {
	doBenchmark(b, 10, 100, true)
}

func Benchmark_20_listeners_and_queueSize_is_100___serial(b *testing.B) {
	doBenchmark(b, 20, 100, false)
}

func Benchmark_20_listeners_and_queueSize_is_100_parallel(b *testing.B) {
	doBenchmark(b, 20, 100, true)
}

func Benchmark_100_listeners_and_queueSize_is_100___serial(b *testing.B) {
	doBenchmark(b, 100, 100, false)
}

func Benchmark_100_listeners_and_queueSize_is_100_parallel(b *testing.B) {
	doBenchmark(b, 100, 100, true)
}

func Benchmark_200_listeners_and_queueSize_is_100___serial(b *testing.B) {
	doBenchmark(b, 200, 100, false)
}

func Benchmark_200_listeners_and_queueSize_is_100_parallel(b *testing.B) {
	doBenchmark(b, 200, 100, true)
}

func Benchmark_1000_listeners_and_queueSize_is_100___serial(b *testing.B) {
	doBenchmark(b, 1000, 100, false)
}

func Benchmark_1000_listeners_and_queueSize_is_100_parallel(b *testing.B) {
	doBenchmark(b, 1000, 100, true)
}

func doBenchmark(b *testing.B, amountOfSub int, queueSize uint32, parallel bool) {
	myHub := fgevent.NewHub("benchmark", nil)
	myTopic := fgevent.CreateTopic(myHub, "myTopic", BenchmarkEvent{})

	for i := 0; i < amountOfSub; i++ {
		myTopic.Sub(fmt.Sprintf("sub %d", i), func(e BenchmarkEvent) {
			// do nothing
		}, queueSize)
	}

	b.ResetTimer()

	if parallel {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				myTopic.Pub(fgevent.PubModeSync, BenchmarkEvent{Name: "fastgh", Address: "github.com/fastgh/fgevent"})
			}
		})
	} else {
		for i := 0; i < b.N; i++ {
			myTopic.Pub(fgevent.PubModeAuto, BenchmarkEvent{Name: "fastgh", Address: "github.com/fastgh/fgevent"})
		}
	}

	b.StopTimer()

	myHub.Close(true)
}
