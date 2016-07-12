package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	h = (h + (m / 60)) % 24
	m = m % 60
	if m < 0 {
		m += 60
		h--
	}
	if h < 0 {
		h += 24
	}
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, (c.m + m))
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
