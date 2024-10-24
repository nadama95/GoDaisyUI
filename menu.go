package godaisyui

import (
	"fmt"

	"github.com/nadama95/gotemplates/components"
)

type MenuParams struct {
	Classes       string
	ListItems     []components.Component
	Title         string
	TitleAsParent bool
	Size
	Direction
}

func Menu(p MenuParams) components.Component {
	cmp := components.NewComponent("ul").AddClass(
		fmt.Sprintf("menu menu-%s menu-%s", p.Size, p.Direction),
	).AddClass(p.Classes)

	if p.Title != "" {
		if p.TitleAsParent {
			cmp = cmp.AddChild(
				components.NewComponent("li").AddChild(
					components.NewComponent("h2").AddClass("menu-title").SetText(p.Title),
				),
			)
		} else {
			cmp = cmp.AddChild(
				components.NewComponent("li").AddClass("menu-title").SetText(p.Title),
			)
		}
	}

	if p.TitleAsParent {
		submenu := components.NewComponent("ul")

		cmp = cmp.AddChild(
			addMenuItems(submenu, p.ListItems),
		)
	} else {
		cmp = addMenuItems(cmp, p.ListItems)
	}

	return cmp
}

func addMenuItems(m components.Component, is []components.Component) components.Component {
	for _, i := range is {
		m = m.AddChild(i)
	}

	return m
}
