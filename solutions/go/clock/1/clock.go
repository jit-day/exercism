package clock

import "fmt"

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	clock := Clock{hour: 0, minute: 0}
	return clock.Add(hour*60 + minute)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	hours := minutes / 60
	minutes = minutes % 60

	newHour := (clock.hour + hours) % 24
	if newHour < 0 {
		newHour += 24
	}
	newMinutes := clock.minute + minutes

	if newMinutes >= 60 {
		newHour = (newHour + 1) % 24
		newMinutes = newMinutes % 60
	}
	if newMinutes < 0 {
		newHour = (newHour - 1) % 24
		if newHour < 0 {
			newHour += 24
		}
		newMinutes = newMinutes + 60
	}

	return Clock{hour: newHour, minute: newMinutes}
}

func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}
