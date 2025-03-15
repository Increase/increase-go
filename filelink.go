// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
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
	opts = append(r.Options[:], opts...)
	path := "file_links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a File Link
func (r *FileLinkService) Get(ctx context.Context, fileLinkID string, opts ...option.RequestOption) (res *FileLink, err error) {
	opts = append(r.Options[:], opts...)
	if fileLinkID == "" {
		err = errors.New("missing required file_link_id parameter")
		return
	}
	path := fmt.Sprintf("file_links/%s", fileLinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List File Links
func (r *FileLinkService) List(ctx context.Context, query FileLinkListParams, opts ...option.RequestOption) (res *pagination.Page[FileLink], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "file_links"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List File Links
func (r *FileLinkService) ListAutoPaging(ctx context.Context, query FileLinkListParams, opts ...option.RequestOption) *pagination.PageAutoPager[FileLink] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Normally Files can only be downloaded via the API using your Increase API key.
// File Links let you generate signed URLs for Files that can be used to download
// the File without an Increase API key.
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
	// A URL where the File can be downloaded. The URL will expire after the
	// `expires_at` time. This URL is unauthenticated and can be used to download the
	// File without an Increase API key.
	PublicDownloadURL string `json:"public_download_url,required"`
	// A constant representing the object's type. For this resource it will always be
	// `file_link`.
	Type FileLinkType `json:"type,required"`
	JSON fileLinkJSON `json:"-"`
}

// fileLinkJSON contains the JSON metadata for the struct [FileLink]
type fileLinkJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	ExpiresAt         apijson.Field
	FileID            apijson.Field
	IdempotencyKey    apijson.Field
	PublicDownloadURL apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
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
	// of the request. The maxiumum is 1 day from the time of the request.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
}

func (r FileLinkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileLinkListParams struct {
	// The identifier of the File to list File Links for.
	FileID    param.Field[string]                      `query:"file_id,required"`
	CreatedAt param.Field[FileLinkListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [FileLinkListParams]'s query parameters as `url.Values`.
func (r FileLinkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FileLinkListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [FileLinkListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r FileLinkListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
