package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RequestError struct {
	Cause    error
	Request  *http.Request
	Response *http.Response
}

func (e RequestError) Error() string {
	return fmt.Errorf("%s: %d: %w", e.Request.URL.String(), e.Response.StatusCode, e.Cause).Error()
}

type APIError interface {
	Status() int
	ErrorBody() interface{}
	Message() string
	Headers() http.Header
	Error() string
}

type GenericAPIError struct {
	request   *CoreRequest
	response  *http.Response
	status    int
	errorBody interface{}
	message   string
	headers   http.Header
}

func (e GenericAPIError) Status() int {
	return e.status
}

func (e GenericAPIError) ErrorBody() interface{} {
	return e.errorBody
}

func (e GenericAPIError) Message() string {
	return e.message
}

func (e GenericAPIError) Headers() http.Header {
	return e.headers
}

func (e GenericAPIError) Error() string {
	return fmt.Sprintf("%s %s: %d", e.request.Request.Method, e.request.Request.URL.String(), e.status)
}

type BadRequestError struct{ GenericAPIError }
type AuthenticationError struct{ GenericAPIError }
type PermissionDeniedError struct{ GenericAPIError }
type NotFoundError struct{ GenericAPIError }
type ConflictError struct{ GenericAPIError }
type UnprocessableEntityError struct{ GenericAPIError }
type RateLimitError struct{ GenericAPIError }
type InternalServerError struct{ GenericAPIError }

func NewAPIError(req *CoreRequest, res *http.Response, status int, error interface{}, message string, headers http.Header) APIError {
	err := GenericAPIError{req, res, status, error, message, headers}

	if status == 400 {
		return BadRequestError{err}
	}
	if status == 401 {
		return AuthenticationError{err}
	}
	if status == 403 {
		return PermissionDeniedError{err}
	}
	if status == 404 {
		return NotFoundError{err}
	}
	if status == 409 {
		return ConflictError{err}
	}
	if status == 422 {
		return UnprocessableEntityError{err}
	}
	if status == 429 {
		return RateLimitError{err}
	}
	if status > 500 {
		return InternalServerError{err}
	}

	return err
}

func extractResponseText(res *http.Response) (string, error) {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, res.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func NewAPIErrorFromResponse(req *CoreRequest, res *http.Response) APIError {
	var errorJSON interface{}
	var message string

	if text, err := extractResponseText(res); err == nil {
		jsonDecoder := json.NewDecoder(strings.NewReader(text))
		_ = jsonDecoder.Decode(&errorJSON)
		if errorJSON != nil {
			message = ""
		} else {
			message = text
		}
	}

	return NewAPIError(req, res, res.StatusCode, errorJSON, message, res.Header)
}
