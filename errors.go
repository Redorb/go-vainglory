package vainglory

import (
	"fmt"
	"net/http"
)

// RESTError stores error information about a request with a bad response code.
// Message is not always present, there are cases where api calls can fail
// without returning a json message.
type RESTError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
}

func newRestError(req *http.Request, resp *http.Response, body []byte) *RESTError {
	return &RESTError{
		Request:      req,
		Response:     resp,
		ResponseBody: body,
	}
}

func (r RESTError) Error() string {
	return fmt.Sprintf("HTTP %s, %s", r.Response.Status, r.ResponseBody)
}
