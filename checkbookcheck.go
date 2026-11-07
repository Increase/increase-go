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

// CheckbookCheckService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckbookCheckService] method instead.
type CheckbookCheckService struct {
	Options []option.RequestOption
}

// NewCheckbookCheckService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckbookCheckService(opts ...option.RequestOption) (r *CheckbookCheckService) {
	r = &CheckbookCheckService{}
	r.Options = opts
	return
}

// Retrieve a Checkbook Check
func (r *CheckbookCheckService) Get(ctx context.Context, checkbookCheckID string, opts ...option.RequestOption) (res *CheckbookCheck, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkbookCheckID == "" {
		err = errors.New("missing required checkbook_check_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("checkbook_checks/%s", checkbookCheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Checkbook Checks
func (r *CheckbookCheckService) List(ctx context.Context, query CheckbookCheckListParams, opts ...option.RequestOption) (res *pagination.Page[CheckbookCheck], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "checkbook_checks"
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

// List Checkbook Checks
func (r *CheckbookCheckService) ListAutoPaging(ctx context.Context, query CheckbookCheckListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CheckbookCheck] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Stop a Checkbook Check. A stopped check cannot be deposited.
func (r *CheckbookCheckService) Stop(ctx context.Context, checkbookCheckID string, opts ...option.RequestOption) (res *CheckbookCheck, err error) {
	opts = slices.Concat(r.Options, opts)
	if checkbookCheckID == "" {
		err = errors.New("missing required checkbook_check_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("checkbook_checks/%s/stop", checkbookCheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A Checkbook Check is a single check within a Checkbook.
type CheckbookCheck struct {
	// The Checkbook Check's identifier.
	ID string `json:"id" api:"required"`
	// The check number printed on the check.
	CheckNumber int64 `json:"check_number" api:"required"`
	// The identifier of the Checkbook the check belongs to.
	CheckbookID string `json:"checkbook_id" api:"required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Checkbook Check was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The status of the Checkbook Check.
	Status CheckbookCheckStatus `json:"status" api:"required"`
	// A constant representing the object's type. For this resource it will always be
	// `checkbook_check`.
	Type CheckbookCheckType `json:"type" api:"required"`
	JSON checkbookCheckJSON `json:"-"`
}

// checkbookCheckJSON contains the JSON metadata for the struct [CheckbookCheck]
type checkbookCheckJSON struct {
	ID          apijson.Field
	CheckNumber apijson.Field
	CheckbookID apijson.Field
	CreatedAt   apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckbookCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkbookCheckJSON) RawJSON() string {
	return r.raw
}

// The status of the Checkbook Check.
type CheckbookCheckStatus string

const (
	CheckbookCheckStatusPendingDeposit    CheckbookCheckStatus = "pending_deposit"
	CheckbookCheckStatusDeposited         CheckbookCheckStatus = "deposited"
	CheckbookCheckStatusStopped           CheckbookCheckStatus = "stopped"
	CheckbookCheckStatusRequiresAttention CheckbookCheckStatus = "requires_attention"
)

func (r CheckbookCheckStatus) IsKnown() bool {
	switch r {
	case CheckbookCheckStatusPendingDeposit, CheckbookCheckStatusDeposited, CheckbookCheckStatusStopped, CheckbookCheckStatusRequiresAttention:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `checkbook_check`.
type CheckbookCheckType string

const (
	CheckbookCheckTypeCheckbookCheck CheckbookCheckType = "checkbook_check"
)

func (r CheckbookCheckType) IsKnown() bool {
	switch r {
	case CheckbookCheckTypeCheckbookCheck:
		return true
	}
	return false
}

type CheckbookCheckListParams struct {
	// Filter Checkbook Checks to those drawn on the specified Account Number.
	AccountNumberID param.Field[string] `query:"account_number_id"`
	// Filter Checkbook Checks to the one with the specified check number.
	CheckNumber param.Field[int64] `query:"check_number"`
	// Filter Checkbook Checks to those belonging to the specified Checkbook.
	CheckbookID param.Field[string] `query:"checkbook_id"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	//
	// Defaults to `100`.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CheckbookCheckListParams]'s query parameters as
// `url.Values`.
func (r CheckbookCheckListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
