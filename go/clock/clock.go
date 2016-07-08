package clock

const testVersion = 4

type Clock

func New(hour, minute int) Clock {
}

func (Clock) String() string {
}

func (Clock) Add(minutes int) Clock {
}
