# Go Benchmark
Super simple tool for measuring run time of your go code

## Example
```go
package main

import (
	benchmark "github.com/Gurenax/go-benchmark"
)

func main() {
  // Starting the Benchmark
	b := benchmark.Benchmark{}
  b.StartBenchmark()
  
  // Your function
  result := calculateSomething()
  
  // Ending the Benchmark
  b.EndBenchmark()
  
  // Displaying the Result
	b.DisplayResultInMilliseconds("Benchmark in Milliseconds:")
	b.DisplayResultInSeconds("Benchmark in Seconds:")
}
```