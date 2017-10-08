package controllers

import (
	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

func (c Api) ListEvents(user int) revel.Result {
	return c.RenderJSON(user)
}
