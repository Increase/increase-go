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

// ProofOfAuthorizationRequestService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProofOfAuthorizationRequestService] method instead.
type ProofOfAuthorizationRequestService struct {
	Options []option.RequestOption
}

// NewProofOfAuthorizationRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProofOfAuthorizationRequestService(opts ...option.RequestOption) (r *ProofOfAuthorizationRequestService) {
	r = &ProofOfAuthorizationRequestService{}
	r.Options = opts
	return
}

// Retrieve a Proof of Authorization Request
func (r *ProofOfAuthorizationRequestService) Get(ctx context.Context, proofOfAuthorizationRequestID string, opts ...option.RequestOption) (res *ProofOfAuthorizationRequest, err error) {
	opts = append(r.Options[:], opts...)
	if proofOfAuthorizationRequestID == "" {
		err = errors.New("missing required proof_of_authorization_request_id parameter")
		return
	}
	path := fmt.Sprintf("proof_of_authorization_requests/%s", proofOfAuthorizationRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Proof of Authorization Requests
func (r *ProofOfAuthorizationRequestService) List(ctx context.Context, query ProofOfAuthorizationRequestListParams, opts ...option.RequestOption) (res *pagination.Page[ProofOfAuthorizationRequest], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "proof_of_authorization_requests"
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

// List Proof of Authorization Requests
func (r *ProofOfAuthorizationRequestService) ListAutoPaging(ctx context.Context, query ProofOfAuthorizationRequestListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ProofOfAuthorizationRequest] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// A request for proof of authorization for one or more ACH debit transfers.
type ProofOfAuthorizationRequest struct {
	// The Proof of Authorization Request identifier.
	ID string `json:"id,required"`
	// The ACH Transfers associated with the request.
	ACHTransfers []ProofOfAuthorizationRequestACHTransfer `json:"ach_transfers,required"`
	// The time the Proof of Authorization Request was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The time the Proof of Authorization Request is due.
	DueOn time.Time `json:"due_on,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `proof_of_authorization_request`.
	Type ProofOfAuthorizationRequestType `json:"type,required"`
	// The time the Proof of Authorization Request was last updated.
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	JSON      proofOfAuthorizationRequestJSON `json:"-"`
}

// proofOfAuthorizationRequestJSON contains the JSON metadata for the struct
// [ProofOfAuthorizationRequest]
type proofOfAuthorizationRequestJSON struct {
	ID           apijson.Field
	ACHTransfers apijson.Field
	CreatedAt    apijson.Field
	DueOn        apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProofOfAuthorizationRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proofOfAuthorizationRequestJSON) RawJSON() string {
	return r.raw
}

type ProofOfAuthorizationRequestACHTransfer struct {
	// The ACH Transfer identifier.
	ID   string                                     `json:"id,required"`
	JSON proofOfAuthorizationRequestACHTransferJSON `json:"-"`
}

// proofOfAuthorizationRequestACHTransferJSON contains the JSON metadata for the
// struct [ProofOfAuthorizationRequestACHTransfer]
type proofOfAuthorizationRequestACHTransferJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProofOfAuthorizationRequestACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proofOfAuthorizationRequestACHTransferJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `proof_of_authorization_request`.
type ProofOfAuthorizationRequestType string

const (
	ProofOfAuthorizationRequestTypeProofOfAuthorizationRequest ProofOfAuthorizationRequestType = "proof_of_authorization_request"
)

func (r ProofOfAuthorizationRequestType) IsKnown() bool {
	switch r {
	case ProofOfAuthorizationRequestTypeProofOfAuthorizationRequest:
		return true
	}
	return false
}

type ProofOfAuthorizationRequestListParams struct {
	CreatedAt param.Field[ProofOfAuthorizationRequestListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ProofOfAuthorizationRequestListParams]'s query parameters
// as `url.Values`.
func (r ProofOfAuthorizationRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ProofOfAuthorizationRequestListParamsCreatedAt struct {
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

// URLQuery serializes [ProofOfAuthorizationRequestListParamsCreatedAt]'s query
// parameters as `url.Values`.
func (r ProofOfAuthorizationRequestListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
