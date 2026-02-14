// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// FileLinkService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileLinkService] method instead.
type FileLinkService struct {
	Options []option.RequestOption
}

// NewFileLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileLinkService(opts ...option.RequestOption) (r *FileLinkService) {
	r = &FileLinkService{}
	r.Options = opts
	return
}

// Create a File Link
func (r *FileLinkService) New(ctx context.Context, body FileLinkNewParams, opts ...option.RequestOption) (res *FileLink, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "file_links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// File Links let you generate a URL that can be used to download a File.
type FileLink struct {
	// The File Link identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the File
	// Link was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the File
	// Link will expire.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The identifier of the File the File Link points to.
	FileID string `json:"file_id,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `file_link`.
	Type FileLinkType `json:"type,required"`
	// A URL where the File can be downloaded. The URL will expire after the
	// `expires_at` time. This URL is unauthenticated and can be used to download the
	// File without an Increase API key.
	UnauthenticatedURL string       `json:"unauthenticated_url,required"`
	JSON               fileLinkJSON `json:"-"`
}

// fileLinkJSON contains the JSON metadata for the struct [FileLink]
type fileLinkJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	ExpiresAt          apijson.Field
	FileID             apijson.Field
	IdempotencyKey     apijson.Field
	Type               apijson.Field
	UnauthenticatedURL apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *FileLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileLinkJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `file_link`.
type FileLinkType string

const (
	FileLinkTypeFileLink FileLinkType = "file_link"
)

func (r FileLinkType) IsKnown() bool {
	switch r {
	case FileLinkTypeFileLink:
		return true
	}
	return false
}

type FileLinkNewParams struct {
	// The File to create a File Link for.
	FileID param.Field[string] `json:"file_id,required"`
	// The time at which the File Link will expire. The default is 1 hour from the time
	// of the request. The maximum is 1 day from the time of the request.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
}

func (r FileLinkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
