package stat

type Direction int

const (
	DirectionHorizontal Direction = iota
	DirectionVertical
)

func (d Direction) Class() string {
	switch d {
	case DirectionVertical:
		return "stats-vertical"
	default:
		return "stats-horizontal"
	}
}
