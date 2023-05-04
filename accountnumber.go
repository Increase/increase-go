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

// AccountNumberService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountNumberService] method
// instead.
type AccountNumberService struct {
	Options []option.RequestOption
}

// NewAccountNumberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountNumberService(opts ...option.RequestOption) (r *AccountNumberService) {
	r = &AccountNumberService{}
	r.Options = opts
	return
}

// Create an Account Number
func (r *AccountNumberService) New(ctx context.Context, body AccountNumberNewParams, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := "account_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Account Number
func (r *AccountNumberService) Get(ctx context.Context, account_number_id string, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Account Number
func (r *AccountNumberService) Update(ctx context.Context, account_number_id string, body AccountNumberUpdateParams, opts ...option.RequestOption) (res *AccountNumber, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_numbers/%s", account_number_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Account Numbers
func (r *AccountNumberService) List(ctx context.Context, query AccountNumberListParams, opts ...option.RequestOption) (res *shared.Page[AccountNumber], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_numbers"
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

// List Account Numbers
func (r *AccountNumberService) ListAutoPaging(ctx context.Context, query AccountNumberListParams, opts ...option.RequestOption) *shared.PageAutoPager[AccountNumber] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Each account can have multiple account and routing numbers. We recommend that
// you use a set per vendor. This is similar to how you use different passwords for
// different websites. Account numbers can also be used to seamlessly reconcile
// inbound payments. Generating a unique account number per vendor ensures you
// always know the originator of an incoming payment.
type AccountNumber struct {
	// The identifier for the account this Account Number belongs to.
	AccountID string `json:"account_id,required"`
	// The account number.
	AccountNumber string `json:"account_number,required"`
	// The Account Number identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name you choose for the Account Number.
	Name string `json:"name,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// This indicates if payments can be made to the Account Number.
	Status AccountNumberStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type AccountNumberType `json:"type,required"`
	JSON accountNumberJSON
}

// accountNumberJSON contains the JSON metadata for the struct [AccountNumber]
type accountNumberJSON struct {
	AccountID     apijson.Field
	AccountNumber apijson.Field
	ID            apijson.Field
	CreatedAt     apijson.Field
	Name          apijson.Field
	RoutingNumber apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountNumberStatus string

const (
	AccountNumberStatusActive   AccountNumberStatus = "active"
	AccountNumberStatusDisabled AccountNumberStatus = "disabled"
	AccountNumberStatusCanceled AccountNumberStatus = "canceled"
)

type AccountNumberType string

const (
	AccountNumberTypeAccountNumber AccountNumberType = "account_number"
)

type AccountNumberNewParams struct {
	// The Account the Account Number should belong to.
	AccountID param.Field[string] `json:"account_id,required"`
	// The name you choose for the Account Number.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountNumberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountNumberUpdateParams struct {
	// The name you choose for the Account Number.
	Name param.Field[string] `json:"name"`
	// This indicates if transfers can be made to the Account Number.
	Status param.Field[AccountNumberUpdateParamsStatus] `json:"status"`
}

func (r AccountNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountNumberUpdateParamsStatus string

const (
	AccountNumberUpdateParamsStatusActive   AccountNumberUpdateParamsStatus = "active"
	AccountNumberUpdateParamsStatusDisabled AccountNumberUpdateParamsStatus = "disabled"
	AccountNumberUpdateParamsStatusCanceled AccountNumberUpdateParamsStatus = "canceled"
)

type AccountNumberListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// The status to retrieve Account Numbers for.
	Status param.Field[AccountNumberListParamsStatus] `query:"status"`
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID param.Field[string]                           `query:"account_id"`
	CreatedAt param.Field[AccountNumberListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes [AccountNumberListParams]'s query parameters as
// `url.Values`.
func (r AccountNumberListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type AccountNumberListParamsStatus string

const (
	AccountNumberListParamsStatusActive   AccountNumberListParamsStatus = "active"
	AccountNumberListParamsStatusDisabled AccountNumberListParamsStatus = "disabled"
	AccountNumberListParamsStatusCanceled AccountNumberListParamsStatus = "canceled"
)

type AccountNumberListParamsCreatedAt struct {
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

// URLQuery serializes [AccountNumberListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r AccountNumberListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// A list of Account Number objects
type AccountNumberListResponse struct {
	// The contents of the list.
	Data []AccountNumber `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       accountNumberListResponseJSON
}

// accountNumberListResponseJSON contains the JSON metadata for the struct
// [AccountNumberListResponse]
type accountNumberListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNumberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
