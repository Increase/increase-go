// File generated from our OpenAPI spec by Stainless.

package apierror

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/increase/increase-go/internal/apijson"
)

// TODO
type Error struct {
	Type   ErrorType   `json:"type,required"`
	Title  string      `json:"title,required"`
	Detail string      `json:"detail,required,nullable"`
	Status ErrorStatus `json:"status,required"`
	// All errors related to parsing the request parameters.
	Errors     []interface{} `json:"errors,required"`
	RetryAfter int64         `json:"retry_after,nullable"`
	JSON       errorJSON
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// errorJSON contains the JSON metadata for the struct [Error]
type errorJSON struct {
	Type        apijson.Field
	Title       apijson.Field
	Detail      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
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

type ErrorType string

const (
	ErrorTypeRateLimitedError ErrorType = "rate_limited_error"
)

type ErrorStatus int64

const (
	ErrorStatus429 ErrorStatus = 429
)
