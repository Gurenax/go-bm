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
<br>

## API
  ### struct BM
    ```go
    type BM struct {
      startTime time.Time
      endTime   time.Time
      duration  time.Duration
    }
    ```
  The struct will contain the start and end time of the benchmark. This needs to be initialised as empty.<br>
  e.g. `b := bm.BM{}`
<br><br>
  ### func Start
    ```go
    func (b *BM) Start()
    ```
  The function populates the `startTime` value.
<br><br>
  ### func End
    ```go
    func (b *BM) End()
    ```
  The function populates the `endTime` value and calculations the `duration` value.
<br><br>
  ### func DurationInNanoseconds
    ```go
    func (b *BM) DurationInNanoseconds(label string)
    ```
  The function prints out the duration in Nanoseconds.
<br><br>
  ### func DurationInMicroseconds(label string)
    ```go
    func (b *BM) DurationInMicroseconds(label string)
    ```
  The function prints out the duration in Microseconds.
<br><br>
  ### func DurationInMilliseconds(label string)
    ```go
    func (b *BM) DurationInMilliseconds(label string)
    ```
  The function prints out the duration in Milliseconds.
<br><br>
  ### func DurationInSeconds(lanel string)
    ```go
    func (b *BM) DurationInSeconds(label string)
    ```
  The function prints out the duration in Seconds.