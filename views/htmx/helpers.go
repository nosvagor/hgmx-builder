package htmx

// Target creates an Options map with the hx-target attribute set.
// It allows you to target a different element for swapping than the one
// issuing the AJAX request.
//
// The target can be a CSS query selector or one of htmx's special values:
//   - "this": targets the element itself.
//   - "closest <selector>": finds the closest ancestor that matches the selector.
//   - "find <selector>": finds the first matching child descendant.
//   - "next" or "next <selector>": targets the next sibling or next matching sibling.
//   - "previous" or "previous <selector>": targets the previous sibling or previous matching sibling.
//
// https://htmx.org/attributes/hx-target/
func Target(target string) Options {
	return Options{"hx-target": target}
}

func (o *Options) Target(target string) {
	o.Merge(Options{"hx-target": target})
}

// Trigger creates an Options map with the hx-trigger attribute set.
// It allows you to specify what triggers the AJAX request.
//
// By default, triggers are based on the "natural" event for an element (e.g.,
// "click" for buttons, "change" for inputs). You can specify a different event
// or a combination of events.
//
// Examples of triggers:
//   - "click": trigger on a click.
//   - "mouseenter": trigger on mouse enter.
//   - "load": trigger on page load.
//   - "every 2s": trigger every 2 seconds.
//
// Trigger modifiers can be used:
//   - `once`: trigger only once.
//   - `changed`: trigger only when the value has changed.
//   - `delay:<time>`: wait for the given time before triggering (e.g., "delay:1s").
//   - `throttle:<time>`: trigger at most once per given time (e.g., "throttle:500ms").
//   - `from:<selector>`: listen for the event on a different element.
//
// https://htmx.org/attributes/hx-trigger/
func Trigger(trigger string) Options {
	return Options{"hx-trigger": trigger}
}

func (o *Options) Trigger(trigger string) {
	o.Merge(Options{"hx-trigger": trigger})
}

// Swap creates an Options map with the hx-swap attribute set.
// It allows you to control how the response content is swapped into the DOM.
//
// The default swap strategy is "innerHTML".
//
// Possible values include:
//   - "innerHTML": replaces the inner HTML of the target element (default).
//   - "outerHTML": replaces the entire target element.
//   - "beforebegin": inserts the response before the target element.
//   - "afterbegin": prepends the response to the target element's content.
//   - "beforeend": appends the response to the target element's content.
//   - "afterend": inserts the response after the target element.
//   - "delete": deletes the target element.
//   - "none": does not swap any content.
//
// You can also modify the swap behavior with modifiers like `swap:<time>` and `settle:<time>`.
//
// https://htmx.org/attributes/hx-swap/
func Swap(swap string) Options {
	return Options{"hx-swap": swap}
}
func (o *Options) Swap(swap string) {
	o.Merge(Options{"hx-swap": swap})
}
