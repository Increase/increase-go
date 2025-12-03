// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// SupplementalDocumentService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupplementalDocumentService] method instead.
type SupplementalDocumentService struct {
	Options []option.RequestOption
}

// NewSupplementalDocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSupplementalDocumentService(opts ...option.RequestOption) (r *SupplementalDocumentService) {
	r = &SupplementalDocumentService{}
	r.Options = opts
	return
}

// Create a supplemental document for an Entity
func (r *SupplementalDocumentService) New(ctx context.Context, body SupplementalDocumentNewParams, opts ...option.RequestOption) (res *EntitySupplementalDocument, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "entity_supplemental_documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Entity Supplemental Document Submissions
func (r *SupplementalDocumentService) List(ctx context.Context, query SupplementalDocumentListParams, opts ...option.RequestOption) (res *pagination.Page[EntitySupplementalDocument], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "entity_supplemental_documents"
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

// List Entity Supplemental Document Submissions
func (r *SupplementalDocumentService) ListAutoPaging(ctx context.Context, query SupplementalDocumentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[EntitySupplementalDocument] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Supplemental Documents are uploaded files connected to an Entity during
// onboarding.
type EntitySupplementalDocument struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Supplemental Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The Entity the supplemental document is attached to.
	EntityID string `json:"entity_id,required"`
	// The File containing the document.
	FileID string `json:"file_id,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `entity_supplemental_document`.
	Type EntitySupplementalDocumentType `json:"type,required"`
	JSON entitySupplementalDocumentJSON `json:"-"`
}

// entitySupplementalDocumentJSON contains the JSON metadata for the struct
// [EntitySupplementalDocument]
type entitySupplementalDocumentJSON struct {
	CreatedAt      apijson.Field
	EntityID       apijson.Field
	FileID         apijson.Field
	IdempotencyKey apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntitySupplementalDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitySupplementalDocumentJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `entity_supplemental_document`.
type EntitySupplementalDocumentType string

const (
	EntitySupplementalDocumentTypeEntitySupplementalDocument EntitySupplementalDocumentType = "entity_supplemental_document"
)

func (r EntitySupplementalDocumentType) IsKnown() bool {
	switch r {
	case EntitySupplementalDocumentTypeEntitySupplementalDocument:
		return true
	}
	return false
}

type SupplementalDocumentNewParams struct {
	// The identifier of the Entity to associate with the supplemental document.
	EntityID param.Field[string] `json:"entity_id,required"`
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r SupplementalDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SupplementalDocumentListParams struct {
	// The identifier of the Entity to list supplemental documents for.
	EntityID param.Field[string] `query:"entity_id,required"`
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

// URLQuery serializes [SupplementalDocumentListParams]'s query parameters as
// `url.Values`.
func (r SupplementalDocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
