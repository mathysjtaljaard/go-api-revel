package controllers

import (
	"github.com/revel/revel"
)

type V1 struct {
	*revel.Controller
}

func (c V1) Index() revel.Result {
	return c.Render()
}

func (c V1) CustomIndex() revel.Result {
	customSaying := "This is a custom Value injected into the render template {{.customSaying}}"
	return c.Render(customSaying)
}
