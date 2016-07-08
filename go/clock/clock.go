package clock

import (
	"fmt"
)

const testVersion = 4

// define the Clock type as a total number of minutes
type Clock int

func New(h, m int) Clock {
	// calculate total number of minutes from given hours & minutes
	// ensure this total is within a single 24 hour (1440 minute)
	// period by taking only the number of minutes left over after
	// removing any multiples of 1440
	c := Clock((h*60 + m) % 1440)
	// if the calculated time is negative add a days worth of minutes
	// so that it is moving forward from 0 rather than backwards
	if c < 0 {
		c += 1440
	}
	return c
}

func (c Clock) Add(m int) Clock {
	// convert the clock type to an int and add the given minutes to it
	return New(0, int(c)+m)
}

func (c Clock) String() string {
	// pull out the hours by dividing by 60 mins
	// pull out the minutes by calculating what's left
	// pretty print the results
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
