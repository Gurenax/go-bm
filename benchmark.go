package benchmark

import (
	"fmt"
	"time"
)

// Benchmark struct
type BM struct {
	startTime time.Time
	endTime   time.Time
	duration  time.Duration
}

// Start a new benchmark
func (b *BM) Start() {
	b.startTime = time.Now()
}

// End the benchmark then record the duration
func (b *BM) End() {
	b.endTime = time.Now()
	b.duration = b.endTime.Sub(b.startTime)
}

// Display the recorded duration in nanoseconds
func (b *BM) DurationInNanoseconds(label string) {
	durationInNs := float64(b.duration.Nanoseconds())
	fmt.Printf("%v %fns\n", label, durationInNs)
}

// Display the recorded duration in microseconds
func (b *BM) DurationInMicroseconds(label string) {
	durationInUs := float64(b.duration.Nanoseconds()) * 0.001
	fmt.Printf("%v %fµs\n", label, durationInUs)
}

// Display the recorded duration in milliseconds
func (b *BM) DurationInMilliseconds(label string) {
	durationInMs := float64(b.duration.Nanoseconds()) * 0.000001
	fmt.Printf("%v %fms\n", label, durationInMs)
}

// Display the recorded duration in seconds
func (b *BM) DurationInSeconds(label string) {
	durationInSecs := float64(b.duration.Nanoseconds()) * 0.000000001
	fmt.Printf("%v %fs\n", label, durationInSecs)
}
