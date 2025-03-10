package menu

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierDisabled
	ModiferActive
	ModifierFocus
)

func (m Modifier) Class() string {
	switch m {
	case ModifierDisabled:
		return "disabled"
	case ModiferActive:
		return "active"
	case ModifierFocus:
		return "focus"
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
	SizeXL
)

func (s Size) Class() string {
	switch s {
	case SizeXS:
		return "menu-xs"
	case SizeSM:
		return "menu-sm"
	case SizeLG:
		return "menu-lg"
	case SizeXL:
		return "menu-xl"
	default:
		return "menu-md"
	}
}

type Direction int

const (
	DirectionHorizontal Direction = iota
	DirectionVertical
)

func (d Direction) Class() string {
	switch d {
	case DirectionHorizontal:
		return "menu-horizontal"
	default:
		return "menu-vertical"
	}
}
