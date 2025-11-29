package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// DrawFPS draws a simple FPS and frame counter in the top-right corner.
func DrawFPS(screen *ebiten.Image, frames int) {
	w, _ := screen.Size()
	text.Draw(screen, "Frames: "+itoa(frames), basicfont.Face7x13, w-120, 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
}

// tiny itoa
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	neg := false
	if v < 0 {
		neg = true
		v = -v
	}
	digits := make([]byte, 0, 16)
	for v > 0 {
		digits = append(digits, byte('0'+(v%10)))
		v /= 10
	}
	if neg {
		digits = append(digits, '-')
	}
	// reverse
	for i := len(digits)/2 - 1; i >= 0; i-- {
		opp := len(digits) - 1 - i
		digits[i], digits[opp] = digits[opp], digits[i]
	}
	return string(digits)
}
