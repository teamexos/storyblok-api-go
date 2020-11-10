package storyblok

import "fmt"

const (
	errCodeRequestDecodeFailed = "failed to decode response"
	errCodeRequestDoFailed     = "failed to make request"
	errCodeRequestSetupFailed  = "failed to setup request"
)

type (
	// ResponseError is used for errors this package will return
	ResponseError struct {
		StatusCode int
		Message    string
	}

	storyblokError struct {
		Error string `json:"error"`
	}
)

func (e *ResponseError) Error() string {
	return fmt.Sprintf("statusCode: %d, error: %s", e.StatusCode, e.Message)
}

// NewResponseError returns a RequestError
func NewResponseError(s int, m string) *ResponseError {
	return &ResponseError{
		StatusCode: s,
		Message:    m,
	}
}
