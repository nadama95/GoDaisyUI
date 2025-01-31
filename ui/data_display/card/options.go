package card

type Style int

const (
	StyleDefault Style = iota
	StyleBorder
	StyleDash
)

func (s Style) Class() string {
	switch s {
	case StyleBorder:
		return "card-border"
	case StyleDash:
		return "card-dash"
	default:
		return ""
	}
}

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierSide
	ModifierImageFill
)

func (m Modifier) Class() string {
	switch m {
	case ModifierSide:
		return "card-side"
	case ModifierImageFill:
		return "image-full"
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
		return "card-xs"
	case SizeSM:
		return "card-sm"
	case SizeMD:
		return "card-md"
	case SizeLG:
		return "card-lg"
	case SizeXL:
		return "card-xl"
	default:
		return ""
	}
}
