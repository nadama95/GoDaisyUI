package accordion

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierArrow
	ModifierPlus
	ModifierOpen
	ModifierClose
)

func (m Modifier) Class() string {
	switch m {
	case ModifierArrow:
		return "collapse-arrow"
	case ModifierPlus:
		return "collapse-plus"
	case ModifierOpen:
		return "collapse-open"
	case ModifierClose:
		return "collapse-close"
	default:
		return ""
	}
}
