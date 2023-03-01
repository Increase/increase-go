package options

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"increase/core"
	"increase/core/query"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

func getNormalizedOS() string {
	switch runtime.GOOS {
	case "ios":
		return "iOS"
	case "android":
		return "Android"
	case "darwin":
		return "MacOS"
	case "window":
		return "Windows"
	case "freebsd":
		return "FreeBSD"
	case "openbsd":
		return "OpenBSD"
	case "linux":
		return "Linux"
	default:
		return fmt.Sprintf("Other:%s", runtime.GOOS)
	}
}

func getNormalizedArchitecture() string {
	switch runtime.GOARCH {
	case "386":
		return "x32"
	case "amd64":
		return "x64"
	case "arm":
		return "arm"
	case "arm64":
		return "arm64"
	default:
		return fmt.Sprintf("other:%s", runtime.GOARCH)
	}
}

func getPlatformProperties() map[string]string {
	return map[string]string{
		"X-Stainless-Lang":            "go",
		"X-Stainless-Package-Version": "unknown",
		"X-Stainless-OS":              getNormalizedOS(),
		"X-Stainless-Arch":            getNormalizedArchitecture(),
		"X-Stainless-Runtime":         "go",
		"X-Stainless-Runtime-Version": runtime.Version(),
	}
}

func NewRequestConfig(ctx context.Context, method string, u string, body interface{}, dst interface{}, opts ...RequestOption) (*RequestConfig, error) {
	var reader io.ReadCloser
	if body, ok := body.(json.Marshaler); ok {
		b, err := body.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if len(b) != 0 {
			reader = io.NopCloser(bytes.NewBuffer(b))
		}
	}
	if body, ok := body.(query.Queryer); ok {
		u = u + "?" + body.URLQuery().Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return nil, err
	}
	if reader != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	for k, v := range getPlatformProperties() {
		req.Header.Add(k, v)
	}
	cfg := RequestConfig{
		MaxRetries: 0,
		Context:    ctx,
		Request:    req,
		HTTPClient: http.DefaultClient,
	}
	cfg.ResponseBodyInto = dst
	cfg.Apply(opts...)
	return &cfg, nil
}

func ExecuteNewRequest(ctx context.Context, method string, u string, body interface{}, dst interface{}, opts ...RequestOption) error {
	cfg, err := NewRequestConfig(ctx, method, u, body, dst, opts...)
	if err != nil {
		return err
	}
	return cfg.Execute()
}

type RequestConfig struct {
	MaxRetries int
	Context    context.Context
	Request    *http.Request
	BaseURL    *url.URL
	HTTPClient *http.Client
	APIKey     string
	// If ResponseBodyInto not nil, then we will attempt to deserialize into
	// ResponseBodyInto. If Destination is a []byte, then it will return the body as
	// is.
	ResponseBodyInto interface{}
	// ResponseInto copies the \*http.Response of the corresponding request into the
	// given address
	ResponseInto **http.Response
}

func (cfg *RequestConfig) Execute() error {
	u, err := cfg.BaseURL.Parse(cfg.Request.URL.String())
	if err != nil {
		return err
	}
	cfg.Request.URL = u
	res, err := cfg.HTTPClient.Do(cfg.Request)
	if err != nil {
		return core.RequestError{Cause: err, Request: cfg.Request, Response: res}
	}
	if res.StatusCode > 299 {
		// TODO appropriately retry
		return core.NewAPIErrorFromResponse(cfg.Request, res)
	}
	if cfg.ResponseInto != nil {
		*cfg.ResponseInto = res
	}

	if cfg.ResponseBodyInto == nil {
		return nil
	}

	contents, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// If we are not json return plaintext
	isJSON := strings.Contains(res.Header.Get("content-type"), "application/json")
	if !isJSON {
		switch dst := cfg.ResponseBodyInto.(type) {
		case *string:
			*dst = string(contents)
		case **string:
			tmp := string(contents)
			*dst = &tmp
		case *[]byte:
			*dst = contents
		default:
			return fmt.Errorf("expected destination type of 'string' or '[]byte' for responses with content-type that is not 'application/json'")
		}
		return nil
	}

	err = json.NewDecoder(bytes.NewReader(contents)).Decode(cfg.ResponseBodyInto)
	if err != nil {
		err = fmt.Errorf("error parsing response json: %w", err)
	}

	return nil
}

func (cfg *RequestConfig) Clone(ctx context.Context) *RequestConfig {
	if cfg == nil {
		return nil
	}
	return &RequestConfig{
		MaxRetries: cfg.MaxRetries,
		Context:    ctx,
		Request:    cfg.Request.Clone(ctx),
		HTTPClient: cfg.HTTPClient,
	}
}

func (cfg *RequestConfig) Apply(opts ...RequestOption) {
	for _, opt := range opts {
		opt(cfg)
	}
}

type RequestOption = func(*RequestConfig)

func WithBaseURL(base string) RequestOption {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatalf("failed to parse BaseURL: %s\n", err)
	}
	return func(r *RequestConfig) {
		r.BaseURL = u
	}
}

func WithHTTPClient(client *http.Client) RequestOption {
	return func(r *RequestConfig) {
		r.HTTPClient = client
	}
}

func WithMaxRetries(retries int) RequestOption {
	return func(r *RequestConfig) {
		r.MaxRetries = retries
	}
}

func WithHeader(key, value string) RequestOption {
	return func(r *RequestConfig) {
		r.Request.Header[key] = []string{value}
	}
}
func WithHeaderAdd(key, value string) RequestOption {
	return func(r *RequestConfig) {
		r.Request.Header[key] = append(r.Request.Header[key], value)
	}
}
func WithHeaderDel(key string) RequestOption {
	return func(r *RequestConfig) {
		delete(r.Request.Header, key)
	}
}

func WithQuery(key, value string) RequestOption {
	return func(r *RequestConfig) {
		query := r.Request.URL.Query()
		query.Set(key, value)
		r.Request.URL.RawQuery = query.Encode()
	}
}
func WithQueryAdd(key, value string) RequestOption {
	return func(r *RequestConfig) {
		query := r.Request.URL.Query()
		query.Add(key, value)
		r.Request.URL.RawQuery = query.Encode()
	}
}
func WithQueryDel(key string) RequestOption {
	return func(r *RequestConfig) {
		query := r.Request.URL.Query()
		query.Del(key)
		r.Request.URL.RawQuery = query.Encode()
	}
}

func WithResponseBodyInto(dst any) RequestOption {
	return func(r *RequestConfig) {
		r.ResponseBodyInto = dst
	}
}

func WithResponseInto(dst **http.Response) RequestOption {
	return func(r *RequestConfig) {
		r.ResponseInto = dst
	}
}

func WithAPIKey(key string) RequestOption {
	return func(r *RequestConfig) {
		r.APIKey = key
		r.Apply(WithHeader("Authorization", fmt.Sprintf("Bearer %s", r.APIKey)))
	}
}

func WithEnvironmentProduction() RequestOption {
	return WithBaseURL("https://api.increase.com")
}

func WithEnvironmentSandbox() RequestOption {
	return WithBaseURL("https://sandbox.increase.com")
}
