package requestconfig

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/increase/increase-go/internal"
	"github.com/increase/increase-go/internal/apierror"
	"github.com/increase/increase-go/internal/apiform"
	"github.com/increase/increase-go/internal/apiquery"
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
		"X-Stainless-Package-Version": internal.PackageVersion,
		"X-Stainless-OS":              getNormalizedOS(),
		"X-Stainless-Arch":            getNormalizedArchitecture(),
		"X-Stainless-Runtime":         "go",
		"X-Stainless-Runtime-Version": runtime.Version(),
	}
}

func NewRequestConfig(ctx context.Context, method string, u string, body interface{}, dst interface{}, opts ...func(*RequestConfig) error) (*RequestConfig, error) {
	var b []byte
	contentType := "application/json"
	if body, ok := body.(json.Marshaler); ok {
		var err error
		b, err = body.MarshalJSON()
		if err != nil {
			return nil, err
		}
	}
	if body, ok := body.(apiform.Marshaler); ok {
		var err error
		b, err = body.MarshalMultipart()
		if err != nil {
			return nil, err
		}
		contentType = "multipart/form-data"
	}
	if body, ok := body.(apiquery.Queryer); ok {
		u = u + "?" + body.URLQuery().Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, u, nil)
	if err != nil {
		return nil, err
	}
	if b != nil {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Idempotency-Key", "stainless-go-"+uuid.New().String())
	req.Header.Set("Accept", "application/json")

	for k, v := range getPlatformProperties() {
		req.Header.Add(k, v)
	}
	cfg := RequestConfig{
		MaxRetries: 2,
		Context:    ctx,
		Request:    req,
		HTTPClient: http.DefaultClient,
		Buffer:     b,
	}
	cfg.ResponseBodyInto = dst
	err = cfg.Apply(opts...)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// RequestConfigruation represents all the state related to one request.
//
// Editing the variables inside RequestConfig directly is unstable api. Prefer
// composing [func(*RequestConfig) error] instead if possible.
type RequestConfig struct {
	MaxRetries     int
	RequestTimeout time.Duration
	Context        context.Context
	Request        *http.Request
	BaseURL        *url.URL
	HTTPClient     *http.Client
	APIKey         string
	// If ResponseBodyInto not nil, then we will attempt to deserialize into
	// ResponseBodyInto. If Destination is a []byte, then it will return the body as
	// is.
	ResponseBodyInto interface{}
	// ResponseInto copies the \*http.Response of the corresponding request into the
	// given address
	ResponseInto **http.Response
	Buffer       []byte
}

func (cfg *RequestConfig) Execute() error {
	u, err := cfg.BaseURL.Parse(cfg.Request.URL.String())
	if err != nil {
		return err
	}
	cfg.Request.URL = u

	if len(cfg.Buffer) != 0 && cfg.Request.Body == nil {
		buf := bytes.NewReader(cfg.Buffer)
		cfg.Request.ContentLength = int64(len(cfg.Buffer))
		cfg.Request.Body = io.NopCloser(buf)
		cfg.Request.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(cfg.Buffer)), nil }
	}

	var res *http.Response
	for i := 0; i <= cfg.MaxRetries; i += 1 {
		ctx := cfg.Request.Context()
		if cfg.RequestTimeout != time.Duration(0) {
			nctx, cancel := context.WithTimeout(ctx, cfg.RequestTimeout)
			ctx = nctx
			defer cancel()
		}
		res, err = cfg.HTTPClient.Do(cfg.Request.Clone(ctx))

		if i == cfg.MaxRetries || err == nil && res.StatusCode != http.StatusConflict && res.StatusCode != http.StatusTooManyRequests && res.StatusCode < http.StatusInternalServerError {
			break
		}

		duration := time.Duration(500) * time.Millisecond * time.Duration(math.Exp(float64(i)))
		if res != nil {
			if parsed, err := strconv.ParseInt(res.Header.Get("Retry-After"), 10, 64); err == nil {
				duration = time.Duration(parsed) * time.Second
			}
		}
		if duration > time.Duration(60)*time.Second {
			duration = time.Duration(60) * time.Second
		}
		duration += time.Millisecond * time.Duration(-500+rand.Intn(1000))
		time.Sleep(duration)
	}

	if err != nil {
		return err
	}

	if res.StatusCode > 399 {
		aerr := apierror.Error{Request: cfg.Request, Response: res, StatusCode: res.StatusCode}
		contents, err := io.ReadAll(res.Body)
		err = aerr.UnmarshalJSON(contents)
		if err != nil {
			return err
		}
		return &aerr
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

func ExecuteNewRequest(ctx context.Context, method string, u string, body interface{}, dst interface{}, opts ...func(*RequestConfig) error) error {
	cfg, err := NewRequestConfig(ctx, method, u, body, dst, opts...)
	if err != nil {
		return err
	}
	return cfg.Execute()
}

func (cfg *RequestConfig) Clone(ctx context.Context) *RequestConfig {
	if cfg == nil {
		return nil
	}
	req := cfg.Request.Clone(ctx)
	var err error
	if req.Body != nil {
		req.Body, err = req.GetBody()
	}
	if err != nil {
		return nil
	}
	new := &RequestConfig{
		MaxRetries: cfg.MaxRetries,
		Context:    ctx,
		Request:    req,
		HTTPClient: cfg.HTTPClient,
	}
	new.Request.Header.Set("Idempotency-Key", "stainless-go-"+uuid.New().String())
	return new
}

func (cfg *RequestConfig) Apply(opts ...func(*RequestConfig) error) error {
	for _, opt := range opts {
		err := opt(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}
