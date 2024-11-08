package godaisyui

import (
	"fmt"
	"html/template"

	"github.com/nadama95/gotemplates/components"
)

type StatParams struct {
	Title       string
	Value       string
	Classes     string
	Description string
	Figure      template.HTML
	Actions     []components.Component
}

func Stats(stats []StatParams, direction Direction) components.Component {
	cmp := components.NewComponent("div").AddClass("stats").AddClass(fmt.Sprintf("stats-%s", direction))

	for _, s := range stats {
		cmp = cmp.AddChild(Stat(s))
	}

	return cmp
}

func Stat(s StatParams) components.Component {
	cmp := components.NewComponent("div").AddClass("stat").AddClass(s.Classes)

	if s.Title != "" {
		cmp = cmp.AddChild(
			components.NewComponent("div").AddClass("stat-title").SetText(s.Title),
		)
	}

	if s.Value != "" {
		cmp = cmp.AddChild(
			components.NewComponent("div").AddClass("stat-value").SetText(s.Value),
		)
	}

	if s.Description != "" {
		cmp = cmp.AddChild(
			components.NewComponent("div").AddClass("stat-desc").SetText(s.Description),
		)
	}

	if s.Figure != "" {
		cmp = cmp.SetText(string(s.Figure))
	}

	if len(s.Actions) > 0 {
		for _, a := range s.Actions {
			cmp = cmp.AddChild(a)
		}
	}

	return cmp
}
