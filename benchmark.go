package benchmark

import (
	"fmt"
	"time"
)

// Benchmark struct
type benchmark struct {
	startTime time.Time
	endTime   time.Time
	duration  time.Duration
}

// Start a new benchmark
func (b *benchmark) startBenchmark() {
	b.startTime = time.Now()
}

// End the benchmark then record the duration
func (b *benchmark) endBenchmark() {
	b.endTime = time.Now()
	b.duration = b.endTime.Sub(b.startTime)
}

// Display the recorded duration in milliseconds
func (b *benchmark) displayResultInMilliseconds(label string) {
	durationInMs := float64(b.duration.Nanoseconds()) * 0.000001
	fmt.Printf("%v %fms\n", label, durationInMs)
}

// Display the recorded duration in seconds
func (b *benchmark) displayResultInSeconds(label string) {
	durationInSecs := float64(b.duration.Nanoseconds()) * 0.000000001
	fmt.Printf("%v %fseconds\n", label, durationInSecs)
}
