package termcontrol

type GraphicRendition int

const (
	Reset GraphicRendition = iota
	Bold
	Faint
	Italic
	Underline
	SlowBlink
	RapidBlink
	ReverseVideo
	Conceal
	Strikethru
	PrimaryFont

	BlackletterFont GraphicRendition = iota + 9
	DoubleUnderline
	NormalIntensity
	NeitherItalicNorBlackletter
	NotUnderlined
	NotBlinking
	ProportionalSpacing
	NotReversed
	NotConcealed
	NotStrikedThru

	DisableProportionalSpacing GraphicRendition = iota + 29
	Framed
	Encircled
	Overlined
	NeitherFramedNorEncircled
	NotOverlined

	Superscript GraphicRendition = iota + 46
	Subscript
	NeitherSuperscriptNorSubscript
)