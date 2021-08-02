package termcontrol

import "fmt"

const (
	csi = "\u001b["
)

func CursorUp(amount int)  string {
	return fmt.Sprintf("%s%dA", csi, amount)
}

func CursorDown(amount int)  string {
	return fmt.Sprintf("%s%dB", csi, amount)
}

func CursorFoward(amount int)  string {
	return fmt.Sprintf("%s%dC", csi, amount)
}

func CursorBack(amount int)  string {
	return fmt.Sprintf("%s%dD", csi, amount)
}

func CursorNextLine(amount int)  string {
	return fmt.Sprintf("%s%dE", csi, amount)
}

func CursorPrevLine(amount int)  string {
	return fmt.Sprintf("%s%dF", csi, amount)
}

func CursorHorizontalPos(column int)  string {
	return fmt.Sprintf("%s%dG", csi, column + 1)
}

func CursorPos(column, row int)  string {
	return fmt.Sprintf("%s%d;%dH", csi, row + 1, column + 1)
}

func CursorPosWithFormatEffector(column, row int)  string {
	return fmt.Sprintf("%s%d;%df", csi, row + 1, column + 1)
}

func ClearFromCursorToScreenEnd() string {
	return fmt.Sprintf("%sJ", csi)
}

func ClearFromCursorToScreenStart() string {
	return fmt.Sprintf("%s1J", csi)
}

func ClearScreen() string {
	return fmt.Sprintf("%s2J", csi)
}

func ClearScreenAndScrollback() string {
	return fmt.Sprintf("%s3J", csi)
}

func ClearFromCursorToLineEnd() string {
	return fmt.Sprintf("%sK", csi)
}

func ClearFromCursorToLineStart() string {
	return fmt.Sprintf("%s1K", csi)
}

func ClearLine() string {
	return fmt.Sprintf("%s2K", csi)
}

func ClearLineAndMoveToStart() string {
	fmt.Printf("%s2K", csi)
	return CursorHorizontalPos(0)
}

func ScrollUp(amount int) string {
	return fmt.Sprintf("%s%dS", csi, amount)
}

func ScrollDown(amount int) string {
	return fmt.Sprintf("%s%dT", csi, amount)
}

func SetGraphicRendition(rendition GraphicRendition) string {
	return fmt.Sprintf("%s%dm", csi, rendition)
}

func SetOneColor(color Color, bg, bright bool) string {
	if bg { color += 10 }
	if bright { color += 60 }
	return fmt.Sprintf("%s%dm", csi, color)
}

func SetBothColors(fgCol, bgCol Color, fgBright, bgBright bool) string {
	if fgBright { fgCol += 60 }
	if bgBright { bgCol += 60 }
	return fmt.Sprintf("%s%d;%dm", csi, fgCol, bgCol + 10)
}

func SetOneColor24Bit(r, g, b int, bg bool) string {
	esc := 38
	if bg { esc += 10 }
	return fmt.Sprintf("%s%d;2;%d;%d;%dm", csi, esc, r, g, b)
}

func SetBothColors24Bit(fgR, fgG, fgB, bgR, bgG, bgB int) string {
	return fmt.Sprintf("%s38;2;%d;%d;%d;48;2;%d;%d;%dm", csi, fgR, fgG, fgB, bgR, bgG, bgB)
}

func EnableAux() string {
	return fmt.Sprintf("%s5i", csi)
}

func DisableAux() string {
	return fmt.Sprintf("%s4i", csi)
}

func DeviceStatusReport() string {
	return fmt.Sprintf("%s6n", csi)
	//TODO: receive and parse the DSR - it sends ESC[r;cR for row and column of cursor
}

func UseAlternateFont(fontNum int) string {
	return SetGraphicRendition(GraphicRendition(11 + fontNum))
}