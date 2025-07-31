package server

// Index handler for the root path using Echo context.
// func Index(c echo.Context) error {
// 	cmp := views.FullPage(views.Page{
// 		Title: "Index",
// 	})

// 	return cmp.Render(c.Request().Context(), c.Response().Writer)
// }

// func Palette(c echo.Context) error {
// 	hex := c.QueryParam("hex")
// 	if hex == "" {
// 		hex = "#222536"
// 	}
// 	// hex = "#ffffff"
// 	p := palette.Generate(hex)
// 	viewModel := p.ToView()

// 	cmp := views.FullPage(views.Page{
// 		Title:   "Palette View",
// 		Content: colors.Palette(viewModel, hex),
// 	})

// 	return cmp.Render(c.Request().Context(), c.Response().Writer)
// }

// func (p Palette) ToView() []colorsPage.ColorScaleView {
// 	var views []colorsPage.ColorScaleView

// 	bgcDetails, ok := p[Base]
// 	if !ok {
// 		return views
// 	}
// 	seed := bgcDetails.Base

// 	convertScale := func(clor Color, scale *ColorDetails) colorsPage.ColorScaleView {
// 		view := colorsPage.ColorScaleView{
// 			Name:   string(scale.Color),
// 			Code:   string(clor),
// 			Value:  scale.Shades[600].Oklch,
// 			Shades: make([]colorsPage.Shade, len(shades)),
// 		}

// 		for i, shadeVal := range shades {
// 			colorDetails, ok := scale.Shades[shadeVal]
// 			if !ok {
// 				continue
// 			}

// 			color := colorDetails.Oklch
// 			rl, cr := OklchCompare(seed, color)
// 			hue := toDegree(color.H)

// 			radius := 37.0

// 			scaledRadiusC := radius * math.Tanh(6.0*color.C)
// 			totalDistanceC := scaledRadiusC

// 			scaledRadiusL := radius * math.Pow(color.L, 1.5)
// 			totalDistanceL := scaledRadiusL

// 			angle := -color.H
// 			view.Shades[i] = colorsPage.Shade{
// 				Code:  string(clor),
// 				Value: shadeVal,
// 				RL:    rl,
// 				CR:    cr,
// 				L:     fmt.Sprintf("%0.1f%%", color.L*100),
// 				C:     fmt.Sprintf("%0.3f", color.C),
// 				H:     fmt.Sprintf("%0.1f", hue),
// 				Hex:   OklchToHex(&color),
// 				Cx:    50.0 + totalDistanceC*math.Cos(angle),
// 				Cy:    50.0 + totalDistanceC*math.Sin(angle),
// 				Clx:   50.0 + totalDistanceL*math.Cos(angle),
// 				Cly:   50.0 + totalDistanceL*math.Sin(angle),
// 			}
// 		}
// 		return view
// 	}

// 	for _, code := range orderedColors {
// 		details, ok := p[code]
// 		if ok {
// 			views = append(views, convertScale(code, details))
// 		}
// 	}

// 	return views
// }
