package core

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type CoreClient struct {
	DefaultClient http.Client
	Agent         Agent

	BaseURL    string
	MaxRetries int

	AuthHeaders func() map[string]string
}

func (c *CoreClient) UserAgent() string {
	return "increase/Go 0.0.0"
}

func (c *CoreClient) agent() Agent {
	if c.Agent != nil {
		return c.Agent
	} else {
		return &c.DefaultClient
	}
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
		for key, value := range c.AuthHeaders() {
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

func (c *CoreClient) Request(path string, cr *CoreRequest, body interface{}) error {
	cr.Path = path
	cr.LoadDefaults(c)

	req := cr.InitializeHTTPRequest()

	c.FillDefaultHeaders(req)

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

	log.Printf("%s", res.Header.Get("content-type"))
	if strings.Contains(res.Header.Get("content-type"), "application/json") {
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return err
		}
	} else {
		if text, err := extractResponseText(res); err != nil {
			return err
		} else {
			body = text
		}
	}
	return nil
}

func (c *CoreClient) Get(path string, req *CoreRequest, res interface{}) error {
	req.Method = "Get"
	return c.Request(path, req, &res)
}
func (c *CoreClient) Post(path string, req *CoreRequest, res interface{}) error {
	req.Method = "Post"
	return c.Request(path, req, &res)
}
func (c *CoreClient) Patch(path string, req *CoreRequest, res interface{}) error {
	req.Method = "Patch"
	return c.Request(path, req, &res)
}
func (c *CoreClient) Put(path string, req *CoreRequest, res interface{}) error {
	req.Method = "Put"
	return c.Request(path, req, &res)
}
func (c *CoreClient) Delete(path string, req *CoreRequest, res interface{}) error {
	req.Method = "Delete"
	return c.Request(path, req, &res)
}
