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

// Create a Document
func (r *DocumentService) New(ctx context.Context, body DocumentNewParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	path := "documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
// they can be listed here.
type Document struct {
	// The Document identifier.
	ID string `json:"id,required"`
	// Properties of an account verification letter document.
	AccountVerificationLetter DocumentAccountVerificationLetter `json:"account_verification_letter,required,nullable"`
	// The type of document.
	Category DocumentCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Entity the document was generated for.
	EntityID string `json:"entity_id,required,nullable"`
	// The identifier of the File containing the Document's contents.
	FileID string `json:"file_id,required"`
	// Properties of a funding instructions document.
	FundingInstructions DocumentFundingInstructions `json:"funding_instructions,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `document`.
	Type DocumentType `json:"type,required"`
	JSON documentJSON `json:"-"`
}

// documentJSON contains the JSON metadata for the struct [Document]
type documentJSON struct {
	ID                        apijson.Field
	AccountVerificationLetter apijson.Field
	Category                  apijson.Field
	CreatedAt                 apijson.Field
	EntityID                  apijson.Field
	FileID                    apijson.Field
	FundingInstructions       apijson.Field
	IdempotencyKey            apijson.Field
	Type                      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentJSON) RawJSON() string {
	return r.raw
}

// Properties of an account verification letter document.
type DocumentAccountVerificationLetter struct {
	// The identifier of the Account Number the document was generated for.
	AccountNumberID string                                `json:"account_number_id,required"`
	JSON            documentAccountVerificationLetterJSON `json:"-"`
}

// documentAccountVerificationLetterJSON contains the JSON metadata for the struct
// [DocumentAccountVerificationLetter]
type documentAccountVerificationLetterJSON struct {
	AccountNumberID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DocumentAccountVerificationLetter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentAccountVerificationLetterJSON) RawJSON() string {
	return r.raw
}

// The type of document.
type DocumentCategory string

const (
	DocumentCategoryForm1099Int               DocumentCategory = "form_1099_int"
	DocumentCategoryForm1099Misc              DocumentCategory = "form_1099_misc"
	DocumentCategoryProofOfAuthorization      DocumentCategory = "proof_of_authorization"
	DocumentCategoryCompanyInformation        DocumentCategory = "company_information"
	DocumentCategoryAccountVerificationLetter DocumentCategory = "account_verification_letter"
	DocumentCategoryFundingInstructions       DocumentCategory = "funding_instructions"
)

func (r DocumentCategory) IsKnown() bool {
	switch r {
	case DocumentCategoryForm1099Int, DocumentCategoryForm1099Misc, DocumentCategoryProofOfAuthorization, DocumentCategoryCompanyInformation, DocumentCategoryAccountVerificationLetter, DocumentCategoryFundingInstructions:
		return true
	}
	return false
}

// Properties of a funding instructions document.
type DocumentFundingInstructions struct {
	// The identifier of the Account Number the document was generated for.
	AccountNumberID string                          `json:"account_number_id,required"`
	JSON            documentFundingInstructionsJSON `json:"-"`
}

// documentFundingInstructionsJSON contains the JSON metadata for the struct
// [DocumentFundingInstructions]
type documentFundingInstructionsJSON struct {
	AccountNumberID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DocumentFundingInstructions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentFundingInstructionsJSON) RawJSON() string {
	return r.raw
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

type DocumentNewParams struct {
	// The type of document to create.
	Category param.Field[DocumentNewParamsCategory] `json:"category,required"`
	// An account verification letter. Required if and only if `category` is
	// `account_verification_letter`.
	AccountVerificationLetter param.Field[DocumentNewParamsAccountVerificationLetter] `json:"account_verification_letter"`
	// Funding instructions. Required if and only if `category` is
	// `funding_instructions`.
	FundingInstructions param.Field[DocumentNewParamsFundingInstructions] `json:"funding_instructions"`
}

func (r DocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of document to create.
type DocumentNewParamsCategory string

const (
	DocumentNewParamsCategoryAccountVerificationLetter DocumentNewParamsCategory = "account_verification_letter"
	DocumentNewParamsCategoryFundingInstructions       DocumentNewParamsCategory = "funding_instructions"
)

func (r DocumentNewParamsCategory) IsKnown() bool {
	switch r {
	case DocumentNewParamsCategoryAccountVerificationLetter, DocumentNewParamsCategoryFundingInstructions:
		return true
	}
	return false
}

// An account verification letter. Required if and only if `category` is
// `account_verification_letter`.
type DocumentNewParamsAccountVerificationLetter struct {
	// The Account Number the bank letter should be generated for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// If provided, the letter will include the Account's balance as of the date.
	BalanceDate param.Field[time.Time] `json:"balance_date" format:"date"`
}

func (r DocumentNewParamsAccountVerificationLetter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Funding instructions. Required if and only if `category` is
// `funding_instructions`.
type DocumentNewParamsFundingInstructions struct {
	// The Account Number the funding instructions should be generated for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
}

func (r DocumentNewParamsFundingInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DocumentListParams struct {
	Category  param.Field[DocumentListParamsCategory]  `query:"category"`
	CreatedAt param.Field[DocumentListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID param.Field[string] `query:"entity_id"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
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
	DocumentListParamsCategoryInForm1099Int               DocumentListParamsCategoryIn = "form_1099_int"
	DocumentListParamsCategoryInForm1099Misc              DocumentListParamsCategoryIn = "form_1099_misc"
	DocumentListParamsCategoryInProofOfAuthorization      DocumentListParamsCategoryIn = "proof_of_authorization"
	DocumentListParamsCategoryInCompanyInformation        DocumentListParamsCategoryIn = "company_information"
	DocumentListParamsCategoryInAccountVerificationLetter DocumentListParamsCategoryIn = "account_verification_letter"
	DocumentListParamsCategoryInFundingInstructions       DocumentListParamsCategoryIn = "funding_instructions"
)

func (r DocumentListParamsCategoryIn) IsKnown() bool {
	switch r {
	case DocumentListParamsCategoryInForm1099Int, DocumentListParamsCategoryInForm1099Misc, DocumentListParamsCategoryInProofOfAuthorization, DocumentListParamsCategoryInCompanyInformation, DocumentListParamsCategoryInAccountVerificationLetter, DocumentListParamsCategoryInFundingInstructions:
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
