package clock

import (
	"fmt"
)

const testVersion = 4

type Clock int

func New(h, m int) Clock {
	mins := (h*60 + m) % 1440
	if mins < 0 {
		mins += 1440
	}
	return Clock(mins)
}

func (c Clock) Add(m int) Clock {
	return New(getHours(c), getMins(c)+m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", getHours(c), getMins(c))
}

func getHours(c Clock) int { return int(c) / 60 }
func getMins(c Clock) int  { return int(c) % 60 }
