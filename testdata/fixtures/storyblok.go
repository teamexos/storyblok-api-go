package fixtures

import (
	"bytes"
	"io"
	"io/ioutil"
)

// ResponseStoryNotFound returns an example response for an invalid story slug/id/uuid
func ResponseStoryNotFound() io.ReadCloser {
	json := `{"error":["This record could not be found"]}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ResponseUnauthorized returns an example Unauthorized response
func ResponseUnauthorized() io.ReadCloser {
	json := `{"error":"Unauthorized"}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}
