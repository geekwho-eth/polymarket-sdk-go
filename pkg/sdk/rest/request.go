package rest

import (
	"context"
	"net/http"
	"net/url"
)

// Do executes the request and returns a result wrapper.
func (r *RequestBuilder) SendRequest(ctx context.Context) *RestRawResult {
	u, err := url.Parse(r.baseURL + r.path)
	if err != nil {
		return &RestRawResult{err: err}
	}
	q := u.Query()
	for k, vs := range r.params {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()

	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, r.method, u.String(), nil)
	if err != nil {
		return &RestRawResult{err: err}
	}
	client := &http.Client{Timeout: r.timeout}

	resp, err := client.Do(req)

	if err != nil {
		return &RestRawResult{err: err}
	}
	body, err := ParseRawResponse(resp)
	if err != nil {
		return &RestRawResult{err: err}
	}
	defer resp.Body.Close()

	return &RestRawResult{body: body, err: err}
}
