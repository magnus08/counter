package event

import (
	"time"
)

type Event struct {
	Type      string
	User      int
	Timestamp int
	Name      string
	Energy      int
}

var events = []Event{
	newEvent("aa", 1, "AA name", 100),
	newEvent("ab", 1, "AB name", 99),
	newEvent("ac", 1, "AC name", 98),
	newEvent("b", 1, "B name", 200) }

func newEvent(typ string, user int, name string, energy int) Event {
	return Event{typ, user, int(time.Now().Unix()), name, energy}
}

func ListEvents(user int) []Event {
	return events
}

func init() {
}
