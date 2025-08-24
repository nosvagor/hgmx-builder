# Icon Development Guide

## 1. Find an SVG

Start by finding an SVG for your icon. A great resource for high-quality, consistent icons is [Lucide](https://lucide.dev/).

When you have the SVG code, you will only need the inner elements like `<path>`, `<circle>`, etc., not the surrounding `<svg>` tag.

## 2. Create a New Icon File

In the `views/components/icon/icons/` directory, create a new file with a descriptive name that matches the icon's appearance or intended action, e.g., `click.templ` or `orbit.templ`.

## 3. Define the Icon Struct

In your new file, define a public function that returns an `*icon.Icon` pointer. This function will instantiate and return an `icon.Icon` struct.

```go
func {IconName}() *icon.Icon {
    return &icon.Icon{
        Name:         name, // type name string (lowerCamelCase)
        Handler:      templ.NewOnceHandle(),
        Script:       {iconName}Animation,
        Icon:         {iconName}Content,
    }
}

```

The name should be added to icons/icons.go. This helps ensure icon names are unique.

```go
package icons

type name string

const (
    ... // others icons
	orbit name = "orbit"
    ...
)
```

### Icon Registry

For web docs usage, update `/internal/handlers/pages/web/icons/registry.go` to include the icon and metadata.

```go

var Registry = map[string]*IconDetails{
    ... // others icons
    "orbit": {
		Icon:         *icons.Orbit(),
		Tags:         []string{"orbit", "circle", "rotation", "spin"},
		Categories:   []string{"shapes", "motion"},
		Contributors: []string{"nosvagor", "lucide"},
	},
	"userCircle": {
		Icon:         *icons.UserCircle(),
		Tags:         []string{"user", "circle", "account", "profile"},
		Categories:   []string{"people", "account", "shapes"},
		Contributors: []string{"nosvagor", "lucide"},
	},
    ...
}
```

### Icon Struct Fields

- `Name`: A unique, lowerCamelCase name for the icon (e.g., "userCheck")
- `Handler`: Use `templ.NewOnceHandle()` to ensure scripts are only rendered once
- `Script`: A reference to the templ component containing the animation script
- `Icon`: A reference to the templ component containing the SVG markup

### Brand Icons

Brand icons won't be as nice as lucide icons, as they might have to be manually deconstructed in order to be animated.
Use the icons/inkscape-template.xml for help for creating a brand icon consitency.

## 4. Create the SVG Content Component

Define a templ component for the SVG markup. The component name should be in format `{iconName}Content`.

```go
templ {iconName}Content() {
    @icon.SVG() {
        // Paste the inner SVG elements here.
        <path d="..."/>
        <circle cx="..." cy="..." r="..."/>
    }
}
```

## 5. Create the Animation Component

Define a templ component for the icon's animation. The component name should be in the format `{iconName}Animation`.

[!IMPORTANT]

1. Select the SVG element as a constant right above the animation function that animates it.
2. If parts of the animation have delays or are sequential, then animations so that you read the sequence from top to bottom.

[motion docs](https://motion.dev/docs/quick-start)

```javascript
templ {iconName}Animation() {
    <script type="module">
        import { animate, easeOut } from "/static/scripts/motion.min.js";
        const τmax = 0.42, τmid = 0.27, τmin = 0.15;

        // The animation function must be attached to the `window` object.
        // The function name must be in the format `{iconName}Animate`.
        window.{iconName}Animate = (el) => {
            // `el` is the <icon> element. You can query its children.
            const path = el.querySelector('path');
            animate(path,
                { opacity: [0, 1] }, // properties to animate
                { duration: τmax, easeOut } // animation settings
            );

            // continue to query parts of svg, if needed
            circle = el.querySelector('circle[r="2"]');
            animate(circle,
                { scale: [0, 1.25, 1] },
                { duration: τmax, easeOut }
            );
        };
    </script>
}
```

### Animation Style Guidelines

When implementing animations, especially those involving geometric concepts like rotations or translations, using Greek letters for variables is encouraged. This creates a highly readable and distinct style, making it clear that the code is dealing with physic based animations, and showing guidelines were read and followed.

#### Commonly used Greek letters and their conventional meanings:

- **Θ (Theta), α (Alpha), β (Beta)**: Angles or rotation values
- **Δ (Delta)**: Change or difference in a value (e.g., change in angle)
- **π (Pi)**: Trigonometric transformations
- **μ (Mu)**: Scaling values
- **ϵ (Epsilon)**: Small values
- **Σ (Sigma)**: Summations or cumulative values
- **τ (Tau)**: Time values

#### Example from `orbit.templ`:

```javascript
const Θ = Number(el.dataset.rotation || 0);
const Δ = Θ + 360 + Math.random() * 100;
animate(moons, { rotate: [Θ, Δ] });
```

### Timing Constants

- **τmax (0.42)**: Maximum duration for main animations
- **τmid (0.27)**: Medium duration for delayed parts
- **τmin (0.15)**: Quick duration for subtle movements, i.e., difference between parts.

### Easing Import Control

Import easing functions directly for better performance and consistency:

```javascript
import { animate, easeOut } from "/static/scripts/motion.min.js";
```

**Preferred usage:**

```javascript
{
  duration: τmid, easeOut;
} // Direct import (recommended)
```

**Alternative:**

```javascript
{ duration: τmid, ease: "easeInOut" }  // String reference
```

### Animation Suggestions

- **Opacity and pathLength**: Great for creating drawing effects for simpler icons
- **Scale**: Can create nice subtle bounce effects
- **Rotations and translations**: Great for creating movement for more complex icons
- **Staggers**: Excellent for icons with lots of parts.
