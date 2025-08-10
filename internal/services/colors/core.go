package colors

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"github.com/alltom/oklab"
)

// HexToOklch converts a hex color string (e.g., "#RRGGBB") to its OKLCH representation.
func HexToOklch(hexColor string) (oklchColor oklab.Oklch, err error) {
	if len(hexColor) != 7 || hexColor[0] != '#' {
		return oklchColor, fmt.Errorf("invalid hex color format: expected #RRGGBB, got %s", hexColor)
	}

	r, errR := strconv.ParseUint(hexColor[1:3], 16, 8)
	g, errG := strconv.ParseUint(hexColor[3:5], 16, 8)
	b, errB := strconv.ParseUint(hexColor[5:7], 16, 8)

	if errR != nil || errG != nil || errB != nil {
		return
	}

	rgbaColor := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
	oklchColor = oklab.OklchModel.Convert(rgbaColor).(oklab.Oklch)

	return
}

// OklchToString converts an OKLCH color to a css string.
func OklchToString(oklchColor *oklab.Oklch) string {
	return fmt.Sprintf("oklch(%.2f %.3f %06.2f)", oklchColor.L, oklchColor.C, toDegree(oklchColor.H))
}

func OklchToHex(oklchColor *oklab.Oklch) string {
	if oklchColor == nil {
		return "#000000"
	}
	r32, g32, b32, _ := oklchColor.RGBA()

	r8 := uint8(r32 >> 8)
	g8 := uint8(g32 >> 8)
	b8 := uint8(b32 >> 8)

	return fmt.Sprintf("#%02x%02x%02x", r8, g8, b8)
}

func toDegree(hue float64) float64 {
	hueDegrees := hue * 180 / math.Pi
	if hueDegrees < 0 {
		hueDegrees += 360
	}
	return hueDegrees
}

// sRGBToLinear converts an sRGB color component (0.0-1.0) to its linear representation.
func sRGBToLinear(c float64) float64 {
	if c <= 0.04045 { // WCAG 2.2 uses 0.04045, slightly different from 0.03928 in some older specs
		return c / 12.92
	} else {
		return math.Pow((c+0.055)/1.055, 2.4)
	}
}

// RelativeLuminance calculates the relative luminance of a color according to WCAG standards.
func RelativeLuminance(c oklab.Oklch) float64 {
	rInt, gInt, bInt, _ := c.Oklab().RGBA()

	r := float64(rInt) / 65535.0
	g := float64(gInt) / 65535.0
	b := float64(bInt) / 65535.0

	rLin := sRGBToLinear(r)
	gLin := sRGBToLinear(g)
	bLin := sRGBToLinear(b)

	return 0.2126*rLin + 0.7152*gLin + 0.0722*bLin
}

// ContrastRatio calculates the contrast ratio between two OKLCH colors (c1, c2).
func ContrastRatio(c1, c2 oklab.Oklch) float64 {
	l1 := RelativeLuminance(c1)
	l2 := RelativeLuminance(c2)

	if l2 > l1 {
		l1, l2 = l2, l1
	}

	return (l1 + 0.05) / (l2 + 0.05)
}

func OklchCompare(c1, c2 oklab.Oklch) (rl float64, cr float64) {
	rl = RelativeLuminance(c2)
	cr = ContrastRatio(c1, c2)
	return
}
