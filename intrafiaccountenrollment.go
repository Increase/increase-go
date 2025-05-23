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

// IntrafiAccountEnrollmentService contains methods and other services that help
// with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntrafiAccountEnrollmentService] method instead.
type IntrafiAccountEnrollmentService struct {
	Options []option.RequestOption
}

// NewIntrafiAccountEnrollmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIntrafiAccountEnrollmentService(opts ...option.RequestOption) (r *IntrafiAccountEnrollmentService) {
	r = &IntrafiAccountEnrollmentService{}
	r.Options = opts
	return
}

// Enroll an account in the IntraFi deposit sweep network
func (r *IntrafiAccountEnrollmentService) New(ctx context.Context, body IntrafiAccountEnrollmentNewParams, opts ...option.RequestOption) (res *IntrafiAccountEnrollment, err error) {
	opts = append(r.Options[:], opts...)
	path := "intrafi_account_enrollments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an IntraFi Account Enrollment
func (r *IntrafiAccountEnrollmentService) Get(ctx context.Context, intrafiAccountEnrollmentID string, opts ...option.RequestOption) (res *IntrafiAccountEnrollment, err error) {
	opts = append(r.Options[:], opts...)
	if intrafiAccountEnrollmentID == "" {
		err = errors.New("missing required intrafi_account_enrollment_id parameter")
		return
	}
	path := fmt.Sprintf("intrafi_account_enrollments/%s", intrafiAccountEnrollmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List IntraFi Account Enrollments
func (r *IntrafiAccountEnrollmentService) List(ctx context.Context, query IntrafiAccountEnrollmentListParams, opts ...option.RequestOption) (res *pagination.Page[IntrafiAccountEnrollment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "intrafi_account_enrollments"
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

// List IntraFi Account Enrollments
func (r *IntrafiAccountEnrollmentService) ListAutoPaging(ctx context.Context, query IntrafiAccountEnrollmentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[IntrafiAccountEnrollment] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Unenroll an account from IntraFi
func (r *IntrafiAccountEnrollmentService) Unenroll(ctx context.Context, intrafiAccountEnrollmentID string, opts ...option.RequestOption) (res *IntrafiAccountEnrollment, err error) {
	opts = append(r.Options[:], opts...)
	if intrafiAccountEnrollmentID == "" {
		err = errors.New("missing required intrafi_account_enrollment_id parameter")
		return
	}
	path := fmt.Sprintf("intrafi_account_enrollments/%s/unenroll", intrafiAccountEnrollmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// IntraFi is a
// [network of financial institutions](https://www.intrafi.com/network-banks) that
// allows Increase users to sweep funds to multiple banks. This enables accounts to
// become eligible for additional Federal Deposit Insurance Corporation (FDIC)
// insurance. An IntraFi Account Enrollment object represents the status of an
// account in the network. Sweeping an account to IntraFi doesn't affect funds
// availability.
type IntrafiAccountEnrollment struct {
	// The identifier of this enrollment at IntraFi.
	ID string `json:"id,required"`
	// The identifier of the Increase Account being swept into the network.
	AccountID string `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the enrollment was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The identifier of the account in IntraFi's system. This identifier will be
	// printed on any IntraFi statements or documents.
	IntrafiID string `json:"intrafi_id,required"`
	// The status of the account in the network. An account takes about one business
	// day to go from `pending_enrolling` to `enrolled`.
	Status IntrafiAccountEnrollmentStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `intrafi_account_enrollment`.
	Type IntrafiAccountEnrollmentType `json:"type,required"`
	JSON intrafiAccountEnrollmentJSON `json:"-"`
}

// intrafiAccountEnrollmentJSON contains the JSON metadata for the struct
// [IntrafiAccountEnrollment]
type intrafiAccountEnrollmentJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	CreatedAt      apijson.Field
	IdempotencyKey apijson.Field
	IntrafiID      apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IntrafiAccountEnrollment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intrafiAccountEnrollmentJSON) RawJSON() string {
	return r.raw
}

// The status of the account in the network. An account takes about one business
// day to go from `pending_enrolling` to `enrolled`.
type IntrafiAccountEnrollmentStatus string

const (
	IntrafiAccountEnrollmentStatusPendingEnrolling   IntrafiAccountEnrollmentStatus = "pending_enrolling"
	IntrafiAccountEnrollmentStatusEnrolled           IntrafiAccountEnrollmentStatus = "enrolled"
	IntrafiAccountEnrollmentStatusPendingUnenrolling IntrafiAccountEnrollmentStatus = "pending_unenrolling"
	IntrafiAccountEnrollmentStatusUnenrolled         IntrafiAccountEnrollmentStatus = "unenrolled"
	IntrafiAccountEnrollmentStatusRequiresAttention  IntrafiAccountEnrollmentStatus = "requires_attention"
)

func (r IntrafiAccountEnrollmentStatus) IsKnown() bool {
	switch r {
	case IntrafiAccountEnrollmentStatusPendingEnrolling, IntrafiAccountEnrollmentStatusEnrolled, IntrafiAccountEnrollmentStatusPendingUnenrolling, IntrafiAccountEnrollmentStatusUnenrolled, IntrafiAccountEnrollmentStatusRequiresAttention:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `intrafi_account_enrollment`.
type IntrafiAccountEnrollmentType string

const (
	IntrafiAccountEnrollmentTypeIntrafiAccountEnrollment IntrafiAccountEnrollmentType = "intrafi_account_enrollment"
)

func (r IntrafiAccountEnrollmentType) IsKnown() bool {
	switch r {
	case IntrafiAccountEnrollmentTypeIntrafiAccountEnrollment:
		return true
	}
	return false
}

type IntrafiAccountEnrollmentNewParams struct {
	// The identifier for the account to be added to IntraFi.
	AccountID param.Field[string] `json:"account_id,required"`
	// The contact email for the account owner, to be shared with IntraFi.
	EmailAddress param.Field[string] `json:"email_address,required"`
}

func (r IntrafiAccountEnrollmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntrafiAccountEnrollmentListParams struct {
	// Filter IntraFi Account Enrollments to the one belonging to an account.
	AccountID param.Field[string] `query:"account_id"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                                    `query:"limit"`
	Status param.Field[IntrafiAccountEnrollmentListParamsStatus] `query:"status"`
}

// URLQuery serializes [IntrafiAccountEnrollmentListParams]'s query parameters as
// `url.Values`.
func (r IntrafiAccountEnrollmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IntrafiAccountEnrollmentListParamsStatus struct {
	// Filter IntraFi Account Enrollments for those with the specified status or
	// statuses. For GET requests, this should be encoded as a comma-delimited string,
	// such as `?in=one,two,three`.
	In param.Field[[]IntrafiAccountEnrollmentListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [IntrafiAccountEnrollmentListParamsStatus]'s query
// parameters as `url.Values`.
func (r IntrafiAccountEnrollmentListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IntrafiAccountEnrollmentListParamsStatusIn string

const (
	IntrafiAccountEnrollmentListParamsStatusInPendingEnrolling   IntrafiAccountEnrollmentListParamsStatusIn = "pending_enrolling"
	IntrafiAccountEnrollmentListParamsStatusInEnrolled           IntrafiAccountEnrollmentListParamsStatusIn = "enrolled"
	IntrafiAccountEnrollmentListParamsStatusInPendingUnenrolling IntrafiAccountEnrollmentListParamsStatusIn = "pending_unenrolling"
	IntrafiAccountEnrollmentListParamsStatusInUnenrolled         IntrafiAccountEnrollmentListParamsStatusIn = "unenrolled"
	IntrafiAccountEnrollmentListParamsStatusInRequiresAttention  IntrafiAccountEnrollmentListParamsStatusIn = "requires_attention"
)

func (r IntrafiAccountEnrollmentListParamsStatusIn) IsKnown() bool {
	switch r {
	case IntrafiAccountEnrollmentListParamsStatusInPendingEnrolling, IntrafiAccountEnrollmentListParamsStatusInEnrolled, IntrafiAccountEnrollmentListParamsStatusInPendingUnenrolling, IntrafiAccountEnrollmentListParamsStatusInUnenrolled, IntrafiAccountEnrollmentListParamsStatusInRequiresAttention:
		return true
	}
	return false
}
