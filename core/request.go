package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"increase/core/query"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// https://github.com/mautrix/go/blob/master/event/content.go#L180-L195
func mergeMaps[P any](into, from map[string]P) {
	for key, newValue := range from {
		existingValue, ok := into[key]
		if !ok {
			into[key] = newValue
			continue
		}
		existingValueMap, okEx := any(existingValue).(map[string]P)
		newValueMap, okNew := any(newValue).(map[string]P)
		if okEx && okNew {
			mergeMaps(existingValueMap, newValueMap)
		} else {
			into[key] = newValue
		}
	}
}

func copyMap[P any](original map[string]P) (copied map[string]P) {
	copied = make(map[string]P)
	for key, value := range original {
		copied[key] = value
	}
	return
}

// https://github.com/mautrix/go/blob/master/event/content.go#L120-L158
func marshalMerged(known interface{}, unknown map[string]interface{}) (bytes []byte, err error) {
	if known == nil && len(unknown) == 0 {
		return nil, nil
	}

	var knownParsed map[string]interface{}

	if known != nil {
		bytes, err = json.Marshal(known)
		if err != nil {
			return
		}
		if len(unknown) == 0 {
			return
		}

		err = json.Unmarshal(bytes, &knownParsed)
		if err != nil {
			return
		}
	}

	if unknown != nil {
		if knownParsed != nil {
			unknown = copyMap(unknown)
			mergeMaps(unknown, knownParsed)
		}
		bytes, err = json.Marshal(unknown)
	}

	return
}

func (r *RequestOpts) MarshalBody(body interface{}) ([]byte, error) {
	return marshalMerged(body, r.ExtraBody)
}

func (r *RequestOpts) MarshalQuery(query interface{}) ([]byte, error) {
	return nil, nil
}

func (r *RequestOpts) MarshalHeaders(headers http.Header) http.Header {
	if headers == nil {
		// headers = make(http.Header)
	}

	if r.Headers == nil || len(r.Headers) == 0 {
		return headers
	}

	for key, value := range r.Headers {
		headers.Set(key, value)
	}

	return headers
}

type RequestConfiguration struct {
}

type RequestOpts struct {
	URL           string
	Method        string
	Path          string
	ExtraQuery    url.Values
	ExtraBody     map[string]interface{}
	Headers       map[string]string
	Timeout       time.Duration
	MaxRetries    int
	QuerySettings *query.QuerySettings

	ResponseInto           **http.Response
	ResponseBodyInto       interface{}
	DeserializeReturnValue *bool
}

func MergeRequestOpts(opts ...*RequestOpts) (merged RequestOpts) {
	switch len(opts) {
	case 0:
		return
	case 1:
		if opt := opts[0]; opt != nil {
			merged = *opt
		}
		return
	default:
		merged = RequestOpts{
			ExtraBody:  make(map[string]interface{}),
			ExtraQuery: make(url.Values),
			Headers:    make(map[string]string),
		}
		for _, opts := range opts {
			mergeMaps(merged.ExtraBody, opts.ExtraBody)
			mergeMaps(merged.ExtraQuery, opts.ExtraQuery)
			mergeMaps(merged.Headers, opts.Headers)
			if opts.Timeout != 0 {
				merged.Timeout = opts.Timeout
			}
			if opts.MaxRetries != 0 {
				merged.MaxRetries = opts.MaxRetries
			}
			if len(opts.URL) != 0 {
				merged.URL = opts.URL
			}
			if opts.ResponseInto != nil {
				merged.ResponseInto = opts.ResponseInto
			}
			if opts.ResponseBodyInto != nil {
				merged.ResponseBodyInto = opts.ResponseBodyInto
			}
		}
		return
	}
}

type CoreRequest struct {
	Params RequestOpts

	URL   string
	Query interface{}

	Body      interface{}
	BodyBytes []byte

	Request *http.Request
	Agent   Agent
}

func (cr *CoreRequest) LoadDefaults(c *CoreClient) {
	if cr.Params.Timeout == 0 {
		cr.Params.Timeout = c.DefaultClient.Timeout
	}
	if cr.Params.MaxRetries == 0 {
		cr.Params.MaxRetries = c.MaxRetries
	}
	if cr.Params.QuerySettings == nil {
		cr.Params.QuerySettings = &c.QuerySettings
	}
}

func (cr *CoreRequest) MarshalBody() ([]byte, error) {
	return cr.Params.MarshalBody(cr.Body)
}

func (cr *CoreRequest) InitializeHTTPRequest(ctx context.Context) *http.Request {
	req := cr.Request
	var err error
	if req == nil {
		if ctx == nil {
			ctx = context.TODO()
		}
		req, err = http.NewRequestWithContext(ctx, strings.ToUpper(cr.Params.Method), cr.Params.URL, nil)
		cr.Request = req
	}
	if err != nil {
		panic(err)
	}

	return req
}

func (cr *CoreRequest) getBodyBytes() ([]byte, error) {
	if len(cr.BodyBytes) != 0 {
		return cr.BodyBytes, nil
	} else {
		return cr.MarshalBody()
	}
}

func MergeURLValues(values ...url.Values) (merged url.Values) {
	merged = make(url.Values)
	add := func(v url.Values) {
		if v == nil {
			return
		}
		for key, values := range v {
			for _, value := range values {
				merged.Add(key, value)
			}
		}
	}
	for _, v := range values {
		add(v)
	}
	return
}

func joinPath(base string, elem ...string) (string, error) {
	if joined, err := url.JoinPath(base, elem...); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("/%s", joined), nil
	}
}

func (cr *CoreRequest) FillRequestURL(req *http.Request, baseURL string) error {
	compileQuery := func(existingValues url.Values) string {
		var merged url.Values
		if q := cr.Query; q != nil {
			merged = query.MarshalWithSettings(q, *cr.Params.QuerySettings)
		}
		return MergeURLValues(existingValues, merged, cr.Params.ExtraQuery).Encode()
	}
	if len(cr.Params.URL) != 0 {
		if parsed, err := url.Parse(cr.Params.URL); err == nil {
			parsed.RawQuery = compileQuery(parsed.Query())
			if len(parsed.Host) == 0 {
				if parsedBase, err := url.Parse(baseURL); err == nil {
					parsed.Scheme = parsedBase.Scheme
					parsed.Host = parsedBase.Host
					parsed.Path, _ = joinPath(parsedBase.Path, parsed.Path)
				}
			}
			req.URL = parsed
			return nil
		} else {
			return err
		}
	}
	if parsed, err := url.Parse(cr.Params.Path); err == nil && len(parsed.Scheme) != 0 && len(parsed.Host) != 0 {
		// provided path is an explicit url
		parsed.RawQuery = compileQuery(parsed.Query())
		req.URL = parsed
	} else if joined, err := url.JoinPath(baseURL, cr.Params.Path); err == nil {
		if parsed, err = url.Parse(joined); err == nil {
			parsed.RawQuery = compileQuery(parsed.Query())
			req.URL = parsed
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (cr *CoreRequest) FillRequestBody(req *http.Request) error {
	if body, err := cr.getBodyBytes(); err != nil {
		return err
	} else if len(body) != 0 {
		req.Body = io.NopCloser(bytes.NewReader(body))
		req.ContentLength = int64(len(body))
	}
	return nil
}

func (cr *CoreRequest) FillRequestHeaders() {
	cr.Request.Header = cr.Params.MarshalHeaders(cr.Request.Header)
}
