package swap

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierActive
)

func (m Modifier) Class() string {
	switch m {
	case ModifierActive:
		return "swap-active"
	default:
		return ""
	}
}

type Style int

const (
	StyleDefault Style = iota
	StyleRotate
	StyleFlip
)

func (s Style) Class() string {
	switch s {
	case StyleRotate:
		return "swap-rotate"
	case StyleFlip:
		return "swap-flip"
	default:
		return ""
	}
}
