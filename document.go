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

// DocumentService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Retrieve a Document
func (r *DocumentService) Get(ctx context.Context, documentID string, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return
	}
	path := fmt.Sprintf("documents/%s", documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Documents
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *pagination.Page[Document], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "documents"
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

// List Documents
func (r *DocumentService) ListAutoPaging(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Document] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Increase generates certain documents / forms automatically for your application;
// they can be listed here. Currently the only supported document type is IRS Form
// 1099-INT.
type Document struct {
	// The Document identifier.
	ID string `json:"id,required"`
	// The type of document.
	Category DocumentCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Entity the document was generated for.
	EntityID string `json:"entity_id,required,nullable"`
	// The identifier of the File containing the Document's contents.
	FileID string `json:"file_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `document`.
	Type DocumentType `json:"type,required"`
	JSON documentJSON `json:"-"`
}

// documentJSON contains the JSON metadata for the struct [Document]
type documentJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	CreatedAt   apijson.Field
	EntityID    apijson.Field
	FileID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentJSON) RawJSON() string {
	return r.raw
}

// The type of document.
type DocumentCategory string

const (
	// Internal Revenue Service Form 1099-INT.
	DocumentCategoryForm1099Int DocumentCategory = "form_1099_int"
	// A document submitted in response to a proof of authorization request for an ACH
	// transfer.
	DocumentCategoryProofOfAuthorization DocumentCategory = "proof_of_authorization"
	// Company information, such a policies or procedures, typically submitted during
	// our due diligence process.
	DocumentCategoryCompanyInformation DocumentCategory = "company_information"
)

func (r DocumentCategory) IsKnown() bool {
	switch r {
	case DocumentCategoryForm1099Int, DocumentCategoryProofOfAuthorization, DocumentCategoryCompanyInformation:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `document`.
type DocumentType string

const (
	DocumentTypeDocument DocumentType = "document"
)

func (r DocumentType) IsKnown() bool {
	switch r {
	case DocumentTypeDocument:
		return true
	}
	return false
}

type DocumentListParams struct {
	Category  param.Field[DocumentListParamsCategory]  `query:"category"`
	CreatedAt param.Field[DocumentListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID param.Field[string] `query:"entity_id"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DocumentListParamsCategory struct {
	// Filter Documents for those with the specified category or categories. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]DocumentListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes [DocumentListParamsCategory]'s query parameters as
// `url.Values`.
func (r DocumentListParamsCategory) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DocumentListParamsCategoryIn string

const (
	// Internal Revenue Service Form 1099-INT.
	DocumentListParamsCategoryInForm1099Int DocumentListParamsCategoryIn = "form_1099_int"
	// A document submitted in response to a proof of authorization request for an ACH
	// transfer.
	DocumentListParamsCategoryInProofOfAuthorization DocumentListParamsCategoryIn = "proof_of_authorization"
	// Company information, such a policies or procedures, typically submitted during
	// our due diligence process.
	DocumentListParamsCategoryInCompanyInformation DocumentListParamsCategoryIn = "company_information"
)

func (r DocumentListParamsCategoryIn) IsKnown() bool {
	switch r {
	case DocumentListParamsCategoryInForm1099Int, DocumentListParamsCategoryInProofOfAuthorization, DocumentListParamsCategoryInCompanyInformation:
		return true
	}
	return false
}

type DocumentListParamsCreatedAt struct {
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

// URLQuery serializes [DocumentListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r DocumentListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
