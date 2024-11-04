package godaisyui

import (
	"fmt"

	"github.com/nadama95/gotemplates/components"
)

type MenuParams struct {
	AsSubMenu     bool
	Classes       string
	ListItems     []components.Component
	Title         string
	TitleAsParent bool
	Size
	Direction
}

func Menu(p MenuParams) components.Component {

	cmp := components.NewComponent("ul")

	if p.AsSubMenu == false {
		cmp = cmp.AddClass(
			fmt.Sprintf("menu menu-%s menu-%s", p.Size, p.Direction),
		)
	}

	cmp = cmp.AddClass(p.Classes)

	submenu := components.NewComponent("li")
	if p.TitleAsParent {
		if p.Title != "" {
			submenu.AddChild(
				createMenuTitle(p.Title, p.TitleAsParent),
			)
		}

		submenuChildren := components.NewComponent("ul")

		cmp = cmp.AddChild(
			submenu.AddChild(
				addMenuItems(submenuChildren, p.ListItems),
			),
		)
	} else {
		if p.Title != "" {
			cmp.AddChild(createMenuTitle(p.Title, p.TitleAsParent))
		}
		cmp = cmp.AddChild(
			addMenuItems(submenu, p.ListItems),
		)
	}

	return cmp
}

func createMenuTitle(t string, asParent bool) components.Component {
	var c components.Component
	if asParent {
		c = components.NewComponent("a")
	} else {
		c = components.NewComponent("li")
	}

	return c.AddClass("menu-title").SetText(t)
}

func addMenuItems(m components.Component, is []components.Component) components.Component {
	for _, i := range is {
		m = m.AddChild(i)
	}

	return m
}
