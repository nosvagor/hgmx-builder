package htmx

func Get(url string, opts ...Options) *Props {
	return Init(GET, url, opts...)
}

func Post(url string, opts ...Options) *Props {
	return Init(POST, url, opts...)
}

func Put(url string, opts ...Options) *Props {
	return Init(PUT, url, opts...)
}

func Patch(url string, opts ...Options) *Props {
	return Init(PATCH, url, opts...)
}

func Delete(url string, opts ...Options) *Props {
	return Init(DELETE, url, opts...)
}
