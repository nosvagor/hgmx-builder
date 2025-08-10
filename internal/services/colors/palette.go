package colors

import (
	"fmt"
	"io"
	"log"
	"math"

	"github.com/alltom/oklab"
	"github.com/nosvagor/hgmx-builder/views/pages/palette"
)

// === Models ==================================================================

type Hex string

func (h Hex) toOklch() *oklab.Oklch {
	c, err := HexToOklch(string(h))
	if err != nil {
		panic(err)
	}
	return &c
}

type Color string

const (
	Base     Color = "base"
	Surface  Color = "surface"
	Rose     Color = "rose"
	Berry    Color = "berry"
	Cherry   Color = "cherry"
	Ruby     Color = "ruby"
	Red      Color = "red"
	Coral    Color = "coral"
	Pumpkin  Color = "pumpkin"
	Orange   Color = "orange"
	Sun      Color = "sun"
	Gold     Color = "gold"
	Honey    Color = "honey"
	Yellow   Color = "yellow"
	Lemon    Color = "lemon"
	Acid     Color = "acid"
	Lime     Color = "lime"
	Spring   Color = "spring"
	Green    Color = "green"
	Emerald  Color = "emerald"
	Jade     Color = "jade"
	Forest   Color = "forest"
	Leaf     Color = "leaf"
	Teal     Color = "teal"
	Cyan     Color = "cyan"
	Aqua     Color = "aqua"
	Robin    Color = "robin"
	Azure    Color = "azure"
	Sky      Color = "sky"
	Blue     Color = "blue"
	Cobalt   Color = "cobalt"
	Sapphire Color = "sapphire"
	Indigo   Color = "indigo"
	Lavender Color = "lavender"
	Purple   Color = "purple"
	Violet   Color = "violet"
	Pink     Color = "pink"
	Magenta  Color = "magenta"
	Brick    Color = "brick"
	Rust     Color = "rust"
	Beige    Color = "beige"
	Olive    Color = "olive"
	Moss     Color = "moss"
	Zinc     Color = "zinc"
	Gray     Color = "gray"
	Slate    Color = "slate"
	Stone    Color = "stone"
	Ash      Color = "ash"
	White    Color = "white"
	Black    Color = "black"
)

type Colors map[Color]Hex

type Details struct {
	oklab.Oklch
	RL float64
	CR float64
}

type ColorDetails struct {
	Color  Color
	Base   oklab.Oklch
	Shades map[int]Details
}

type Palette map[Color]*ColorDetails

type Motif string

const (
	// Signaling
	Info    Motif = "info"
	Success Motif = "success"
	Warning Motif = "warning"
	Error   Motif = "error"

	// Branding
	Primary Motif = "primary"
	Subtle  Motif = "secondary"
	Accent  Motif = "accent"

	// Dimensions
	Positive Motif = "positive"
	Negative Motif = "negative"
	True     Motif = "true"
	False    Motif = "false"
	In       Motif = "in"
	Out      Motif = "out"

	// Actions
	Change Motif = "change"
	Link   Motif = "link"
	Delete Motif = "delete"
)

type Pair struct {
	Alpha Color
	Beta  Color
}

type Mappings map[Motif]Pair

// === Globals =================================================================

var shadesMap = map[int]int{50: 0, 100: 1, 200: 2, 300: 3, 400: 4, 500: 5, 600: 6, 700: 7, 800: 8, 900: 9, 950: 10}
var shades = []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}

var orderedColors = []Color{
	Base, Surface,
	Rose, Berry, Cherry, Ruby, Red, Coral,
	Pumpkin, Orange, Sun, Gold,
	Honey, Yellow, Lemon, Acid,
	Lime, Spring, Green, Emerald, Jade, Forest, Leaf,
	Teal, Cyan, Aqua,
	Robin, Azure, Sky, Blue,
	Cobalt, Sapphire, Indigo,
	Lavender, Purple, Violet,
	Pink, Magenta,
	Brick, Rust, Beige, Olive, Moss, Zinc, Gray, Slate, Stone, Ash,
	White, Black,
}

var colors = Colors{
	Base:     "",
	Surface:  "",
	Rose:     "#fc0086",
	Berry:    "#fd016f",
	Cherry:   "#ff0457",
	Ruby:     "#f9043a",
	Red:      "#fd181a",
	Coral:    "#fb3d03",
	Pumpkin:  "#fd5802",
	Orange:   "#ff7220",
	Sun:      "#ff9004",
	Gold:     "#fead05",
	Honey:    "#ffcc00",
	Yellow:   "#fddf00",
	Lemon:    "#ecec00",
	Acid:     "#cdf118",
	Lime:     "#aae801",
	Spring:   "#86e401",
	Green:    "#58d300",
	Emerald:  "#28c624",
	Jade:     "#01b947",
	Forest:   "#03bb65",
	Leaf:     "#01c37e",
	Teal:     "#0ed39a",
	Cyan:     "#00e7cb",
	Aqua:     "#02eeef",
	Robin:    "#07e3fe",
	Azure:    "#0acbff",
	Sky:      "#0aafff",
	Blue:     "#0184fe",
	Cobalt:   "#256eff",
	Sapphire: "#4158fa",
	Indigo:   "#5a4aff",
	Lavender: "#6e40ff",
	Purple:   "#972eff",
	Violet:   "#c602fe",
	Pink:     "#ea0aeb",
	Magenta:  "#fd01b9",
	Brick:    "#847e80",
	Rust:     "#847f7e",
	Beige:    "#82807c",
	Olive:    "#7f817c",
	Moss:     "#7d817e",
	Zinc:     "#7c8280",
	Gray:     "#7c8182",
	Slate:    "#7d8184",
	Stone:    "#7e8084",
	Ash:      "#817f84",
	White:    "#cfcfcf",
	Black:    "#0d0d0d",
}

var mappings = Mappings{
	Info:    Pair{Alpha: Azure, Beta: Sky},
	Success: Pair{Alpha: Green, Beta: Emerald},
	Warning: Pair{Alpha: Yellow, Beta: Honey},
	Error:   Pair{Alpha: Red, Beta: Ruby},

	Primary: Pair{Alpha: Blue, Beta: Cobalt},
	Subtle:  Pair{Alpha: Slate, Beta: Stone},
	Accent:  Pair{Alpha: Orange, Beta: Sun},

	Positive: Pair{Alpha: Coral, Beta: Pumpkin},
	Negative: Pair{Alpha: Sapphire, Beta: Indigo},
	True:     Pair{Alpha: Aqua, Beta: Teal},
	False:    Pair{Alpha: Pink, Beta: Rose},
	In:       Pair{Alpha: Lavender, Beta: Violet},
	Out:      Pair{Alpha: Honey, Beta: Lemon},

	Change: Pair{Alpha: Azure, Beta: Blue},
	Link:   Pair{Alpha: Jade, Beta: Leaf},
	Delete: Pair{Alpha: Cherry, Beta: Ruby},
}

// === Handlers ================================================================

func Generate(seed string) Palette {
	p := make(Palette)
	for _, Color := range orderedColors {
		color, ok := colors[Color]
		if !ok {
			log.Println("[WARN]", Color, "not found in palette for seeding.")
			continue
		}
		if Color == Base || Color == Surface {
			color = Hex(seed)
		}
		oklch := color.toOklch()
		details := &ColorDetails{Color: Color, Base: *oklch, Shades: make(map[int]Details, len(shades))}
		p[Color] = details

		switch Color {
		case Base:
			details.generateBg()
		case Surface:
			details.generateFg(p[Base].Shades[50].Oklch)
		case Brick, Rust, Olive, Moss, Zinc, Slate, Gray, Stone, Ash, Beige:
			details.generateGrey()
		case White, Black:
			details.generateBW(Color)
		default:
			details.generateColor()
		}
	}
	return p
}

func (p *Palette) ToCSS(w io.Writer) {
	fmt.Fprintln(w, ":root {")
	for _, code := range orderedColors {
		colorDetails, ok := (*p)[code]
		if ok {
			colorDetails.ToCSS(w, code)
		}
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "@theme {")
	for _, code := range orderedColors {
		colorDetails, ok := (*p)[code]
		if ok {
			colorDetails.ToTheme(w, code)
		}
	}
	fmt.Fprintln(w, "")
	mappings.ToMotifs(w)
	fmt.Fprintln(w, "}")
}

func (c ColorDetails) ToCSS(w io.Writer, color Color) {
	for _, shadeKey := range shades {
		shade, ok := c.Shades[shadeKey]
		if !ok {
			continue
		}
		css := OklchToString(&shade.Oklch)
		fmt.Fprintf(w, "  --%s-%d: %s;\n", string(color), shadeKey, css)
	}
	fmt.Fprintln(w, "")
}

func (c ColorDetails) ToTheme(w io.Writer, color Color) {
	for _, shadeKey := range shades {
		_, ok := c.Shades[shadeKey]
		if !ok {
			continue
		}
		rootVar := fmt.Sprintf("var(--%s-%d)", string(color), shadeKey)
		fmt.Fprintf(w, "  --color-%s-%d: %s;\n", string(color), shadeKey, rootVar)
	}
	fmt.Fprintln(w, "")
}

func (m Mappings) ToMotifs(w io.Writer) {
	for motif, pair := range m {
		for _, shadeKey := range shades {
			rootVar := fmt.Sprintf("var(--%s-%d)", pair.Alpha, shadeKey)
			fmt.Fprintf(w, "  --color-%s-%d: %s;\n", motif, shadeKey, rootVar)
		}
		fmt.Fprintln(w, "")
		for _, shadeKey := range shades {
			rootVar := fmt.Sprintf("var(--%s-%d)", pair.Beta, shadeKey)
			fmt.Fprintf(w, "  --color-%s-%d: %s;\n", motif, shadeKey+1, rootVar)
		}
		fmt.Fprintln(w, "")
	}
}

func (p Palette) ToView() []palette.ColorScaleView {
	var views []palette.ColorScaleView

	bgcDetails, ok := p[Base]
	if !ok {
		return views
	}
	seed := bgcDetails.Base

	convertScale := func(clr Color, scale *ColorDetails) palette.ColorScaleView {
		view := palette.ColorScaleView{
			Name:   string(scale.Color),
			Code:   string(clr),
			Value:  scale.Shades[600].Oklch,
			Shades: make([]palette.Shade, len(shades)),
		}

		for i, shadeVal := range shades {
			colorDetails, ok := scale.Shades[shadeVal]
			if !ok {
				continue
			}

			color := colorDetails.Oklch
			rl, cr := OklchCompare(seed, color)
			hue := toDegree(color.H)

			radius := 37.0

			scaledRadiusC := radius * math.Tanh(6.0*color.C)
			totalDistanceC := scaledRadiusC

			scaledRadiusL := radius * math.Pow(color.L, 1.5)
			totalDistanceL := scaledRadiusL

			angle := -color.H
			view.Shades[i] = palette.Shade{
				Code:  string(clr),
				Value: shadeVal,
				RL:    rl,
				CR:    cr,
				L:     fmt.Sprintf("%0.1f%%", color.L*100),
				C:     fmt.Sprintf("%0.3f", color.C),
				H:     fmt.Sprintf("%0.1f", hue),
				Hex:   OklchToHex(&color),
				Cx:    50.0 + totalDistanceC*math.Cos(angle),
				Cy:    50.0 + totalDistanceC*math.Sin(angle),
				Clx:   50.0 + totalDistanceL*math.Cos(angle),
				Cly:   50.0 + totalDistanceL*math.Sin(angle),
			}
		}
		return view
	}

	for _, code := range orderedColors {
		details, ok := p[code]
		if ok {
			views = append(views, convertScale(code, details))
		}
	}

	return views
}
