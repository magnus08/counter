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
	fmt.Printf("What!\n")

	return c.RenderJSON(event.ListEvents(user))
}

func (c Api) AddEvent() revel.Result {
  fmt.Printf("What?")
  var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	event.AddEvent(event.NewEvent("aa", 1, "AA name", 100))
	return c.RenderJSON(1)
}
