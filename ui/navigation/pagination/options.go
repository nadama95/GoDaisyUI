package pagination

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
		return "btn-xs"
	case SizeSM:
		return "btn-sm"
	case SizeLG:
		return "btn-lg"
	case SizeXL:
		return "btn-xl"
	default:
		return "btn-md"
	}
}
