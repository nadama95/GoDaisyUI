package table

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierZebra
	ModifierPinRows
	ModifierPinCols
)

func (m Modifier) Class() string {
	switch m {
	case ModifierZebra:
		return "table-zebra"
	case ModifierPinRows:
		return "table-pin-rows"
	case ModifierPinCols:
		return "table-pin-cols"
	default:
		return ""
	}
}

type Size int

const (
	SizeXS Size = iota
	SizeSM
	SizeMD
	SizeLG
)

func (s Size) Class() string {
	switch s {
	case SizeXS:
		return "table-xs"
	case SizeSM:
		return "table-sm"
	case SizeLG:
		return "table-lg"
	default:
		return "table-md"
	}
}
