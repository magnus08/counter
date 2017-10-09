package controllers

import (
	"github.com/revel/revel"
	"github.com/magnus08/counter/app/event"
)

type Api struct {
	*revel.Controller
}

func (c Api) ListEvents(user int) revel.Result {

	return c.RenderJSON(event.ListEvents(user))
}
