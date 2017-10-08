package controllers

import (
	"github.com/revel/revel"
)

type Test struct {
	*revel.Controller
}

func (c Test) Test() revel.Result {
	return c.RenderJSON(122)
}
