// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
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

// AccountStatementService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStatementService] method instead.
type AccountStatementService struct {
	Options []option.RequestOption
}

// NewAccountStatementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStatementService(opts ...option.RequestOption) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Options = opts
	return
}

// Retrieve an Account Statement
func (r *AccountStatementService) Get(ctx context.Context, accountStatementID string, opts ...option.RequestOption) (res *AccountStatement, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountStatementID == "" {
		err = errors.New("missing required account_statement_id parameter")
		return
	}
	path := fmt.Sprintf("account_statements/%s", accountStatementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Account Statements
func (r *AccountStatementService) List(ctx context.Context, query AccountStatementListParams, opts ...option.RequestOption) (res *pagination.Page[AccountStatement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_statements"
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

// List Account Statements
func (r *AccountStatementService) ListAutoPaging(ctx context.Context, query AccountStatementListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AccountStatement] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Account Statements are generated monthly for every active Account. You can
// access the statement's data via the API or retrieve a PDF with its details via
// its associated File.
type AccountStatement struct {
	// The Account Statement identifier.
	ID string `json:"id,required"`
	// The identifier for the Account this Account Statement belongs to.
	AccountID string `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Statement was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The Account's balance at the start of its statement period.
	EndingBalance int64 `json:"ending_balance,required"`
	// The identifier of the File containing a PDF of the statement.
	FileID string `json:"file_id,required"`
	// The Account's balance at the start of its statement period.
	StartingBalance int64 `json:"starting_balance,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
	// of the period the Account Statement covers.
	StatementPeriodEnd time.Time `json:"statement_period_end,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
	// start of the period the Account Statement covers.
	StatementPeriodStart time.Time `json:"statement_period_start,required" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `account_statement`.
	Type        AccountStatementType   `json:"type,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        accountStatementJSON   `json:"-"`
}

// accountStatementJSON contains the JSON metadata for the struct
// [AccountStatement]
type accountStatementJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	CreatedAt            apijson.Field
	EndingBalance        apijson.Field
	FileID               apijson.Field
	StartingBalance      apijson.Field
	StatementPeriodEnd   apijson.Field
	StatementPeriodStart apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStatementJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `account_statement`.
type AccountStatementType string

const (
	AccountStatementTypeAccountStatement AccountStatementType = "account_statement"
)

func (r AccountStatementType) IsKnown() bool {
	switch r {
	case AccountStatementTypeAccountStatement:
		return true
	}
	return false
}

type AccountStatementListParams struct {
	// Filter Account Statements to those belonging to the specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit                param.Field[int64]                                          `query:"limit"`
	StatementPeriodStart param.Field[AccountStatementListParamsStatementPeriodStart] `query:"statement_period_start"`
}

// URLQuery serializes [AccountStatementListParams]'s query parameters as
// `url.Values`.
func (r AccountStatementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountStatementListParamsStatementPeriodStart struct {
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

// URLQuery serializes [AccountStatementListParamsStatementPeriodStart]'s query
// parameters as `url.Values`.
func (r AccountStatementListParamsStatementPeriodStart) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
