package pagination

import "strconv"
import "net/url"
import "increase/core"
import "context"
import "net/http"

type PageResponseInterface[Model interface{}] interface {
	GetItems() []Model
	GetItem(index int) *Model
	GetLength() int
}

type PageInterface[Model any, Response PageResponseInterface[Model]] interface {
	// different per endpoint
	GetResponse() Response
	GetRawResponse() *http.Response
	HasNextPage() bool
	GetNextPageGeneric() (PageInterface[Model, Response], error)
	GetNextPageParams() *PageParams
	// shared
	Current() *Model // and/or maybe have custom, eg. .Customer()?
	Next() bool
	Err() error
	Index() int
}

func getInt(response *http.Response, key string, defaultValue int) int {
	if response == nil {
		return defaultValue
	}
	if request := response.Request; request == nil {
		return defaultValue
	} else if u := request.URL; u == nil {
		return defaultValue
	} else {
		p := u.Query()
		if rawInt, ok := p[key]; ok && len(rawInt) != 0 {
			if i, err := strconv.Atoi(rawInt[0]); err == nil {
				return i
			}
		}
	}
	return defaultValue
}

type PageOptions struct {
	RequestParams interface{}
	Params        url.Values
	Path          string
	URL           string
}

type PageParams struct {
	Params  url.Values
	Headers map[string]string
	URL     string
}

//
type PageResponse[Model interface{}] struct {
	Data *[]Model `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

func (r *PageResponse[Model]) GetData() (Data []Model) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *PageResponse[Model]) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

var _ PageResponseInterface[interface{}] = (*PageResponse[interface{}])(nil)

func (r *PageResponse[Model]) GetItems() []Model {
	return r.GetData()
}

func (r *PageResponse[Model]) GetItem(index int) *Model {
	return &r.GetItems()[index]
}

func (r *PageResponse[Model]) GetLength() int {
	return len(r.GetItems())
}

type Page[Model interface{}] struct {
	Options        PageOptions
	Requester      core.Requester
	Context        context.Context
	requestOptions *core.RequestOpts
	index          int
	runningIndex   int
	fired          bool
	err            error
	rawResponse    *http.Response
	response       *PageResponse[Model]
}

var _ PageInterface[interface{}, *PageResponse[interface{}]] = (*Page[interface{}])(nil)

func (r *Page[Model]) GetNextPageParams() *PageParams {
	if cursor := r.response.GetNextCursor(); len(cursor) == 0 {
		if r.fired {
			return nil
		} else {
			return &PageParams{}
		}
	} else {
		return &PageParams{
			Params: url.Values{
				"cursor": {cursor},
			},
		}
	}
}

func (r *Page[Model]) GetResponse() *PageResponse[Model] {
	return r.response
}

func (r *Page[Model]) GetRawResponse() *http.Response {
	return r.rawResponse
}

func (r *Page[Model]) GetResponseHeader(key string) string {
	if response := r.rawResponse; response == nil {
		return ""
	} else if header := response.Header; header == nil {
		return ""
	} else if values, ok := header[key]; !ok || len(values) == 0 {
		return ""
	} else {
		return values[0]
	}
}

func (r *Page[Model]) getPageItems() []Model {
	return r.response.GetItems()
}

func (r *Page[Model]) paramsFromURL(str string) url.Values {
	params := r.Options.Params
	if params == nil {
		params = url.Values{}
	}
	if u, err := url.Parse(str); err == nil {
		for key, values := range u.Query() {
			for _, value := range values {
				params.Add(key, value)
			}
		}
	}
	return params
}

func (r *Page[Model]) infoToOptions(info *PageParams) core.RequestOpts {
	var options core.RequestOpts
	if r.requestOptions != nil {
		options = *r.requestOptions
	}
	rawURL := core.CoalesceStrings(info.URL, r.Options.Path)
	if len(rawURL) != 0 {
		params := r.paramsFromURL(rawURL)
		newURL, _ := url.Parse(rawURL)
		query := core.MergeURLValues(newURL.Query(), params, info.Params)
		newURL.RawQuery = query.Encode()
		options.Headers = info.Headers
		options.URL = newURL.String()
	}
	return options
}

func (r *Page[Model]) HasNextPage() bool {
	if items := r.getPageItems(); len(items) == 0 {
		return false
	}
	return r.GetNextPageParams() != nil
}

func (r *Page[Model]) getRawNextPage(response **http.Response) (res *PageResponse[Model], err error) {
	if info := r.GetNextPageParams(); info == nil {
		return nil, nil
	} else {
		opts := r.infoToOptions(info)
		opts.ResponseInto = response
		err = r.Requester.Request(r.Context, "", &core.CoreRequest{
			Params: opts,
			Query:  r.Options.RequestParams,
		}, &res)
		r.fired = true
	}
	return
}

func (r *Page[Model]) GetNextPage() (page *Page[Model], err error) {
	var response *http.Response
	rawPage, err := r.getRawNextPage(&response)
	if err != nil {
		return nil, err
	}
	newPage := &Page[Model]{
		rawResponse:  response,
		response:     rawPage,
		runningIndex: r.runningIndex,
		Options:      r.Options,
		Requester:    r.Requester,
		fired:        r.fired,
		Context:      r.Context,
	}
	return newPage, nil
}

func (r *Page[Model]) GetNextPageGeneric() (PageInterface[Model, *PageResponse[Model]], error) {
	return r.GetNextPage()
}

func (r *Page[Model]) Next() bool {
	if r.err != nil {
		return false
	}
	if r.index >= r.response.GetLength() {
		if res, err := r.GetNextPage(); err != nil {
			r.err = err
			return false
		} else if res == nil || res.response == nil {
			return false
		} else if res.response.GetLength() == 0 {
			return false
		} else {
			*r = *res
			return r.Next()
		}
	}
	r.runningIndex += 1
	r.index += 1
	return true
}

// Current returns the element corresponding to the current index. Calling Current
// before calling Next is a programming error and will cause your program to crash.
func (r *Page[Model]) Current() *Model {
	// The index is subtracted by one to differentiate between 0 (unset) and 1 (first element)
	// Calling Current before calling Next is a programming error
	return r.response.GetItem(r.index - 1)
}

func (r *Page[Model]) Err() error {
	return r.err
}

func (r *Page[Model]) Index() int {
	return r.runningIndex - 1
}
