<p align="right">
    <a href="https://github.com/nosvagor/hgmx-builder/stargazers">
        <img
            src="https://img.shields.io/github/stars/nosvagor/hgmx-builder?color=ecc45d&logo=apachespark&labelColor=222536&logoColor=ecc45d&style=for-the-badge"
            title="what is love, baby don't hurt me"
        >
    </a>
</p>

<p align="center">
<a href="https://hgmx.dev">
<kbd>
 <img 
   src="https://github.com/nosvagor/hgmx-builder/blob/master/views/static/svgs/htmx-full.svg?raw=true"
   title="hgmx.dev"
   alt="hgmx logo" width="500"
  /> 
</kbd>
</a>

> a Go powered, hypermedia driven, open design system;
> utilizing HTML-first tools like [htmx](https://htmx.org), [\_hyperscript](https://hyperscript.org), [templ](https://templ.guide), and [tailwind](https://tailwindcss.com).

## Open Design System?

An ODS is cousin to the idea of a framework, but inverts ownership. That is, it allows you to build a library that you own, one you can improve.

The general idea comes from [shadcn/ui](https://ui.shadcn.com/docs). hgmx is an evolution of this concept, one that fills niche carved out of the javascript ecosystem by [htmx](https://htmx.org/docs), which encourages using [**H**ypermedia **A**s **T**he **E**ngine **O**f **A**pplication **S**tate](https://htmx.org/essays/hypermedia-friendly-scripting/#prime_directive). The divergence from htmx stems from the fact is that htmx does not care what backend you use&mdash;but hgmx does&mdash;it's built specifically for Go. Futhermore, it diverges from shadcn/ui in that it extends the open design system into the backend with more opinionated defaults.

Overall, the goal of hgmx is to blur the boundary between frontend and backend, and to create a holistic design system with foundation that can be grown into any direction you need.

## Unique Philosophies (UI)

- **Semantic Theming**: Baked in structural theming to create a consistent, customizable color language to communicate with.
- **Boutique Components**: Components should feel artisan crafted, sparking joy when used.

### Semantic Theming

Traditional dark/light mode toggles are a false dichotomy. hgmx's theming system generates entire color palettes from a single seed color using the OKLCH color space. This creates harmonious color relationships automatically to pick and choose colors to build up a theme easily.

The goal is to lean into the undertalized bandwidth that intentional color assocatiations can do to reduce cognitive load. Think of how films subtle use color to communciate emotions, or how syntax highlighting can make reading code easier.

#### Key Principles

- **Heirarchical**: Inspired by 60/30/10 rule. Choose background color (60%), primary accent (30%), and action color (10%). &rarr; get a cohesive palette built into components.
- **Semantic Motifs**: Components use semantic names (surface, base, primary, accent, error, success, change, delete, true/false, etc..); not specific colors. Geometrically determined from core choices.
- **Runtime Flexibility**: Swap entire theme by changing root CSS variables, no changes to components needed.
- **Color Agnostic**: Components built to work without colors, but effortlessly work with them to add depth&mdash;you choose the syntatic meaning of the colors.

### Boutique Components

In an age of slop, hgmx's components are crafted to stand out in a effortful way to signal intentionality.

The barbel hypothesis of maturing markets suggests a convex relationship betwen quality and quantitiy. That is, either you lean into cheap slop and scale to the masses, or reject and go boutique and establish a brand with specilized focus on quality over a particular dimension.

There is nohting wrong with the perusing quantity, it's needed. But hgmx leans into the boutique approach, creating components tailored to stand out, sparking joy when used.

#### Key Principles

- **Motion Enabled**: Subtle animations built in to command attention and signal state.
- **Keyboard Driven**: Frequently used components built with optional keyboard navigation.
- **MAYA (Most Adavanced, Yet Acceptable)**: Create unique takes on components that push the boundaries, yet remain intutive and usable.

## Architecture

hgmx is designed to be built on top of an event-driven modular monolith foundation, extending the event focused prime directive of HATEOAS. It can be used with alternaitve architectural patterns and backend libraries, but was developed with tools like [rabbitmq](https://www.rabbitmq.com), [echo](https://echo.labstack.com), [gorm](https://gorm.io), and [signoz](https://signoz.io) in use.

```
 .
 │
 ├─ cmd
 │  └─ main.go
 │
 ├─ internal
 │  ├─ handlers
 │  │  ├─ api
 │  │  └─ web
 │  │
 │  ├─ jobs
 │  │
 │  ├─ server
 │  │  ├─ bus
 │  │  ├─ logging
 │  │  ├─ middleware
 │  │  └─ routes
 │  │
 │  ├─ services
 │  │  ├─ {service}
 │  │  │  └─ models
 │  │  ├─ ...
 │  │  └─ {service}
 │  │     └─ models
 │  │
 │  ├─ store
 │  │
 │  └─ utils
 │
 └─ views
    ├─ components
    │  ├─ {component}
    │  ├─ ...
    │  └─ {component}
    │
    ├─ icons
    │
    ├─ pages
    │  ├─ {page}
    │  ├─ ...
    │  └─ {page}
    │
    ├─ static
    │  ├─ css
    │  ├─ favicon
    │  ├─ fonts
    │  ├─ images
    │  ├─ scripts
    │  └─ svgs
    │
    └─ utils
        ├─ html
        ├─ htmx
        ├─ ...
        └─ {util}
```
