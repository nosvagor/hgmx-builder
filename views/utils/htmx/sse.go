package htmx

import "fmt"

// SSE enables the Server-Sent Events extension.
// This attribute is required to use SSE features.
// https://htmx.org/extensions/sse/
func SSE() Options {
	return Options{"hx-ext": "sse"}
}

func (o *Options) SSE() {
	o.Merge(SSE())
}

// SSEConnect specifies the URL to connect to for Server-Sent Events.
//
// https://htmx.org/extensions/sse/#connecting-to-an-sse-server
func SSEConnect(url string) Options {
	return Options{"sse-connect": url}
}

func (o *Options) SSEConnect(url string) {
	o.Merge(SSEConnect(url))
}

// SSESwap specifies the name of the SSE event to listen for.
// When an event with this name is received, its data will be swapped
// into the element's innerHTML. Multiple events can be listened for
// by separating them with commas.
//
// https://htmx.org/extensions/sse/#receiving-named-events
func SSESwap(eventName string) Options {
	return Options{"sse-swap": eventName}
}

func (o *Options) SSESwap(eventName string) {
	o.Merge(SSESwap(eventName))
}

// SSEClose specifies an event name that will cause the SSE connection
// to be closed gracefully.
//
// https://htmx.org/extensions/sse/
func SSEClose(eventName string) Options {
	return Options{"sse-close": eventName}
}

func (o *Options) SSEClose(eventName string) {
	o.Merge(SSEClose(eventName))
}

// SSETrigger triggers an htmx request when an SSE event is received.
// The value should be in the format 'sse:<event_name>'.
//
// https://htmx.org/extensions/sse/#trigger-server-callbacks
func SSETrigger(eventName string) Options {
	return Options{"hx-trigger": fmt.Sprintf("sse:%s", eventName)}
}

func (o *Options) SSETrigger(eventName string) {
	o.Merge(SSETrigger(eventName))
}
