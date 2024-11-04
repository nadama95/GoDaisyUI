package godaisyui

import (
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
	return components.NewComponent("ul")
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
