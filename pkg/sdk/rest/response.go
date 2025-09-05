package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

// RestRawResult raw body and error
type RestRawResult struct {
	body []byte
	err  error
}

// DecodeInto unmarshals the response body into v if no previous error.
func (res *RestRawResult) DecodeInto(v interface{}) error {
	if res.err != nil {
		return res.err
	}
	return json.Unmarshal(res.body, v)
}

type Response interface {
	ParseErrorFromHTTPResponse(code int, body []byte) error
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type BaseResponse struct {
	ErrorResponse
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(code int, body []byte) error {
	// http code 200 is ok
	if code != http.StatusOK {
		if err := json.Unmarshal(body, r); err != nil {
			return err
		}
		return NewRestError(code, r.Error, "--")
	}

	return nil
}

// ParseRawResponse parse raw http response
func ParseRawResponse(rawResponse *http.Response) ([]byte, error) {
	defer rawResponse.Body.Close()
	body, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, err
	}

	baseResponse := BaseResponse{}

	if err := baseResponse.ParseErrorFromHTTPResponse(rawResponse.StatusCode, body); err != nil {
		return nil, err
	}

	return body, nil
}
