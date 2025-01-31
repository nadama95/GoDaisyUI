package dropdown

type Placement int

const (
	PlacementStart Placement = iota
	PlacementCenter
	PlacementEnd
	PlacementTop
	PlacementBottom
	PlacementLeft
	PlacementRight
)

func (p Placement) Class() string {
	switch p {
	case PlacementStart:
		return "dropdown-start"
	case PlacementCenter:
		return "dropdown-center"
	case PlacementEnd:
		return "dropdown-end"
	case PlacementTop:
		return "dropdown-top"
	case PlacementBottom:
		return "dropdown-bottom"
	case PlacementLeft:
		return "dropdown-left"
	case PlacementRight:
		return "dropdown-right"
	}

	return ""
}

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierHover
	ModifierOpen
)

func (m Modifier) Class() string {
	switch m {
	case ModifierHover:
		return "dropdown-hover"
	case ModifierOpen:
		return "dropdown-open"
	}
	return ""
}
