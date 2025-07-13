package htmx

// Get configures an element to make a GET request to the given URL.
func Get(url string, opts ...Options) *Props {
	return Init(GET, url, opts...)
}

// Post configures an element to make a POST request to the given URL.
func Post(url string, opts ...Options) *Props {
	return Init(POST, url, opts...)
}

// Put configures an element to make a PUT request to the given URL.
func Put(url string, opts ...Options) *Props {
	return Init(PUT, url, opts...)
}

// Patch configures an element to make a PATCH request to the given URL.
func Patch(url string, opts ...Options) *Props {
	return Init(PATCH, url, opts...)
}

// Delete configures an element to make a DELETE request to the given URL.
func Delete(url string, opts ...Options) *Props {
	return Init(DELETE, url, opts...)
}
