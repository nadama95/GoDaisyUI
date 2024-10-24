package main

import (
	"fmt"

	"github.com/nadama95/gotemplates/components"
)

type ButtonParams struct {
	Size    Size
	Variant Variant
}

func Button(p ButtonParams) components.Component {
	c := components.NewComponent("button").OverrideClasses(fmt.Sprintf("btn btn-%s btn-%s", p.Variant, p.Size))
	return c
}
