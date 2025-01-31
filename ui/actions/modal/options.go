package modal

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierOpen
)

func (m Modifier) Class() string {
	switch m {
	case ModifierOpen:
		return "modal-open"
	default:
		return ""
	}
}

type Placement int

const (
	PlacementTop Placement = iota
	PlacementMiddle
	PlacementBottom
	PlacementStart
	PlacementEnd
)

func (p Placement) Class() string {
	switch p {
	case PlacementTop:
		return "modal-top"
	case PlacementMiddle:
		return "modal-middle"
	case PlacementBottom:
		return "modal-bottom"
	case PlacementStart:
		return "modal-start"
	case PlacementEnd:
		return "modal-end"
	default:
		return ""
	}
}
