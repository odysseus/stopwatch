# Stopwatch
Simple stopwatch class for timing whatever your heart desires.

### Examples
```go
package main

import (
	"fmt"
	"github.com/odysseus/stopwatch"
	"time"
)

func main() {
	// New creates a stopwatch and starts the timer
	sw := stopwatch.New()

	time.Sleep(500 * time.Millisecond)

	// If Stop() has not been called manually, it is called as soon as
	// you try to use a method that needs an elapsed time
	fmt.Println(sw)
	// 0h 0m 0s 502.435ms

	// Reset start/stop and implicit stop
	sw.Reset()
	fmt.Println(sw)
	// 0h 0m 0s 0.0ms

	// Start manually sets a start time
	sw.Start()
	time.Sleep(100 * time.Millisecond)
	// And stop manually stops it
	sw.Stop()
	time.Sleep(200 * time.Millisecond)
	fmt.Println(sw)
	// 0h 0m 0s 103.730ms

	// Finally, the time can be expressed in a number of different units
	ns := sw.Nanoseconds()
	ms := sw.Milliseconds()
	s := sw.Seconds()
	m := sw.Minutes()
	h := sw.Hours()
}
```
