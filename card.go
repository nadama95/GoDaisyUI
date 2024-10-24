package main

import (
	"github.com/nadama95/gotemplates/components"
)

type CardParams struct {
	BodyClass    string
	Title        string
	TitleClass   string
	Actions      []components.Component
	ActionsClass string
	Bordered     bool
	Compact      bool
}

func Card(p CardParams) components.Component {
	body := components.NewComponent("div").AddClass("card-body").AddClass(p.BodyClass)

	if p.Bordered {
		body = body.AddClass("card-bordered")
	}

	if p.Compact {
		body = body.AddClass("card-compact")
	}

	if p.Title != "" {
		body = body.AddChild(
			components.NewComponent("h2").OverrideClasses("card-title").AddClass(p.TitleClass).SetText(p.Title),
		)
	}

	if len(p.Actions) > 0 {
		actions := components.NewComponent("div").AddClass("card-actions").AddClass(p.ActionsClass)
		for _, a := range p.Actions {
			actions = actions.AddChild(a)
		}

		body.AddChild(actions)
	}

	return components.NewComponent("div").AddClass("card shadow-xl").AddChild(body)
}
