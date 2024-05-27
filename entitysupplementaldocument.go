// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/pagination"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// EntitySupplementalDocumentService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitySupplementalDocumentService] method instead.
type EntitySupplementalDocumentService struct {
	Options []option.RequestOption
}

// NewEntitySupplementalDocumentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEntitySupplementalDocumentService(opts ...option.RequestOption) (r *EntitySupplementalDocumentService) {
	r = &EntitySupplementalDocumentService{}
	r.Options = opts
	return
}

// Create a supplemental document for an Entity
func (r *EntitySupplementalDocumentService) New(ctx context.Context, entityID string, body EntitySupplementalDocumentNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/supplemental_documents", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Entity Supplemental Document Submissions
func (r *EntitySupplementalDocumentService) List(ctx context.Context, query EntitySupplementalDocumentListParams, opts ...option.RequestOption) (res *pagination.Page[SupplementalDocument], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *EntitySupplementalDocumentService) ListAutoPaging(ctx context.Context, query EntitySupplementalDocumentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[SupplementalDocument] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Supplemental Documents are uploaded files connected to an Entity during
// onboarding.
type SupplementalDocument struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Supplemental Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The File containing the document.
	FileID string `json:"file_id,required"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `entity_supplemental_document`.
	Type SupplementalDocumentType `json:"type,required"`
	JSON supplementalDocumentJSON `json:"-"`
}

// supplementalDocumentJSON contains the JSON metadata for the struct
// [SupplementalDocument]
type supplementalDocumentJSON struct {
	CreatedAt      apijson.Field
	FileID         apijson.Field
	IdempotencyKey apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SupplementalDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r supplementalDocumentJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `entity_supplemental_document`.
type SupplementalDocumentType string

const (
	SupplementalDocumentTypeEntitySupplementalDocument SupplementalDocumentType = "entity_supplemental_document"
)

func (r SupplementalDocumentType) IsKnown() bool {
	switch r {
	case SupplementalDocumentTypeEntitySupplementalDocument:
		return true
	}
	return false
}

type EntitySupplementalDocumentNewParams struct {
	// The identifier of the File containing the document.
	FileID param.Field[string] `json:"file_id,required"`
}

func (r EntitySupplementalDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntitySupplementalDocumentListParams struct {
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

// URLQuery serializes [EntitySupplementalDocumentListParams]'s query parameters as
// `url.Values`.
func (r EntitySupplementalDocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
