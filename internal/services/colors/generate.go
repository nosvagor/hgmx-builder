package colors

import (
	"math"

	"github.com/alltom/oklab"
)

func (c *ColorDetails) generateBg() {
	c.Shades[600] = Details{Oklch: c.Base}
	hue := c.Base.H

	// --- Light Shades (50-600) ---
	targetL50 := min(c.Base.L*1.8, 1.0)
	targetC50 := c.Base.C * 3.0
	const lightPower = 1.3

	numIntervalsLight := float64(shadesMap[600] - shadesMap[50])
	for shadeValue, index := range shadesMap {
		if shadeValue >= 600 {
			continue
		}
		t_light := 1.0 - (float64(index) / numIntervalsLight)
		t_light_curved := math.Pow(t_light, lightPower)

		l := c.Base.L + (targetL50-c.Base.L)*t_light_curved
		chroma := c.Base.C + (targetC50-c.Base.C)*t_light

		detail := c.Shades[shadeValue]
		detail.Oklch = oklab.Oklch{L: max(0, min(l, 1)), C: max(chroma, 0), H: hue}
		c.Shades[shadeValue] = detail
	}

	// --- Darker Shades (600-950) ---
	targetC950 := c.Base.C * 0.667
	targetL950 := c.Base.L * 0.667
	const darkPower = 1.5

	numIntervalsDark := float64(shadesMap[950] - shadesMap[600])
	baseIndexDark := float64(shadesMap[600])
	for shadeValue, index := range shadesMap {
		if shadeValue < 700 {
			continue
		}
		t_dark := (float64(index) - baseIndexDark) / numIntervalsDark
		t_dark_curved := math.Pow(t_dark, darkPower)

		l := c.Base.L + (targetL950-c.Base.L)*t_dark_curved
		chroma := c.Base.C + (targetC950-c.Base.C)*t_dark

		detail := c.Shades[shadeValue]
		detail.Oklch = oklab.Oklch{L: max(0, min(l, 1)), C: max(chroma, 0), H: hue}
		c.Shades[shadeValue] = detail
	}
}

func (c *ColorDetails) generateFg(bgc50 oklab.Oklch) {
	baseFg := c.Base
	if baseFg.L <= 0.5 {
		baseFg = oklab.Oklch{L: 0.85, C: baseFg.C, H: baseFg.H}
	} else {
		baseFg = oklab.Oklch{L: 0.25, C: baseFg.C, H: baseFg.H}
	}
	c.Shades[400] = Details{Oklch: baseFg}
	hue := baseFg.H

	// --- Generate Lighter Shades (50-300) ---
	targetL50 := min(baseFg.L*1.25, 0.98)
	targetC50 := baseFg.C * 1.5
	lightPower := 1.5

	numIntervalsLight := float64(shadesMap[400] - shadesMap[50])
	baseIndexLight := float64(shadesMap[400])
	for shadeValue, index := range shadesMap {
		if shadeValue >= 400 {
			continue
		}
		t_light := (baseIndexLight - float64(index)) / numIntervalsLight
		t_light_curved := math.Pow(t_light, lightPower)

		l := baseFg.L + (targetL50-baseFg.L)*t_light_curved
		chroma := baseFg.C + (targetC50-baseFg.C)*t_light

		detail := c.Shades[shadeValue]
		detail.Oklch = oklab.Oklch{L: max(0, min(l, 1)), C: max(chroma, 0), H: hue}
		c.Shades[shadeValue] = detail
	}

	// --- Generate Darker Shades (500-950) ---
	targetL950 := bgc50.L
	targetC950 := bgc50.C
	darkPower := 1.2

	numIntervalsDark := float64(shadesMap[950] - shadesMap[400])
	baseIndexDark := float64(shadesMap[400])

	index900 := float64(shadesMap[900])
	t_dark_900 := (index900 - baseIndexDark) / numIntervalsDark
	t_dark_curved_900 := math.Pow(t_dark_900, darkPower)

	calculatedL900 := baseFg.L + (targetL950-baseFg.L)*t_dark_curved_900
	calculatedC900 := baseFg.C + (targetC950-baseFg.C)*t_dark_900

	newTargetL950 := calculatedL900
	newTargetC950 := calculatedC900

	for shadeValue, index := range shadesMap {
		if shadeValue <= 400 {
			continue
		}
		t_dark := (float64(index) - baseIndexDark) / numIntervalsDark
		t_dark_curved := math.Pow(t_dark, darkPower)

		l := baseFg.L + (newTargetL950-baseFg.L)*t_dark_curved
		chroma := baseFg.C + (newTargetC950-baseFg.C)*t_dark

		detail := c.Shades[shadeValue]
		detail.Oklch = oklab.Oklch{L: max(0, min(l, 1)), C: max(chroma, 0), H: hue}
		c.Shades[shadeValue] = detail
	}
}

func (c *ColorDetails) generateColor() {
	c.Shades[600] = Details{Oklch: c.Base}
	hue := c.Base.H

	// --- Constants for Light Shades (50-500) ---
	const targetL50 = 0.97
	const targetC50 = 0.01
	const adjustL50 = 0.25
	const adjustC50 = 0.37
	numIntervalsLight := float64(shadesMap[600] - shadesMap[50])
	for shadeValue, index := range shadesMap {
		if shadeValue > 500 {
			continue
		}
		t := float64(index) / numIntervalsLight
		l := targetL50 + (c.Base.L-targetL50)*t
		chroma := max(targetC50+(c.Base.C-targetC50)*t, 0)
		detail := c.Shades[shadeValue]
		detail.Oklch = oklab.Oklch{L: max(0, min(l, 1)), C: chroma, H: hue}
		c.Shades[shadeValue] = detail
	}

	details100, ok100 := c.Shades[100]
	details50, ok50 := c.Shades[50]
	if ok100 && ok50 {
		l100_linear := details100.L
		l50_initial_linear := details50.L
		l50_adjusted := l50_initial_linear + (l100_linear-l50_initial_linear)*adjustL50
		details50.L = max(0, min(l50_adjusted, 1))

		c100_linear := details100.C
		c50_initial_linear := details50.C
		c50_adjusted := c50_initial_linear + (c100_linear-c50_initial_linear)*adjustC50
		details50.C = max(c50_adjusted, 0)
		c.Shades[50] = details50
	}

	// --- Constants for Dark Shades (700-950) ---
	const targetL950 = 0.25
	const targetC950 = 0.05
	const adjustL950 = 0.37
	const adjustC950 = 0.42
	numIntervalsDark := float64(shadesMap[950] - shadesMap[600])
	baseIndexDark := float64(shadesMap[600])
	for shadeValue, index := range shadesMap {
		if shadeValue < 700 {
			continue
		}
		t := (float64(index) - baseIndexDark) / numIntervalsDark
		l := c.Base.L + (targetL950-c.Base.L)*t
		chroma := c.Base.C + (targetC950-c.Base.C)*t
		detail := c.Shades[shadeValue]
		detail.Oklch = oklab.Oklch{L: max(0, min(l, 1)), C: max(chroma, 0), H: hue}
		c.Shades[shadeValue] = detail
	}

	details900, ok900 := c.Shades[900]
	details950, ok950 := c.Shades[950]
	if ok900 && ok950 {
		l900 := details900.L
		l950_initial := details950.L
		l950_adjusted := l950_initial + (l900-l950_initial)*adjustL950
		details950.L = max(0, min(l950_adjusted, 1))

		c900 := details900.C
		c950_initial := details950.C
		c950_adjusted := c950_initial + (c900-c950_initial)*adjustC950
		details950.C = max(c950_adjusted, 0)
		c.Shades[950] = details950
	}
}

func (c *ColorDetails) generateBW(color Color) {
	baseL := c.Base.L
	const fixedChroma = 0.0
	const fixedHue = 0.0

	var baseShadeValue int
	var lightestLTarget, darkestLTarget float64

	switch color {
	case White:
		baseShadeValue = 500
		lightestLTarget = 1.0
		darkestLTarget = 0.75
	case Black:
		baseShadeValue = 600
		lightestLTarget = 0.25
		darkestLTarget = 0.04
	default:
		return
	}

	c.Shades[baseShadeValue] = Details{Oklch: oklab.Oklch{L: baseL, C: fixedChroma, H: fixedHue}}

	indexBase := float64(shadesMap[baseShadeValue])
	indexMin := float64(shadesMap[50])
	indexMax := float64(shadesMap[950])

	for _, shadeVal := range shades {
		if shadeVal == baseShadeValue {
			continue
		}

		currentIndex := float64(shadesMap[shadeVal])
		var l float64
		var t float64

		if shadeVal < baseShadeValue {
			t = (indexBase - currentIndex) / (indexBase - indexMin)
			l = baseL + (lightestLTarget-baseL)*t
		} else {
			t = (currentIndex - indexBase) / (indexMax - indexBase)
			l = baseL + (darkestLTarget-baseL)*t
		}
		c.Shades[shadeVal] = Details{Oklch: oklab.Oklch{L: max(0.0, min(l, 1.0)), C: fixedChroma, H: fixedHue}}
	}
}

func (c *ColorDetails) generateGrey() {
	c.Shades[500] = Details{Oklch: c.Base}
	hue := c.Base.H
	baseOklch := c.Base
	const targetLLight = 0.98
	const targetCLight = 0.002
	const targetLDark = 0.20
	const targetCDark = 0.005
	for _, shadeValue := range shades {
		if shadeValue == 500 {
			continue
		} else if shadeValue < 500 {
			t := float64(500-shadeValue) / float64(500-50)
			l := targetLLight*t + baseOklch.L*(1-t)
			chroma := targetCLight*t + baseOklch.C*(1-t)
			detail := Details{Oklch: oklab.Oklch{L: max(0, min(l, 1)), C: max(chroma, 0), H: hue}}
			c.Shades[shadeValue] = detail
		} else {
			t := float64(shadeValue-500) / float64(950-500)
			l := baseOklch.L*(1-t) + targetLDark*t
			chroma := baseOklch.C*(1-t) + targetCDark*t
			detail := Details{Oklch: oklab.Oklch{L: max(0, min(l, 1)), C: max(chroma, 0), H: hue}}
			c.Shades[shadeValue] = detail
		}
	}
}
