package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	HeaderContentType = "Content-Type"
)

var (
	HeaderContentTypeJSON = []string{"application/json"}
)

type Transport interface {
	Perform(*http.Request) (*http.Response, error)
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       io.ReadCloser
}

// String returns the response as a string.
//
// The intended usage is for testing or debugging only.
//
func (r *Response) String() string {
	var (
		out = new(bytes.Buffer)
		b1  = bytes.NewBuffer([]byte{})
		b2  = bytes.NewBuffer([]byte{})
		tr  io.Reader
	)

	if r != nil && r.Body != nil {
		tr = io.TeeReader(r.Body, b1)
		defer r.Body.Close()

		if _, err := io.Copy(b2, tr); err != nil {
			out.WriteString(fmt.Sprintf("<error reading response body: %v>", err))
			return out.String()
		}
		defer func() { r.Body = ioutil.NopCloser(b1) }()
	}

	if r != nil {
		out.WriteString(fmt.Sprintf("[%d %s]", r.StatusCode, http.StatusText(r.StatusCode)))
		if r.StatusCode > 0 {
			out.WriteRune(' ')
		}
	} else {
		out.WriteString("[0 <nil>]")
	}

	if r != nil && r.Body != nil {
		out.ReadFrom(b2) // errcheck exclude (*bytes.Buffer).ReadFrom
	}

	return out.String()
}

// IsError returns true when the response status indicates failure.
//
func (r *Response) IsError() bool {
	return r.StatusCode > 299
}

// Status returns the response status as a string.
//
func (r *Response) Status() string {
	var b strings.Builder
	if r != nil {
		b.WriteString(strconv.Itoa(r.StatusCode))
		b.WriteString(" ")
		b.WriteString(http.StatusText(r.StatusCode))
	}
	return b.String()
}

// Warnings returns the deprecation warnings from response headers.
//
func (r *Response) Warnings() []string {
	return r.Header["Warning"]
}

// HasWarnings returns true when the response headers contain deprecation warnings.
//
func (r *Response) HasWarnings() bool {
	return len(r.Warnings()) > 0
}

// Request is the basic type for encapsulating API request
//
type Request struct {
	Context   context.Context
	Transport Transport
}

// NewRequest creates an HTTP request.
//
func NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, path, body)
}
