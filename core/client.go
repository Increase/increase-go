package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"increase/core/query"
	"io"
	"net/http"
	"strings"
)

type CoreClient struct {
	DefaultClient http.Client
	Agent         Agent

	BaseURL    string
	MaxRetries int

	QuerySettings query.QuerySettings

	AuthHeaders map[string]string
}

var _ Requester = (*CoreClient)(nil)

func (c *CoreClient) UserAgent() string {
	return "increase/Go 0.0.0"
}

func (c *CoreClient) agent() Agent {
	if c.Agent != nil {
		return c.Agent
	}
	return &c.DefaultClient
}

func (c *CoreClient) Do(req *http.Request) (*http.Response, error) {
	return c.agent().Do(req)
}

func (c *CoreClient) DefaultHeaders() map[string]string {
	defaultHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"User-Agent":   c.UserAgent(),
	}
	for key, value := range getPlatformProperties() {
		defaultHeaders[key] = value
	}
	if c.AuthHeaders != nil {
		for key, value := range c.AuthHeaders {
			defaultHeaders[key] = value
		}
	}
	return defaultHeaders
}

func (c *CoreClient) FillDefaultHeaders(req *http.Request) {
	if req.Header == nil {
		req.Header = http.Header{}
	}

	defaultHeaders := c.DefaultHeaders()
	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}
}

func (c *CoreClient) Request(ctx context.Context, path string, cr *CoreRequest, responseBody interface{}) error {
	cr.Params.Path = CoalesceStrings(cr.Params.Path, path)
	cr.LoadDefaults(c)

	req := cr.InitializeHTTPRequest(ctx)

	c.FillDefaultHeaders(req)

	cr.Params.URL = fmt.Sprintf("%s%s", c.BaseURL, path)
	cr.FillRequestURL(req, c.BaseURL)
	cr.FillRequestBody(req)
	cr.FillRequestHeaders()

	res, err := c.Do(req)

	if err != nil {
		// if c.canRetry(cr) {
		// 	return c.retry(ctx, cr)
		// }
		return RequestError{Cause: err, Request: req, Response: res}
	}

	if res.StatusCode > 299 {
		// if c.canRetry(cr) {
		// 	return c.retry(ctx, cr)
		// }
		return NewAPIErrorFromResponse(cr, res)
	}

	if cr.Params.ResponseInto != nil {
		*cr.Params.ResponseInto = res
	}

	isJSON := strings.Contains(res.Header.Get("content-type"), "application/json")
	// If we are not json return plaintext
	if !isJSON {
		text, err := extractResponseText(res)
		if err != nil {
			return fmt.Errorf("error extracting response text: %w", err)
		}
		responseBody = text
		return nil
	}

	// Load the body into bytes
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	parseJSON := func(v any) error {
		err := json.NewDecoder(io.NopCloser(bytes.NewReader(contents))).Decode(&v)
		if err != nil {
			err = fmt.Errorf("error parsing response json: %w", err)
		}
		return err
	}
	if deserialize := cr.Params.DeserializeReturnValue; deserialize == nil || *deserialize {
		if err := parseJSON(&responseBody); err != nil {
			return err
		}
	}
	if cr.Params.ResponseBodyInto != nil {
		if err := parseJSON(&cr.Params.ResponseBodyInto); err != nil {
			return err
		}
	}

	return nil
}

func (c *CoreClient) Get(ctx context.Context, path string, req *CoreRequest, res interface{}) error {
	req.Params.Method = http.MethodGet
	return c.Request(ctx, path, req, &res)
}
func (c *CoreClient) Post(ctx context.Context, path string, req *CoreRequest, res interface{}) error {
	req.Params.Method = http.MethodPost
	return c.Request(ctx, path, req, &res)
}
func (c *CoreClient) Patch(ctx context.Context, path string, req *CoreRequest, res interface{}) error {
	req.Params.Method = http.MethodPatch
	return c.Request(ctx, path, req, &res)
}
func (c *CoreClient) Put(ctx context.Context, path string, req *CoreRequest, res interface{}) error {
	req.Params.Method = http.MethodPut
	return c.Request(ctx, path, req, &res)
}
func (c *CoreClient) Delete(ctx context.Context, path string, req *CoreRequest, res interface{}) error {
	req.Params.Method = http.MethodDelete
	return c.Request(ctx, path, req, &res)
}
