package godaisyui

type Size int

const (
	MD Size = iota
	LG
	SM
	XS
)

var sizeName = map[Size]string{
	MD: "md",
	LG: "lg",
	SM: "sm",
	XS: "xs",
}

func (s Size) String() string {
	return sizeName[s]
}

type Variant int

const (
	Neutral Variant = iota
	Primary
	Secondary
	Accent
	Info
	Success
	Warning
	Error
	Ghost
	Link
	Outlink
	Active
	Disabled
)

var variantName = map[Variant]string{
	Neutral:   "neutral",
	Primary:   "primary",
	Secondary: "secondary",
	Accent:    "accent",
	Info:      "info",
	Success:   "success",
	Warning:   "warning",
	Error:     "error",
	Ghost:     "ghost",
	Link:      "link",
	Outlink:   "outlink",
	Active:    "active",
	Disabled:  "disabled",
}

func (v Variant) String() string {
	return variantName[v]
}
