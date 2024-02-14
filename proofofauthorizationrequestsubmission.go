// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// ProofOfAuthorizationRequestSubmissionService contains methods and other services
// that help with interacting with the increase API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewProofOfAuthorizationRequestSubmissionService] method instead.
type ProofOfAuthorizationRequestSubmissionService struct {
	Options []option.RequestOption
}

// NewProofOfAuthorizationRequestSubmissionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewProofOfAuthorizationRequestSubmissionService(opts ...option.RequestOption) (r *ProofOfAuthorizationRequestSubmissionService) {
	r = &ProofOfAuthorizationRequestSubmissionService{}
	r.Options = opts
	return
}

// Submit Proof of Authorization
func (r *ProofOfAuthorizationRequestSubmissionService) New(ctx context.Context, body ProofOfAuthorizationRequestSubmissionNewParams, opts ...option.RequestOption) (res *ProofOfAuthorizationRequestSubmission, err error) {
	opts = append(r.Options[:], opts...)
	path := "proof_of_authorization_request_submissions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Proof of Authorization Request Submission
func (r *ProofOfAuthorizationRequestSubmissionService) Get(ctx context.Context, proofOfAuthorizationRequestSubmissionID string, opts ...option.RequestOption) (res *ProofOfAuthorizationRequestSubmission, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("proof_of_authorization_request_submissions/%s", proofOfAuthorizationRequestSubmissionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Proof of Authorization Request Submissions
func (r *ProofOfAuthorizationRequestSubmissionService) List(ctx context.Context, query ProofOfAuthorizationRequestSubmissionListParams, opts ...option.RequestOption) (res *shared.Page[ProofOfAuthorizationRequestSubmission], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "proof_of_authorization_request_submissions"
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

// List Proof of Authorization Request Submissions
func (r *ProofOfAuthorizationRequestSubmissionService) ListAutoPaging(ctx context.Context, query ProofOfAuthorizationRequestSubmissionListParams, opts ...option.RequestOption) *shared.PageAutoPager[ProofOfAuthorizationRequestSubmission] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Information submitted in response to a proof of authorization request. Per
// Nacha's guidance on proof of authorization, the originator must ensure that the
// authorization complies with applicable legal requirements, is readily
// identifiable as an authorization, and has clear and readily understandable
// terms.
type ProofOfAuthorizationRequestSubmission struct {
	// The Proof of Authorization Request Submission identifier.
	ID string `json:"id,required"`
	// Terms of authorization.
	AuthorizationTerms string `json:"authorization_terms,required"`
	// Time of authorization.
	AuthorizedAt time.Time `json:"authorized_at,required" format:"date-time"`
	// Company of the authorizer.
	AuthorizerCompany string `json:"authorizer_company,required,nullable"`
	// Email of the authorizer.
	AuthorizerEmail string `json:"authorizer_email,required,nullable"`
	// IP address of the authorizer.
	AuthorizerIPAddress string `json:"authorizer_ip_address,required,nullable"`
	// Name of the authorizer.
	AuthorizerName string `json:"authorizer_name,required,nullable"`
	// The time the Proof of Authorization Request Submission was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// ID of the proof of authorization request.
	ProofOfAuthorizationRequestID string `json:"proof_of_authorization_request_id,required"`
	// Status of the proof of authorization request submission.
	Status ProofOfAuthorizationRequestSubmissionStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `proof_of_authorization_request_submission`.
	Type ProofOfAuthorizationRequestSubmissionType `json:"type,required"`
	// The time the Proof of Authorization Request Submission was last updated.
	UpdatedAt time.Time                                 `json:"updated_at,required" format:"date-time"`
	JSON      proofOfAuthorizationRequestSubmissionJSON `json:"-"`
}

// proofOfAuthorizationRequestSubmissionJSON contains the JSON metadata for the
// struct [ProofOfAuthorizationRequestSubmission]
type proofOfAuthorizationRequestSubmissionJSON struct {
	ID                            apijson.Field
	AuthorizationTerms            apijson.Field
	AuthorizedAt                  apijson.Field
	AuthorizerCompany             apijson.Field
	AuthorizerEmail               apijson.Field
	AuthorizerIPAddress           apijson.Field
	AuthorizerName                apijson.Field
	CreatedAt                     apijson.Field
	IdempotencyKey                apijson.Field
	ProofOfAuthorizationRequestID apijson.Field
	Status                        apijson.Field
	Type                          apijson.Field
	UpdatedAt                     apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *ProofOfAuthorizationRequestSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the proof of authorization request submission.
type ProofOfAuthorizationRequestSubmissionStatus string

const (
	// The proof of authorization request submission is pending review.
	ProofOfAuthorizationRequestSubmissionStatusPendingReview ProofOfAuthorizationRequestSubmissionStatus = "pending_review"
	// The proof of authorization request submission was rejected.
	ProofOfAuthorizationRequestSubmissionStatusRejected ProofOfAuthorizationRequestSubmissionStatus = "rejected"
	// The proof of authorization request submission is pending sending.
	ProofOfAuthorizationRequestSubmissionStatusPendingSending ProofOfAuthorizationRequestSubmissionStatus = "pending_sending"
	// The proof of authorization request submission was sent.
	ProofOfAuthorizationRequestSubmissionStatusSent ProofOfAuthorizationRequestSubmissionStatus = "sent"
)

// A constant representing the object's type. For this resource it will always be
// `proof_of_authorization_request_submission`.
type ProofOfAuthorizationRequestSubmissionType string

const (
	ProofOfAuthorizationRequestSubmissionTypeProofOfAuthorizationRequestSubmission ProofOfAuthorizationRequestSubmissionType = "proof_of_authorization_request_submission"
)

type ProofOfAuthorizationRequestSubmissionNewParams struct {
	// Terms of authorization.
	AuthorizationTerms param.Field[string] `json:"authorization_terms,required"`
	// Time of authorization.
	AuthorizedAt param.Field[time.Time] `json:"authorized_at,required" format:"date-time"`
	// Email of the authorizer.
	AuthorizerEmail param.Field[string] `json:"authorizer_email,required"`
	// Name of the authorizer.
	AuthorizerName param.Field[string] `json:"authorizer_name,required"`
	// ID of the proof of authorization request.
	ProofOfAuthorizationRequestID param.Field[string] `json:"proof_of_authorization_request_id,required"`
	// Company of the authorizer.
	AuthorizerCompany param.Field[string] `json:"authorizer_company"`
	// IP address of the authorizer.
	AuthorizerIPAddress param.Field[string] `json:"authorizer_ip_address"`
}

func (r ProofOfAuthorizationRequestSubmissionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProofOfAuthorizationRequestSubmissionListParams struct {
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
	// ID of the proof of authorization request.
	ProofOfAuthorizationRequestID param.Field[string] `query:"proof_of_authorization_request_id"`
}

// URLQuery serializes [ProofOfAuthorizationRequestSubmissionListParams]'s query
// parameters as `url.Values`.
func (r ProofOfAuthorizationRequestSubmissionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
