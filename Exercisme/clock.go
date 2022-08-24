package main

import (
	"fmt"
	"strconv"
)

// // Define the Clock type here.

// type Clock struct {
// 	hours   int
// 	minutes int
// }

// func New(h, m int) Clock {

// 	if h > 24 {
// 		panic("No such hours!")
// 	}
// 	return Clock{
// 		hours:   h,
// 		minutes: m,
// 	}
// }

// func (c Clock) Add(m int) Clock {

// 	newMinutes := c.minutes + m
// 	if newMinutes >= 60 {
// 		c.hours = c.hours + newMinutes/60
// 		c.minutes = newMinutes % 60
// 	} else {
// 		c.minutes = newMinutes
// 	}

// 	return Clock{
// 		hours:   c.hours,
// 		minutes: c.minutes,
// 	}
// }

// func (c Clock) Subtract(m int) Clock {
// 	newMinutes := c.minutes - m
// 	if newMinutes < 0 {
// 		// 67 / 2
// 		var newHours float64 = float64(-1 * newMinutes / 60)
// 		newHours = math.Ceil(float64(newHours))
// 		c.hours = c.hours - int(newHours)
// 		c.minutes = 60 - (-1 * newMinutes % 60)
// 	} else {
// 		c.minutes = newMinutes
// 	}

// 	return Clock{
// 		hours:   c.hours,
// 		minutes: c.minutes,
// 	}
// }

// func (c Clock) String() string {
// 	// fmt.Println(c.hours, c.minutes)
// 	return fmt.Sprintf("%v:%v ", c.hours, c.minutes)

// }
type Clock struct {
	hours, minutes int
}

func fixTime(hours, minutes int) (int, int) {
	tMinutes := minutes + hours*60
	// Fix negative amount of tMinutes
	if tMinutes < 0 {
		for tMinutes < 0 {
			tMinutes = 24*60 + tMinutes
		}
	}
	Hours := tMinutes / 60
	Minutes := tMinutes % 60
	// Deal with Hours > 24, we dont manage dates
	if Hours > 24 {
		n := Hours
		for n > 24 {
			Hours = n % 24
			n = n / 24
		}
	}
	// Fix 24 as 00 in Hours
	if Hours == 24 {
		Hours = 0
	}
	return Hours, Minutes
}

// String method for Clock
func (t Clock) String() string {
	var h, m string
	hours := t.hours
	minutes := t.minutes
	if hours < 10 {
		h = "0" + strconv.Itoa(hours)
	} else {
		h = strconv.Itoa(hours)
	}
	if minutes < 10 {
		m = "0" + strconv.Itoa(minutes)
	} else {
		m = strconv.Itoa(minutes)
	}
	return h + ":" + m
}

// Add method for Clock
func (t Clock) Add(minutes int) Clock {
	return New(t.hours, t.minutes+minutes)
}

// Subtract minutes for Clock
func (t Clock) Subtract(a int) Clock {
	return New(t.hours, t.minutes-a)
}

// New creates a clock based on two integers
func New(hours, minutes int) Clock {
	fHours, fMinutes := fixTime(hours, minutes)
	clock := Clock{fHours, fMinutes}
	return clock
}

func main() {

	oneClock := New(10, 20)
	twoClock := New(10, 3)
	// twoClock = twoClock.Add(31)
	twoClock = twoClock.Subtract(70)
	fmt.Println(oneClock.String(), twoClock.String())

}
