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
	NewEvent("aa", 1, "AA name", 100),
	NewEvent("ab", 1, "AB name", 99),
	NewEvent("ac", 2, "AC name", 98),
	NewEvent("b", 1, "B name", 200) }

	func NewEvent(typ string, user int, name string, energy int) Event {
		return Event{typ, user, int(time.Now().Unix()), name, energy}
	}

	func ListEvents(user int) (ret []Event){
		for _, v := range events {
			if(v.User == user) {
				ret = append(ret, v)
			}
		}
		return ret
	}

	func AddEvent(event Event) {
		events = append(events, event)
	}

	func init() {
	}
