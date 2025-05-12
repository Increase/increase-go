// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/Increase/increase-go/internal/apijson"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	Detail string `json:"detail,required,nullable"`
	// All errors related to parsing the request parameters.
	Errors     []interface{} `json:"errors,required"`
	Reason     ErrorReason   `json:"reason,required"`
	ResourceID string        `json:"resource_id,required"`
	Status     ErrorStatus   `json:"status,required"`
	Title      string        `json:"title,required"`
	Type       ErrorType     `json:"type,required"`
	RetryAfter int64         `json:"retry_after,nullable"`
	JSON       errorJSON     `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// errorJSON contains the JSON metadata for the struct [Error]
type errorJSON struct {
	Detail      apijson.Field
	Errors      apijson.Field
	Reason      apijson.Field
	ResourceID  apijson.Field
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

func (r errorJSON) RawJSON() string {
	return r.raw
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.RawJSON())
}

func (r *Error) DumpRequest(body bool) []byte {
	if r.Request.GetBody != nil {
		r.Request.Body, _ = r.Request.GetBody()
	}
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}

type ErrorReason string

const (
	ErrorReasonDeletedCredential ErrorReason = "deleted_credential"
	ErrorReasonExpiredCredential ErrorReason = "expired_credential"
	ErrorReasonNoCredential      ErrorReason = "no_credential"
	ErrorReasonNoHeader          ErrorReason = "no_header"
	ErrorReasonNoAPIAccess       ErrorReason = "no_api_access"
	ErrorReasonWrongEnvironment  ErrorReason = "wrong_environment"
)

func (r ErrorReason) IsKnown() bool {
	switch r {
	case ErrorReasonDeletedCredential, ErrorReasonExpiredCredential, ErrorReasonNoCredential, ErrorReasonNoHeader, ErrorReasonNoAPIAccess, ErrorReasonWrongEnvironment:
		return true
	}
	return false
}

type ErrorStatus int64

const (
	ErrorStatus429 ErrorStatus = 429
	ErrorStatus403 ErrorStatus = 403
	ErrorStatus404 ErrorStatus = 404
	ErrorStatus400 ErrorStatus = 400
	ErrorStatus409 ErrorStatus = 409
	ErrorStatus401 ErrorStatus = 401
	ErrorStatus500 ErrorStatus = 500
)

func (r ErrorStatus) IsKnown() bool {
	switch r {
	case ErrorStatus429, ErrorStatus403, ErrorStatus404, ErrorStatus400, ErrorStatus409, ErrorStatus401, ErrorStatus500:
		return true
	}
	return false
}

type ErrorType string

const (
	ErrorTypeRateLimitedError               ErrorType = "rate_limited_error"
	ErrorTypePrivateFeatureError            ErrorType = "private_feature_error"
	ErrorTypeObjectNotFoundError            ErrorType = "object_not_found_error"
	ErrorTypeMalformedRequestError          ErrorType = "malformed_request_error"
	ErrorTypeInvalidParametersError         ErrorType = "invalid_parameters_error"
	ErrorTypeInvalidOperationError          ErrorType = "invalid_operation_error"
	ErrorTypeInvalidAPIKeyError             ErrorType = "invalid_api_key_error"
	ErrorTypeInternalServerError            ErrorType = "internal_server_error"
	ErrorTypeInsufficientPermissionsError   ErrorType = "insufficient_permissions_error"
	ErrorTypeIdempotencyKeyAlreadyUsedError ErrorType = "idempotency_key_already_used_error"
	ErrorTypeEnvironmentMismatchError       ErrorType = "environment_mismatch_error"
	ErrorTypeAPIMethodNotFoundError         ErrorType = "api_method_not_found_error"
)

func (r ErrorType) IsKnown() bool {
	switch r {
	case ErrorTypeRateLimitedError, ErrorTypePrivateFeatureError, ErrorTypeObjectNotFoundError, ErrorTypeMalformedRequestError, ErrorTypeInvalidParametersError, ErrorTypeInvalidOperationError, ErrorTypeInvalidAPIKeyError, ErrorTypeInternalServerError, ErrorTypeInsufficientPermissionsError, ErrorTypeIdempotencyKeyAlreadyUsedError, ErrorTypeEnvironmentMismatchError, ErrorTypeAPIMethodNotFoundError:
		return true
	}
	return false
}
