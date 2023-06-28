// File generated from our OpenAPI spec by Stainless.

package apierror

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/increase/increase-go/internal/apijson"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	Detail string `json:"detail,required,nullable"`
	// All errors related to parsing the request parameters.
	Errors     []interface{} `json:"errors,required"`
	Status     ErrorStatus   `json:"status,required"`
	Title      string        `json:"title,required"`
	Type       ErrorType     `json:"type,required"`
	RetryAfter int64         `json:"retry_after,nullable"`
	JSON       errorJSON
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// errorJSON contains the JSON metadata for the struct [Error]
type errorJSON struct {
	Detail      apijson.Field
	Errors      apijson.Field
	Status      apijson.Field
	Title       apijson.Field
	Type        apijson.Field
	RetryAfter  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	body, _ := io.ReadAll(r.Response.Body)
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), string(body))
}

func (r *Error) DumpRequest(body bool) []byte {
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}

type ErrorStatus int64

const (
	ErrorStatus429 ErrorStatus = 429
)

type ErrorType string

const (
	ErrorTypeRateLimitedError ErrorType = "rate_limited_error"
)
