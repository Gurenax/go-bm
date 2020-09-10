# Go BM (Benchmark)
Super simple tool for measuring run time of your go code

## Example
```go
package main

import (
  bm "github.com/Gurenax/go-bm"
)

func main() {
  // Starting the Benchmark
  b := bm.BM{}
  b.Start()
  
  // Your function
  result := calculateSomething()
  
  // Ending the Benchmark
  b.End()
  
  // Displaying the Result
  b.DurationInNanoseconds("Duration in Nanoseconds:") // Duration in Nanoseconds: 1720.000000ns
  b.DurationInMicroseconds("Duration in Microseconds:") // Duration in Microseconds: 1.720000µs
  b.DurationInMilliseconds("Duration in Milliseconds:") // Duration in Milliseconds: 0.001720ms
  b.DurationInSeconds("Duration in Seconds:") // Duration in Seconds: 0.000002s
}
```