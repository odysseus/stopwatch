package stopwatch

import (
	"fmt"
	"time"
)

// For the sake of simplicity this class boils down to two methods:
// Start() creates and starts a new stopwatch instance
// String() stops, resets and returns the string of the time
// The implicit stop removes the generally redundant line of sw.Stop()
// and ensures that you will always get a value
// The implicit reset allows you to place multiple print calls in a
// script to time different segments individually

// Ultimately I will probably add a 'break' feature of some kind that
// would allow you access to both the timing of the entire script and
// the timing of different segments

type Stopwatch struct {
	startt, stopt time.Time
	elapsed       time.Duration
}

func Start() *Stopwatch {
	return &Stopwatch{startt: time.Now()}
}

func (s *Stopwatch) String() string {
	// Calculate each component of the overall time
	ns := s.nanoseconds()
	ms := ns / 1000000
	sec := ms / 1000
	min := sec / 60

	// Get only the excess amt of each component
	ns %= 1000000
	ms %= 1000
	sec %= 60

	// Express ns as ms to 3 significant digits
	ns /= 1000

	s.reset()
	return fmt.Sprintf("%dm %ds %d.%dms",
		min, sec, ms, ns)
}

func (s *Stopwatch) stop() {
	// Manually setting the stop time
	s.stopt = time.Now()
	s.elapsed = s.stopt.Sub(s.startt)
}

func (s *Stopwatch) reset() {
	// Sets the elapsed time to zero and resets the start time
	s.startt = time.Now()
	s.elapsed = 0
}

func (s *Stopwatch) nanoseconds() int64 {
	// Return the elpased stopwatch time in nanoseconds
	// Every other method for diplaying the time calls this one first, so
	// call Stop() every time this is called.
	s.stop()
	return s.elapsed.Nanoseconds()
}

func (s *Stopwatch) milliseconds() float64 {
	// Returns the elapsed stopwatch time in milliseconds
	return float64(s.nanoseconds()) / 1000000
}

func (s *Stopwatch) seconds() float64 {
	// Returns the elapsed stopwatch time in seconds
	return s.milliseconds() / 1000
}

func (s *Stopwatch) minutes() float64 {
	// Returns the elapsed stopwatch time in minutes
	return s.seconds() / 60
}

func (s *Stopwatch) hours() float64 {
	// Returns the elapsed stopwatch time in hours
	return s.minutes() / 60
}
