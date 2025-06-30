# HGMX

> A boutique, open design system (ODS) built with Go Templ, HTMX, and Hyperscript, and Motion.js

## Vision

In an age where UI can be instantly generated, hgmx design system takes a different approach&mdash; _crafting_ a unique, yet cohesive, set of componenets with small touches that spark joy when used.

## Core Philosophy

1. **Boutique Components**: Each component should feel artisan crafted, yet customizable.
2. **Opinionated First Principles**: Start minimal, with opinionated defaults to manage complexity.
3. **Semantic Theming**: Rich color system with semantic variables for easy theme adjustments.
4. **HATEOAS**: Hypermedia driven state management.

### Theme

Traditional dark/light mode toggles are a false dichotomy. Our theming system generates entire color palettes from a single seed color using the perceptually uniform OKLCH color space. This creates harmonious color relationships automatically.

#### How It Works

1. **Seed-Based Generation**: inspired by 60/30/10 rule. choose background color (60%), primary accent (30%), and action color (10%). &rarr; get an cohesive palette.
2. **Semantic Mappings**: Components use semantic names (primary, error, success); not specific colors. Geometrically determined from core choices.
3. **Automatic Relationships**: Colors are generated with proper contrast ratios and perceptual balance.
4. **Runtime Flexibility**: Swap entire themes by changing root CSS variables, no component changes needed.

#### Semantic Motifs

```
Signaling: info, success, warning, error
Branding:  primary, secondary, accent
States:    positive/negative, true/false, in/out
Actions:   change, link, delete
```

Each motif maps to a color pair for subtle variations to add another simple dimension to add choice themes.

## Current Status

### Completed

- [x] Base HTML structure with smart asset hashing
- [x] Rich color system generated from seed color alone.
- [x] Navbar

### Rough Roadmap

#### Phase 0 (continuous; ongoing): Abstract or Enhance

1. **Navbar decomposition** Identiy and create key components likely re-used in other places outside of navbar
2. **Active state** Estalish shared visual lanugage for focus, active, hover, and other common interactive states.

#### Phase 1: Navigation Foundation

1. **Sidebar**, for keeping track of current position if navigation is complex (like in a docs page).

2. **Command Mode**
   - Fuzzy search
   - Action shortcuts
   - Recent/frequent destinations
   - Context-aware suggestions
