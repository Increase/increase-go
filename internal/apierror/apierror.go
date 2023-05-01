package apierror

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/increase/increase-go/internal/apijson"
)

type Error struct {
	Type   ErrorType   `json:"type,required"`
	Title  string      `json:"title,required"`
	Detail string      `json:"detail,required,nullable"`
	Status ErrorStatus `json:"status,required"`
	// All errors related to parsing the request parameters.
	Errors     []interface{} `json:"errors,required"`
	RetryAfter int64         `json:"retry_after,nullable"`
	JSON       ErrorJSON
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

type ErrorJSON struct {
	Type       apijson.Metadata
	Title      apijson.Metadata
	Detail     apijson.Metadata
	Status     apijson.Metadata
	Errors     apijson.Metadata
	RetryAfter apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Error using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
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
	ErrorTypeRateLimitedError              ErrorType = "rate_limited_error"
	ErrorTypePrivateFeatureError           ErrorType = "private_feature_error"
	ErrorTypeObjectNotFoundError           ErrorType = "object_not_found_error"
	ErrorTypeMalformedRequestError         ErrorType = "malformed_request_error"
	ErrorTypeInvalidParametersError        ErrorType = "invalid_parameters_error"
	ErrorTypeInvalidOperationError         ErrorType = "invalid_operation_error"
	ErrorTypeInvalidAPIKeyError            ErrorType = "invalid_api_key_error"
	ErrorTypeInternalServerError           ErrorType = "internal_server_error"
	ErrorTypeInsufficientPermissionsError  ErrorType = "insufficient_permissions_error"
	ErrorTypeIdempotencyUnprocessableError ErrorType = "idempotency_unprocessable_error"
	ErrorTypeIdempotencyConflictError      ErrorType = "idempotency_conflict_error"
	ErrorTypeEnvironmentMismatchError      ErrorType = "environment_mismatch_error"
	ErrorTypeAPIMethodNotFoundError        ErrorType = "api_method_not_found_error"
)

type ErrorStatus int64

const (
	ErrorStatus400 ErrorStatus = 400
	ErrorStatus401 ErrorStatus = 401
	ErrorStatus403 ErrorStatus = 403
	ErrorStatus404 ErrorStatus = 404
	ErrorStatus409 ErrorStatus = 409
	ErrorStatus422 ErrorStatus = 422
	ErrorStatus429 ErrorStatus = 429
	ErrorStatus500 ErrorStatus = 500
)
