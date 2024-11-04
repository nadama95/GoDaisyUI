package godaisyui

import (
	"fmt"

	"github.com/nadama95/gotemplates/components"
)

type TableParams struct {
	Classes       string
	TableBodyRows []components.TD
	TableHeadRows []components.TR
	Striped       bool
	Size
}

func (p TableParams) Table() components.Component {
	div := components.NewDiv()
	div.AddClass("overflow-x-auto")

	table := components.NewComponent("table").AddClass(fmt.Sprintf("table table-%s", p.Size))

	if p.Striped {
		table.AddClass("table-zebra")
	}

	thead := components.NewComponent("thead")

	for _, r := range p.TableHeadRows {
		thead.AddChild(r)
	}

	tbody := components.NewComponent("tbody")

	for _, r := range p.TableBodyRows {
		tbody.AddChild(r)
	}

	table.AddChild(thead).AddChild(tbody)

	div.AddChild(
		table,
	)

	return div
}
