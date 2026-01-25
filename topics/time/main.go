package main

import (
	"context"
	"fmt"
	"time"
)

// A compact, runnable tutorial demonstrating common usages of the
// Go `time` package. Run with `go run main.go`.

func main() {
	fmt.Println("Go time package tutorial — quick examples")
	fmt.Println("-------------------------------------------")

	nowExample()
	fmt.Println()

	formatParseExample()
	fmt.Println()

	durationArithmetic()
	fmt.Println()

	timersAndTickers()
	fmt.Println()

	timezonesAndLocations()
	fmt.Println()

	monotonicExample()
}

func nowExample() {
	fmt.Println("1) Now and basic components")
	t := time.Now()
	fmt.Println("Now:", t)
	fmt.Println("Year/Month/Day:", t.Year(), t.Month(), t.Day())
	fmt.Println("Hour/Min/Sec:", t.Hour(), t.Minute(), t.Second())
}

func formatParseExample() {
	fmt.Println("2) Formatting and parsing times")
	t := time.Date(2023, 3, 18, 15, 30, 45, 0, time.UTC)

	// Go uses a reference time to define formats:
	// Mon Jan 2 15:04:05 MST 2006
	rfc := t.Format(time.RFC3339)
	custom := t.Format("2006-01-02 15:04:05")
	fmt.Println("RFC3339:", rfc)
	fmt.Println("Custom:", custom)

	// Parsing from text
	parsed, err := time.Parse(time.RFC3339, rfc)
	if err != nil {
		fmt.Println("parse error:", err)
	} else {
		fmt.Println("Parsed equals original?", parsed.Equal(t))
	}
}

func durationArithmetic() {
	fmt.Println("3) Durations")
	d1 := 2 * time.Hour
	d2, _ := time.ParseDuration("30m")
	sum := d1 + d2
	fmt.Println("2h + 30m =", sum)

	// ParseDuration accepts units: "300ms", "-1.5h" or "2h45m" (composed)
	d3, _ := time.ParseDuration("150ms")
	fmt.Println("150ms in nanoseconds:", d3.Nanoseconds())
}

func timersAndTickers() {
	fmt.Println("4) Timers and tickers (non-blocking demos)")

	// time.After returns <-chan Time that receives once after duration
	after := time.After(200 * time.Millisecond)

	// time.NewTicker provides a channel that ticks periodically
	ticker := time.NewTicker(150 * time.Millisecond)
	done := make(chan struct{})

	// Stop ticker after 4 ticks
	go func() {
		count := 0
		for t := range ticker.C {
			fmt.Println(" tick at", t.Format("15:04:05.000"))
			count++
			if count >= 4 {
				ticker.Stop()
				close(done)
				return
			}
		}
	}()

	// Show how time.After can be used in select
	select {
	case t := <-after:
		fmt.Println("after fired at", t.Format("15:04:05.000"))
	case <-done:
		fmt.Println("ticker finished before after")
	}

	// Example of NewTimer (simple safe demo)
	timer := time.NewTimer(100 * time.Millisecond)
	<-timer.C
	fmt.Println("NewTimer fired")

	// If you need to reuse a timer, it's safer to create a new one
	// or follow the documented Stop/Drain pattern. Here we simply
	// create a second timer for clarity.
	timer2 := time.NewTimer(50 * time.Millisecond)
	<-timer2.C
	fmt.Println("Second timer fired")
}

func timezonesAndLocations() {
	fmt.Println("5) Time zones and locations")

	// Local time vs UTC
	t := time.Now()
	fmt.Println("Local:", t)
	fmt.Println("UTC:  ", t.UTC())

	// LoadLocation (IANA names). On macOS the zone database is available.
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("LoadLocation error:", err)
		return
	}
	ny := time.Now().In(loc)
	fmt.Println("New York time:", ny.Format(time.RFC1123))

	// Parse in location
	ts := "2023-03-18 08:00:00"
	parsed, err := time.ParseInLocation("2006-01-02 15:04:05", ts, loc)
	if err != nil {
		fmt.Println("parse in location error:", err)
	} else {
		fmt.Println("Parsed in New York loc:", parsed)
	}
}

func monotonicExample() {
	fmt.Println("6) Measuring elapsed time (use monotonic clock)")
	start := time.Now()
	// Simulate work
	time.Sleep(120 * time.Millisecond)
	elapsed := time.Since(start) // preserves monotonic duration
	fmt.Println("Elapsed:", elapsed)

	// With context and timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(60 * time.Millisecond):
		fmt.Println("work finished")
	case <-ctx.Done():
		fmt.Println("context timed out:", ctx.Err())
	}
}
