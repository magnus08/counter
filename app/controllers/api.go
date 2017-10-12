package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/magnus08/counter/app/event"
)

type Api struct {
	*revel.Controller
}

func (c Api) ListEvents(user int) revel.Result {
	fmt.Printf("*** ListEvents here\n")

	return c.RenderJSON(event.ListEvents(user))
}

type EventBody struct {
	Typ      string
	Name      string
	Energy      int
}

func (c Api) AddEvent(user int, eventBody EventBody) revel.Result {
	// Probably a bit questionable to use the model as body...
	fmt.Printf("*** AddEvent here %d, %+v\n", user, eventBody)
	event.AddEvent(event.NewEvent(user, eventBody.Typ, eventBody.Name, eventBody.Energy))
	return c.RenderJSON(1)
}
