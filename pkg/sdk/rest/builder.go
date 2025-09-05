package rest

import (
	"net/url"
	"strings"
	"time"
)

// HTTP methods we support
const (
	POST    = "POST"
	GET     = "GET"
	HEAD    = "HEAD"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"

	// Types we support.
	TypeJSON = "json"
)

/*
RequestInterface defines the minimal request builder surface used by services.
*/
type RequestInterface interface {
	//Post(path string) *RequestBuilder
	//Put(path string) *RequestBuilder
	Get(path string) *RequestBuilder
	//Delete(path string) *RequestBuilder
	APIVersion() GroupAPIVersion
}

// GroupAPIVersion contains the "group" and the "version", which uniquely identifies the API.
type GroupAPIVersion struct {
	Group   string
	Version string
}

/*
RequestBuilder builds and configures HTTP requests in a fluent, chainable style.
*/
type RequestBuilder struct {
	baseURL string
	method  string
	path    string
	url     string
	params  url.Values
	timeout time.Duration
}

// NewRequestBuilder creates a Rest Request builder.
func NewRequestBuilder(baseURL string) RequestBuilder {
	return RequestBuilder{
		baseURL: baseURL,
		params:  make(url.Values),
	}
}

// Get sets the path for GET request.
func (r *RequestBuilder) Get(path string) *RequestBuilder {
	r.method = GET
	r.path = path
	return r
}

// Params sets a query parameter k=v (only if v is not empty).
func (r *RequestBuilder) Params(k, v string) *RequestBuilder {
	if strings.TrimSpace(k) != "" && strings.TrimSpace(v) != "" {
		r.params.Set(k, v)
	}
	return r
}

// Timeout sets per-request timeout duration for Do().
func (r *RequestBuilder) Timeout(d time.Duration) *RequestBuilder {
	r.timeout = d
	return r
}

/*
APIVersion sets API version when needed (placeholder for future extensions).
*/
func (r *RequestBuilder) APIVersion() *RequestBuilder {
	//r.apiVersion = "v1"
	return r
}
