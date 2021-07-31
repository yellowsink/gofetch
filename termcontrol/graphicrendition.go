package termcontrol

type graphicRendition int

const (
	Reset graphicRendition = iota
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

	BlackletterFont graphicRendition = iota + 9
	DoubleUnderline
	NormalIntensity
	NeitherItalicNorBlackletter
	NotUnderlined
	NotBlinking
	ProportionalSpacing
	NotReversed
	NotConcealed
	NotStrikedThru

	DisableProportionalSpacing graphicRendition = iota + 29
	Framed
	Encircled
	Overlined
	NeitherFramedNorEncircled
	NotOverlined

	Superscript graphicRendition = iota + 46
	Subscript
	NeitherSuperscriptNorSubscript
)