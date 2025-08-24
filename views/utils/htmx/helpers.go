package htmx

import "fmt"

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

// Boost adds the hx-boost attribute to an element, which "boosts"
// regular anchors and form submissions into AJAX requests.
// Setting it to "true" enables boosting for all child anchors and forms.
//
// https://htmx.org/attributes/hx-boost/
func Boost() Options {
	return Options{"hx-boost": "true"}
}

func (o *Options) Boost() {
	o.Merge(Options{"hx-boost": "true"})
}

// PushURL pushes a new URL into the browser's history.
// It can be set to "true" to use the request URL, or a specific URL.
//
// https://htmx.org/attributes/hx-push-url/
func PushURL(val string) Options {
	return Options{"hx-push-url": val}
}

func (o *Options) PushURL(val string) {
	o.Merge(Options{"hx-push-url": val})
}

// ReplaceURL replaces the URL in the browser location bar.
// Can be a URL string, "true", or "false".
//
// https://htmx.org/attributes/hx-replace-url/
func ReplaceURL(val string) Options {
	return Options{"hx-replace-url": val}
}

func (o *Options) ReplaceURL(value string) {
	o.Merge(ReplaceURL(value))
}

// Indicator shows a loading indicator during the AJAX request.
// The value should be a CSS selector for the indicator element.
//
// https://htmx.org/attributes/hx-indicator/
func Indicator(selector string) Options {
	return Options{"hx-indicator": selector}
}

func (o *Options) Indicator(selector string) {
	o.Merge(Options{"hx-indicator": selector})
}

// Confirm shows a confirmation dialog before issuing the request.
//
// https://htmx.org/attributes/hx-confirm/
func Confirm(message string) Options {
	return Options{"hx-confirm": message}
}

func (o *Options) Confirm(message string) {
	o.Merge(Options{"hx-confirm": message})
}

// Select allows you to select a subset of the response to be swapped.
// The value should be a CSS selector.
//
// https://htmx.org/attributes/hx-select/
func Select(selector string) Options {
	return Options{"hx-select": selector}
}

func (o *Options) Select(selector string) {
	o.Merge(Options{"hx-select": selector})
}

// On allows you to handle events with inline scripts on elements.
// The format is hx-on:<event-name>="<script>".
//
// https://htmx.org/attributes/hx-on/
func On(event, script string) Options {
	return Options{fmt.Sprintf("hx-on:%s", event): script}
}

func (o *Options) On(event, script string) {
	o.Merge(On(event, script))
}

// SelectOOB selects content to swap in from a response,
// somewhere other than the target (out of band).
// The value is a CSS selector.
//
// https://htmx.org/attributes/hx-select-oob/
func SelectOOB(selector string) Options {
	return Options{"hx-select-oob": selector}
}

func (o *Options) SelectOOB(selector string) {
	o.Merge(SelectOOB(selector))
}

// SwapOOB marks an element to be swapped in from a response out of band.
// The value can be a swap style (e.g. "innerHTML") or "true".
//
// https://htmx.org/attributes/hx-swap-oob/
func SwapOOB(value any) Options {
	return Options{"hx-swap-oob": value}
}

func (o *Options) SwapOOB(value any) {
	o.Merge(SwapOOB(value))
}

// Vals adds values to submit with the request in JSON format.
//
// https://htmx.org/attributes/hx-vals/
func Vals(json string) Options {
	return Options{"hx-vals": json}
}

func (o *Options) Vals(json string) {
	o.Merge(Vals(json))
}

// Disable disables htmx processing for the given node and any children nodes.
//
// https://htmx.org/attributes/hx-disable/
func Disable() Options {
	return Options{"hx-disable": "true"}
}

func (o *Options) Disable() {
	o.Merge(Disable())
}

// DisabledElt adds the 'disabled' attribute to the specified elements
// while a request is in flight. The value is a CSS selector.
//
// https://htmx.org/attributes/hx-disabled-elt/
func DisabledElt(selector string) Options {
	return Options{"hx-disabled-elt": selector}
}

func (o *Options) DisabledElt(selector string) {
	o.Merge(DisabledElt(selector))
}

// Disinherit controls and disables automatic attribute inheritance for child nodes.
// The value can be an attribute name or "*".
//
// https://htmx.org/attributes/hx-disinherit/
func Disinherit(value string) Options {
	return Options{"hx-disinherit": value}
}

func (o *Options) Disinherit(value string) {
	o.Merge(Disinherit(value))
}

// Encoding changes the request encoding type, e.g., "multipart/form-data".
//
// https://htmx.org/attributes/hx-encoding/
func Encoding(encoding string) Options {
	return Options{"hx-encoding": encoding}
}

func (o *Options) Encoding(encoding string) {
	o.Merge(Encoding(encoding))
}

// Headers adds to the headers that will be submitted with the request.
// The value should be a JSON string.
//
// https://htmx.org/attributes/hx-headers/
func Headers(json string) Options {
	return Options{"hx-headers": json}
}

func (o *Options) Headers(json string) {
	o.Merge(Headers(json))
}

// Include includes additional data in requests. The value is a CSS selector.
//
// https://htmx.org/attributes/hx-include/
func Include(selector string) Options {
	return Options{"hx-include": selector}
}

func (o *Options) Include(selector string) {
	o.Merge(Include(selector))
}

// Inherit controls and enables automatic attribute inheritance for child nodes
// if it has been disabled by default.
//
// https://htmx.org/attributes/hx-inherit/
func Inherit(value string) Options {
	return Options{"hx-inherit": value}
}

func (o *Options) Inherit(value string) {
	o.Merge(Inherit(value))
}

// Params filters the parameters that will be submitted with a request.
//
// https://htmx.org/attributes/hx-params/
func Params(filter string) Options {
	return Options{"hx-params": filter}
}

func (o *Options) Params(filter string) {
	o.Merge(Params(filter))
}

// Preserve specifies elements to keep unchanged between requests.
// Can be "true" or a comma-separated list of IDs.
//
// https://htmx.org/attributes/hx-preserve/
func Preserve(value string) Options {
	return Options{"hx-preserve": value}
}

func (o *Options) Preserve(value string) {
	o.Merge(Preserve(value))
}

// Prompt shows a prompt() before submitting a request.
//
// https://htmx.org/attributes/hx-prompt/
func Prompt(message string) Options {
	return Options{"hx-prompt": message}
}

func (o *Options) Prompt(message string) {
	o.Merge(Prompt(message))
}

// Request configures various aspects of the request.
// The value should be a JSON string.
//
// https://htmx.org/attributes/hx-request/
func Request(json string) Options {
	return Options{"hx-request": json}
}

func (o *Options) Request(json string) {
	o.Merge(Request(json))
}

// Sync controls how requests made by different elements are synchronized.
//
// https://htmx.org/attributes/hx-sync/
func Sync(spec string) Options {
	return Options{"hx-sync": spec}
}

func (o *Options) Sync(spec string) {
	o.Merge(Sync(spec))
}

// Validate forces elements to validate themselves before a request.
//
// https://htmx.org/attributes/hx-validate/
func Validate() Options {
	return Options{"hx-validate": "true"}
}

func (o *Options) Validate() {
	o.Merge(Validate())
}

// Preload enables the preload extension which allows you to load content
// in advance of user interaction.
//
// e.g., "mouseenter" for hover preloading, which has built in 100ms threshold.
//
// The default event is "mousedown" which gives about 100ms of preloading.
//
// https://htmx.org/extensions/preload/
func Preload(event ...string) Options {
	if len(event) == 0 || event[0] == "" {
		return Options{"preload": "mousedown"}
	}

	return Options{"preload": event[0]}
}

func (o *Options) Preload(event ...string) {
	o.Merge(Preload(event...))
}
