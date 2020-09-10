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
  b.StartBenchmark()
  
  // Your function
  result := calculateSomething()
  
  // Ending the Benchmark
  b.EndBenchmark()
  
  // Displaying the Result
  b.DisplayResultInNanoseconds("Benchmark in Nanoseconds:") // Benchmark in Nanoseconds: 1720.000000ns
  b.DisplayResultInMicroseconds("Benchmark in Microseconds:") // Benchmark in Microseconds: 1.720000µs
  b.DisplayResultInMilliseconds("Benchmark in Milliseconds:") // Benchmark in Milliseconds: 0.001720ms
  b.DisplayResultInSeconds("Benchmark in Seconds:") // Benchmark in Seconds: 0.000002s
}
```