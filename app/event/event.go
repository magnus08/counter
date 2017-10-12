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
	NewEvent(1, "aa", "AA name", 100),
	NewEvent(1, "ab", "AB name", 99),
	NewEvent(2, "ac", "AC name", 98),
	NewEvent(1, "b", "B name", 200) }

	func NewEvent(user int, typ string, name string, energy int) Event {
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
