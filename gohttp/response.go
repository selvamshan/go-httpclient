package gohttp

import (
	"encoding/json"
	"net/http"
)

// Response ...
type Response struct{
	status string
	statusCode int
	headers http.Header
	body []byte
}

// Status ...
func (r *Response) Status() string {
	return r.status
}

// StatusCode ...
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Headers ...
func (r *Response) Headers() http.Header {
	return r.headers
}

// Bytes ...
func (r *Response) Bytes() []byte {
	return r.body
}

//String ...
func (r *Response) String() string {
	return string(r.body)
}


//UnmarshalJson ...
func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
