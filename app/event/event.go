package event

import (
	"time"
)

type Event struct {
	Type      string
	User      int
	Timestamp int
	Name      string
	Kcal      int
}

var Events []Event

func newEvent(typ string, user int, name string, kcal int) Event {
	return Event{typ, user, int(time.Now().Unix()), name, kcal}
}

func listEvents() []Event {
	return Events
}

func init() {
}
