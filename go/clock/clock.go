package main

import (
	"fmt"
)

const testVersion = 4

type Clock int

func new(h, m int) Clock {
	c := Clock((h*60 + m) % 1440)
	if c < 0 {
		c += 1440
	}
	fmt.Println(c)
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
