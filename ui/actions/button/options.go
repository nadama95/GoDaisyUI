package button

type Color int

const (
	ColorPrimary Color = iota
	ColorSecondary
	ColorAccent
	ColorInfo
	ColorSuccess
	ColorWarning
	ColorError
)

func (c Color) Class() string {
	switch c {
	case ColorPrimary:
		return "btn-primary"
	case ColorSecondary:
		return "btn-secondary"
	case ColorAccent:
		return "btn-accent"
	case ColorInfo:
		return "btn-info"
	case ColorSuccess:
		return "btn-success"
	case ColorWarning:
		return "btn-warning"
	case ColorError:
		return "btn-error"
	}
	return ""
}

type Style int

const (
	StyleDefault Style = iota
	StyleOutline
	StyleSoft
	StyleGhost
	StyleLink
)

func (s Style) Class() string {
	switch s {
	case StyleOutline:
		return "btn-outline"
	case StyleSoft:
		return "btn-soft"
	case StyleGhost:
		return "btn-ghost"
	case StyleLink:
		return "btn-link"
	}
	return ""
}

type Behaviour int

const (
	BehaviourDefault Behaviour = iota
	BehaviourActive
	BehaviourDisabled
)

func (b Behaviour) Class() string {
	switch b {
	case BehaviourActive:
		return "btn-active"
	case BehaviourDisabled:
		return "btn-disabled"
	}
	return ""
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
		return "btn-xs"
	case SizeSM:
		return "btn-sm"
	case SizeMD:
		return "btn-md"
	case SizeLG:
		return "btn-lg"
	case SizeXL:
		return "btn-xl"
	}

	return ""
}

type Modifier int

const (
	ModifierDefault Modifier = iota
	ModifierWide
	ModifierBlock
	ModifierSquare
	ModifierCircle
)

func (m Modifier) Class() string {
	switch m {
	case ModifierWide:
		return "btn-wide"
	case ModifierBlock:
		return "btn-block"
	case ModifierSquare:
		return "btn-square"
	case ModifierCircle:
		return "btn-circle"
	}
	return ""
}
