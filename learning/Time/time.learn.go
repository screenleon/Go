package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p("Year:", then.Year())
	p("Month:", then.Month())
	p("Day:", then.Day())
	p("Hour:", then.Hour())
	p("Minute:", then.Minute())
	p("Second:", then.Second())
	p("Nanosecond", then.Nanosecond())
	p("Location:", then.Location())

	p("Weekday:", then.Weekday())

	p("Before:", then.Before(now))
	p("After:", then.After(now))
	p("Equal:", then.Equal(now))

	// get time difference
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
