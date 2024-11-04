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
		// Menu Classes
		fmt.Sprintf("menu menu-%s menu-%s", p.Size, p.Direction),
	).AddClass(p.Classes)

	if p.Title != "" {
		if p.TitleAsParent {
			section := components.NewComponent("li")
			section = section.AddChild(createMenuTitle(p.Title, p.TitleAsParent))

			return cmp.AddChild(addMenuItems(section, p.ListItems))

		} else {
			cmp = cmp.AddChild(createMenuTitle(p.Title, p.TitleAsParent))

		}
	}
	return cmp.AddChild(addMenuItems(cmp, p.ListItems))
}

type SubMenuParams struct {
	Classes       string
	ListItems     []components.Component
	Title         string
	TitleAsParent bool
}

func SubMenu(p SubMenuParams) components.Component {
	cmp := components.NewComponent("li")

	if p.Title != "" {
		cmp = cmp.AddChild(createMenuTitle(p.Title, p.TitleAsParent))
	}

	menuItems := components.NewComponent("ul")

	return cmp.AddChild(addMenuItems(menuItems, p.ListItems))
}

func createMenuTitle(t string, asParent bool) components.Component {
	var c components.Component
	if asParent {
		c = components.NewComponent("h2")
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
