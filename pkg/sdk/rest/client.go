package rest

/*
Client wraps RequestBuilder to provide a minimal REST client abstraction.
*/
type Client struct {
	RequestBuilder
}

/*
New creates a REST client bound to the given base URL.
*/
func New(baseURL string) Client {
	return Client{
		NewRequestBuilder(baseURL),
	}
}
