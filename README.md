# Increase Go API Library

<a href="https://pkg.go.dev/github.com/Increase/increase-go"><img src="https://pkg.go.dev/badge/github.com/Increase/increase-go.svg" alt="Go Reference"></a>

The Increase Go library provides convenient access to the [Increase REST API](https://increase.com/documentation)
from applications written in Go.

## Installation

<!-- x-release-please-start-version -->

```go
import (
	"github.com/Increase/increase-go" // imported as increase
)
```

<!-- x-release-please-end -->

Or to pin the version:

<!-- x-release-please-start-version -->

```sh
go get -u 'github.com/Increase/increase-go@v0.266.0'
```

<!-- x-release-please-end -->

## Requirements

This library requires Go 1.18+.

## Usage

The full API of this library can be found in [api.md](api.md).

```go
package main

import (
	"context"
	"fmt"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/option"
)

func main() {
	client := increase.NewClient(
		option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("INCREASE_API_KEY")
		option.WithEnvironmentSandbox(), // defaults to option.WithEnvironmentProduction()
	)
	account, err := client.Accounts.New(context.TODO(), increase.AccountNewParams{
		Name:      increase.F("New Account!"),
		EntityID:  increase.F("entity_n8y8tnk2p9339ti393yi"),
		ProgramID: increase.F("program_i2v2os4mwza1oetokh9i"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", account.ID)
}

```

### Request fields

All request parameters are wrapped in a generic `Field` type,
which we use to distinguish zero values from null or omitted fields.

This prevents accidentally sending a zero value if you forget a required parameter,
and enables explicitly sending `null`, `false`, `''`, or `0` on optional parameters.
Any field not specified is not sent.

To construct fields with values, use the helpers `String()`, `Int()`, `Float()`, or most commonly, the generic `F[T]()`.
To send a null, use `Null[T]()`, and to send a nonconforming value, use `Raw[T](any)`. For example:

```go
params := FooParams{
	Name: increase.F("hello"),

	// Explicitly send `"description": null`
	Description: increase.Null[string](),

	Point: increase.F(increase.Point{
		X: increase.Int(0),
		Y: increase.Int(1),

		// In cases where the API specifies a given type,
		// but you want to send something else, use `Raw`:
		Z: increase.Raw[int64](0.01), // sends a float
	}),
}
```

### Response objects

All fields in response structs are value types (not pointers or wrappers).

If a given field is `null`, not present, or invalid, the corresponding field
will simply be its zero value.

All response structs also include a special `JSON` field, containing more detailed
information about each property, which you can use like so:

```go
if res.Name == "" {
	// true if `"name"` is either not present or explicitly null
	res.JSON.Name.IsNull()

	// true if the `"name"` key was not present in the response JSON at all
	res.JSON.Name.IsMissing()

	// When the API returns data that cannot be coerced to the expected type:
	if res.JSON.Name.IsInvalid() {
		raw := res.JSON.Name.Raw()

		legacyName := struct{
			First string `json:"first"`
			Last  string `json:"last"`
		}{}
		json.Unmarshal([]byte(raw), &legacyName)
		name = legacyName.First + " " + legacyName.Last
	}
}
```

These `.JSON` structs also include an `Extras` map containing
any properties in the json response that were not specified
in the struct. This can be useful for API features not yet
present in the SDK.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := increase.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.Accounts.New(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

See the [full list of request options](https://pkg.go.dev/github.com/Increase/increase-go/option).

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all pages:

```go
iter := client.Accounts.ListAutoPaging(context.TODO(), increase.AccountListParams{})
// Automatically fetches more pages as needed.
for iter.Next() {
	account := iter.Current()
	fmt.Printf("%+v\n", account)
}
if err := iter.Err(); err != nil {
	panic(err.Error())
}
```

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

```go
page, err := client.Accounts.List(context.TODO(), increase.AccountListParams{})
for page != nil {
	for _, account := range page.Data {
		fmt.Printf("%+v\n", account)
	}
	page, err = page.GetNextPage()
}
if err != nil {
	panic(err.Error())
}
```

### Errors

When the API returns a non-success status code, we return an error with type
`*increase.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.Accounts.New(context.TODO(), increase.AccountNewParams{
	Name: increase.F("New Account!"),
})
if err != nil {
	var apierr *increase.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response
	}
	panic(err.Error()) // GET "/accounts": 400 Bad Request { ... }
}
```

When other errors occur, they are returned unwrapped; for example,
if HTTP transport fails, you might receive `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a timeout for a request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not start over.
To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.Accounts.New(
	ctx,
	increase.AccountNewParams{
		Name:      increase.F("New Account!"),
		EntityID:  increase.F("entity_n8y8tnk2p9339ti393yi"),
		ProgramID: increase.F("program_i2v2os4mwza1oetokh9i"),
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

### File uploads

Request parameters that correspond to file uploads in multipart requests are typed as
`param.Field[io.Reader]`. The contents of the `io.Reader` will by default be sent as a multipart form
part with the file name of "anonymous_file" and content-type of "application/octet-stream".

The file name and content-type can be customized by implementing `Name() string` or `ContentType()
string` on the run-time type of `io.Reader`. Note that `os.File` implements `Name() string`, so a
file returned by `os.Open` will be sent with the file name on disk.

We also provide a helper `increase.FileParam(reader io.Reader, filename string, contentType string)`
which can be used to wrap any `io.Reader` with the appropriate file name and content type.

```go
// A file from the file system
file, err := os.Open("my/file.txt")
increase.FileNewParams{
	File:    increase.F[io.Reader](file),
	Purpose: increase.F(increase.FileNewParamsPurposeCheckImageFront),
}

// A file from a string
increase.FileNewParams{
	File:    increase.F[io.Reader](strings.NewReader("my file contents")),
	Purpose: increase.F(increase.FileNewParamsPurposeCheckImageFront),
}

// With a custom filename and contentType
increase.FileNewParams{
	File:    increase.FileParam(strings.NewReader(`{"hello": "foo"}`), "file.go", "application/json"),
	Purpose: increase.F(increase.FileNewParamsPurposeCheckImageFront),
}
```

## Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
We retry by default all connection errors, 408 Request Timeout, 409 Conflict, 429 Rate Limit,
and >=500 Internal errors.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := increase.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.Accounts.New(
	context.TODO(),
	increase.AccountNewParams{
		Name:      increase.F("New Account!"),
		EntityID:  increase.F("entity_n8y8tnk2p9339ti393yi"),
		ProgramID: increase.F("program_i2v2os4mwza1oetokh9i"),
	},
	option.WithMaxRetries(5),
)
```

### Middleware

We provide `option.WithMiddleware` which applies the given
middleware to requests.

```go
func Logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	// Before the request
	start := time.Now()
	LogReq(req)

	// Forward the request to the next handler
	res, err = next(req)

	// Handle stuff after the request
	end := time.Now()
	LogRes(res, err, start - end)

    return res, err
}

client := increase.NewClient(
	option.WithMiddleware(Logger),
)
```

When multiple middlewares are provided as variadic arguments, the middlewares
are applied left to right. If `option.WithMiddleware` is given
multiple times, for example first in the client then the method, the
middleware in the client will run first and the middleware given in the method
will run next.

You may also replace the default `http.Client` with
`option.WithHTTPClient(client)`. Only one http client is
accepted (this overwrites any previous client) and receives requests after any
middleware has been applied.

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/Increase/increase-go/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
