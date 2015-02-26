package stopwatch

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start, stop time.Time
	elapsed     time.Duration
}

func New() *Stopwatch {
	// Start the stopwatch with the time of creation as a default
	return &Stopwatch{start: time.Now()}
}

func (s *Stopwatch) String() string {
	// Calculate each component of the overall time
	ns := s.Nanoseconds()
	ms := ns / 1000000
	sec := ms / 1000
	min := sec / 60
	hr := min / 60

	// Get only the excess amt of each component
	ns %= 1000000
	ms %= 1000
	sec %= 60
	min %= 60
	hr %= 60

	// Express ns as ms to 3 significant digits
	ns /= 1000

	return fmt.Sprintf("%dh %dm %ds %d.%dms",
		hr, min, sec, ms, ns)
}

func (s *Stopwatch) Start() {
	// Manually set the start time
	s.start = time.Now()
}

func (s *Stopwatch) Stop() {
	// Manually setting the stop time
	s.stop = time.Now()
	s.elapsed = s.stop.Sub(s.start)
}

func (s *Stopwatch) Reset() {
	// Sets the elapsed time to zero and resets the start time
	s.start = time.Now()
	s.elapsed = 0
}

func (s *Stopwatch) Nanoseconds() int64 {
	// Return the elpased stopwatch time in nanoseconds
	if s.elapsed == 0 {
		// If Stop() hasn't been called yet, call it, every other method
		// for displaying elapsed time calls this one so if you try to
		// print the value or pull an elapsed time it sets a stop as default
		s.Stop()
	}
	return s.elapsed.Nanoseconds()
}

func (s *Stopwatch) Milliseconds() float64 {
	// Returns the elapsed stopwatch time in milliseconds
	return float64(s.Nanoseconds()) / 1000000
}

func (s *Stopwatch) Seconds() float64 {
	// Returns the elapsed stopwatch time in seconds
	return s.Milliseconds() / 1000
}

func (s *Stopwatch) Minutes() float64 {
	// Returns the elapsed stopwatch time in minutes
	return s.Seconds() / 60
}

func (s *Stopwatch) Hours() float64 {
	// Returns the elapsed stopwatch time in hours
	return s.Minutes() / 60
}
