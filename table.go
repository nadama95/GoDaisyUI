package godaisyui

import (
	"fmt"

	"github.com/nadama95/gotemplates/components"
)

type TableParams struct {
	Classes       string
	TableBodyRows []components.Component
	TableHeadRows []components.Component
	Striped       bool
	Size
}

func Table(p TableParams) components.Component {
	div := components.NewDiv()
	div = div.AddClass("overflow-x-auto")

	table := components.NewComponent("table").AddClass(fmt.Sprintf("table table-%s", p.Size))

	if p.Striped {
		table = table.AddClass("table-zebra")
	}

	thead := components.NewComponent("thead")

	for _, r := range p.TableHeadRows {
		thead = thead.AddChild(r)
	}

	tbody := components.NewComponent("tbody")

	for _, r := range p.TableBodyRows {
		tbody = tbody.AddChild(r)
	}

	table = table.AddChild(thead).AddChild(tbody)

	return div.AddChild(
		table,
	)
}
