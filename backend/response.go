package backend

import "net/http"

// Response modelise a REST API response, the model is meant to be "compiled"
// by the engine overlay.
type Response struct {
	// Err during the execution
	Err error
	// ErrMessage user-friendly message or identifier for the source of the error
	ErrMessage string
	// Data to return in the response's body
	Data interface{}
	// StatusCode of the response, default is 200
	StatusCode int
	// ContentType of the response, default is application/json
	ContentType string
}

// NewResponse create a new Answer with the default values
func NewResponse() *Response {
	return &Response{
		ContentType: "application/json",
		StatusCode:  http.StatusOK,
	}
}

// WithErr add the error to the answer
func (r *Response) WithErr(err error, message string) *Response {
	r.Err = err
	r.ErrMessage = message
	return r
}

// WithData add the data to the answer
func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

// WithStatusCode add the status code to the answer
func (r *Response) WithStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

// WithContentType add the Content Type to the answer
func (r *Response) WithContentType(contentType string) *Response {
	r.ContentType = contentType
	return r
}
