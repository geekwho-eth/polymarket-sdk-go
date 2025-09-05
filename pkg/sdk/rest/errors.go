package rest

import (
	"fmt"
)

/*
RestError represents an HTTP error returned by the remote API.
*/
type RestError struct {
	Code      int
	Message   string
	RequestID string
}

func (err *RestError) Error() string {
	return fmt.Sprintf("[RestError] code=%d, message=%s, requestID=%s", err.Code, err.Message, err.RequestID)
}

/*
NewRestError constructs a RestError with HTTP status code, message and request ID.
*/
func NewRestError(code int, message, requestID string) error {
	return &RestError{
		Code:      code,
		Message:   message,
		RequestID: requestID,
	}
}

func (err *RestError) GetCode() int {
	return err.Code
}

func (err *RestError) GetMessage() string {
	return err.Message
}

func (err *RestError) GetRequestID() string {
	return err.RequestID
}
