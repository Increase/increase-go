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
)

// ExportService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExportService] method instead.
type ExportService struct {
	Options []option.RequestOption
}

// NewExportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExportService(opts ...option.RequestOption) (r *ExportService) {
	r = &ExportService{}
	r.Options = opts
	return
}

// Create an Export
func (r *ExportService) New(ctx context.Context, body ExportNewParams, opts ...option.RequestOption) (res *Export, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "exports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Export
func (r *ExportService) Get(ctx context.Context, exportID string, opts ...option.RequestOption) (res *Export, err error) {
	opts = slices.Concat(r.Options, opts)
	if exportID == "" {
		err = errors.New("missing required export_id parameter")
		return
	}
	path := fmt.Sprintf("exports/%s", exportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Exports
func (r *ExportService) List(ctx context.Context, query ExportListParams, opts ...option.RequestOption) (res *ExportListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "exports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Exports are batch summaries of your Increase data. You can make them from the
// API or dashboard. Since they can take a while, they are generated
// asynchronously. We send a webhook when they are ready. For more information,
// please read our
// [Exports documentation](https://increase.com/documentation/exports).
type Export struct {
	// The Export identifier.
	ID string `json:"id,required"`
	// The category of the Export. We may add additional possible values for this enum
	// over time; your application should be able to handle that gracefully.
	Category ExportCategory `json:"category,required"`
	// The time the Export was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A URL at which the Export's file can be downloaded. This will be present when
	// the Export's status transitions to `complete`.
	FileDownloadURL string `json:"file_download_url,required,nullable"`
	// The File containing the contents of the Export. This will be present when the
	// Export's status transitions to `complete`.
	FileID string `json:"file_id,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The status of the Export.
	Status ExportStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `export`.
	Type ExportType `json:"type,required"`
	JSON exportJSON `json:"-"`
}

// exportJSON contains the JSON metadata for the struct [Export]
type exportJSON struct {
	ID              apijson.Field
	Category        apijson.Field
	CreatedAt       apijson.Field
	FileDownloadURL apijson.Field
	FileID          apijson.Field
	IdempotencyKey  apijson.Field
	Status          apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Export) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportJSON) RawJSON() string {
	return r.raw
}

// The category of the Export. We may add additional possible values for this enum
// over time; your application should be able to handle that gracefully.
type ExportCategory string

const (
	ExportCategoryAccountStatementOfx          ExportCategory = "account_statement_ofx"
	ExportCategoryAccountStatementBai2         ExportCategory = "account_statement_bai2"
	ExportCategoryTransactionCsv               ExportCategory = "transaction_csv"
	ExportCategoryBalanceCsv                   ExportCategory = "balance_csv"
	ExportCategoryBookkeepingAccountBalanceCsv ExportCategory = "bookkeeping_account_balance_csv"
	ExportCategoryEntityCsv                    ExportCategory = "entity_csv"
	ExportCategoryVendorCsv                    ExportCategory = "vendor_csv"
	ExportCategoryDashboardTableCsv            ExportCategory = "dashboard_table_csv"
)

func (r ExportCategory) IsKnown() bool {
	switch r {
	case ExportCategoryAccountStatementOfx, ExportCategoryAccountStatementBai2, ExportCategoryTransactionCsv, ExportCategoryBalanceCsv, ExportCategoryBookkeepingAccountBalanceCsv, ExportCategoryEntityCsv, ExportCategoryVendorCsv, ExportCategoryDashboardTableCsv:
		return true
	}
	return false
}

// The status of the Export.
type ExportStatus string

const (
	ExportStatusPending  ExportStatus = "pending"
	ExportStatusComplete ExportStatus = "complete"
	ExportStatusFailed   ExportStatus = "failed"
)

func (r ExportStatus) IsKnown() bool {
	switch r {
	case ExportStatusPending, ExportStatusComplete, ExportStatusFailed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `export`.
type ExportType string

const (
	ExportTypeExport ExportType = "export"
)

func (r ExportType) IsKnown() bool {
	switch r {
	case ExportTypeExport:
		return true
	}
	return false
}

// A list of Export objects.
type ExportListResponse struct {
	// The contents of the list.
	Data []Export `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor  string                 `json:"next_cursor,required,nullable"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        exportListResponseJSON `json:"-"`
}

// exportListResponseJSON contains the JSON metadata for the struct
// [ExportListResponse]
type exportListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportListResponseJSON) RawJSON() string {
	return r.raw
}

type ExportNewParams struct {
	// The type of Export to create.
	Category param.Field[ExportNewParamsCategory] `json:"category,required"`
	// Options for the created export. Required if `category` is equal to
	// `account_statement_bai2`.
	AccountStatementBai2 param.Field[ExportNewParamsAccountStatementBai2] `json:"account_statement_bai2"`
	// Options for the created export. Required if `category` is equal to
	// `account_statement_ofx`.
	AccountStatementOfx param.Field[ExportNewParamsAccountStatementOfx] `json:"account_statement_ofx"`
	// Options for the created export. Required if `category` is equal to
	// `balance_csv`.
	BalanceCsv param.Field[ExportNewParamsBalanceCsv] `json:"balance_csv"`
	// Options for the created export. Required if `category` is equal to
	// `bookkeeping_account_balance_csv`.
	BookkeepingAccountBalanceCsv param.Field[ExportNewParamsBookkeepingAccountBalanceCsv] `json:"bookkeeping_account_balance_csv"`
	// Options for the created export. Required if `category` is equal to `entity_csv`.
	EntityCsv param.Field[ExportNewParamsEntityCsv] `json:"entity_csv"`
	// Options for the created export. Required if `category` is equal to
	// `transaction_csv`.
	TransactionCsv param.Field[ExportNewParamsTransactionCsv] `json:"transaction_csv"`
	// Options for the created export. Required if `category` is equal to `vendor_csv`.
	VendorCsv param.Field[ExportNewParamsVendorCsv] `json:"vendor_csv"`
}

func (r ExportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of Export to create.
type ExportNewParamsCategory string

const (
	ExportNewParamsCategoryAccountStatementOfx          ExportNewParamsCategory = "account_statement_ofx"
	ExportNewParamsCategoryAccountStatementBai2         ExportNewParamsCategory = "account_statement_bai2"
	ExportNewParamsCategoryTransactionCsv               ExportNewParamsCategory = "transaction_csv"
	ExportNewParamsCategoryBalanceCsv                   ExportNewParamsCategory = "balance_csv"
	ExportNewParamsCategoryBookkeepingAccountBalanceCsv ExportNewParamsCategory = "bookkeeping_account_balance_csv"
	ExportNewParamsCategoryEntityCsv                    ExportNewParamsCategory = "entity_csv"
	ExportNewParamsCategoryVendorCsv                    ExportNewParamsCategory = "vendor_csv"
)

func (r ExportNewParamsCategory) IsKnown() bool {
	switch r {
	case ExportNewParamsCategoryAccountStatementOfx, ExportNewParamsCategoryAccountStatementBai2, ExportNewParamsCategoryTransactionCsv, ExportNewParamsCategoryBalanceCsv, ExportNewParamsCategoryBookkeepingAccountBalanceCsv, ExportNewParamsCategoryEntityCsv, ExportNewParamsCategoryVendorCsv:
		return true
	}
	return false
}

// Options for the created export. Required if `category` is equal to
// `account_statement_bai2`.
type ExportNewParamsAccountStatementBai2 struct {
	// The Account to create a BAI2 report for. If not provided, all open accounts will
	// be included.
	AccountID param.Field[string] `json:"account_id"`
	// The date to create a BAI2 report for. If not provided, the current date will be
	// used. The timezone is UTC. If the current date is used, the report will include
	// intraday balances, otherwise it will include end-of-day balances for the
	// provided date.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// The Program to create a BAI2 report for. If not provided, all open accounts will
	// be included.
	ProgramID param.Field[string] `json:"program_id"`
}

func (r ExportNewParamsAccountStatementBai2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to
// `account_statement_ofx`.
type ExportNewParamsAccountStatementOfx struct {
	// The Account to create a statement for.
	AccountID param.Field[string] `json:"account_id,required"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsAccountStatementOfxCreatedAt] `json:"created_at"`
}

func (r ExportNewParamsAccountStatementOfx) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter results by time range on the `created_at` attribute.
type ExportNewParamsAccountStatementOfxCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

func (r ExportNewParamsAccountStatementOfxCreatedAt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to
// `balance_csv`.
type ExportNewParamsBalanceCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID param.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsBalanceCsvCreatedAt] `json:"created_at"`
	// Filter exported Transactions to the specified Program.
	ProgramID param.Field[string] `json:"program_id"`
}

func (r ExportNewParamsBalanceCsv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter results by time range on the `created_at` attribute.
type ExportNewParamsBalanceCsvCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

func (r ExportNewParamsBalanceCsvCreatedAt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to
// `bookkeeping_account_balance_csv`.
type ExportNewParamsBookkeepingAccountBalanceCsv struct {
	// Filter exported Transactions to the specified Bookkeeping Account.
	BookkeepingAccountID param.Field[string] `json:"bookkeeping_account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsBookkeepingAccountBalanceCsvCreatedAt] `json:"created_at"`
}

func (r ExportNewParamsBookkeepingAccountBalanceCsv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter results by time range on the `created_at` attribute.
type ExportNewParamsBookkeepingAccountBalanceCsvCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

func (r ExportNewParamsBookkeepingAccountBalanceCsvCreatedAt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to `entity_csv`.
type ExportNewParamsEntityCsv struct {
	// Entity statuses to filter by.
	Status param.Field[ExportNewParamsEntityCsvStatus] `json:"status"`
}

func (r ExportNewParamsEntityCsv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Entity statuses to filter by.
type ExportNewParamsEntityCsvStatus struct {
	// Entity statuses to filter by. For GET requests, this should be encoded as a
	// comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]ExportNewParamsEntityCsvStatusIn] `json:"in,required"`
}

func (r ExportNewParamsEntityCsvStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExportNewParamsEntityCsvStatusIn string

const (
	ExportNewParamsEntityCsvStatusInActive   ExportNewParamsEntityCsvStatusIn = "active"
	ExportNewParamsEntityCsvStatusInArchived ExportNewParamsEntityCsvStatusIn = "archived"
	ExportNewParamsEntityCsvStatusInDisabled ExportNewParamsEntityCsvStatusIn = "disabled"
)

func (r ExportNewParamsEntityCsvStatusIn) IsKnown() bool {
	switch r {
	case ExportNewParamsEntityCsvStatusInActive, ExportNewParamsEntityCsvStatusInArchived, ExportNewParamsEntityCsvStatusInDisabled:
		return true
	}
	return false
}

// Options for the created export. Required if `category` is equal to
// `transaction_csv`.
type ExportNewParamsTransactionCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID param.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsTransactionCsvCreatedAt] `json:"created_at"`
	// Filter exported Transactions to the specified Program.
	ProgramID param.Field[string] `json:"program_id"`
}

func (r ExportNewParamsTransactionCsv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter results by time range on the `created_at` attribute.
type ExportNewParamsTransactionCsvCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

func (r ExportNewParamsTransactionCsvCreatedAt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to `vendor_csv`.
type ExportNewParamsVendorCsv struct {
}

func (r ExportNewParamsVendorCsv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExportListParams struct {
	Category  param.Field[ExportListParamsCategory]  `query:"category"`
	CreatedAt param.Field[ExportListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                  `query:"limit"`
	Status param.Field[ExportListParamsStatus] `query:"status"`
}

// URLQuery serializes [ExportListParams]'s query parameters as `url.Values`.
func (r ExportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ExportListParamsCategory struct {
	// Filter Exports for those with the specified category or categories. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]ExportListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes [ExportListParamsCategory]'s query parameters as
// `url.Values`.
func (r ExportListParamsCategory) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ExportListParamsCategoryIn string

const (
	ExportListParamsCategoryInAccountStatementOfx          ExportListParamsCategoryIn = "account_statement_ofx"
	ExportListParamsCategoryInAccountStatementBai2         ExportListParamsCategoryIn = "account_statement_bai2"
	ExportListParamsCategoryInTransactionCsv               ExportListParamsCategoryIn = "transaction_csv"
	ExportListParamsCategoryInBalanceCsv                   ExportListParamsCategoryIn = "balance_csv"
	ExportListParamsCategoryInBookkeepingAccountBalanceCsv ExportListParamsCategoryIn = "bookkeeping_account_balance_csv"
	ExportListParamsCategoryInEntityCsv                    ExportListParamsCategoryIn = "entity_csv"
	ExportListParamsCategoryInVendorCsv                    ExportListParamsCategoryIn = "vendor_csv"
	ExportListParamsCategoryInDashboardTableCsv            ExportListParamsCategoryIn = "dashboard_table_csv"
)

func (r ExportListParamsCategoryIn) IsKnown() bool {
	switch r {
	case ExportListParamsCategoryInAccountStatementOfx, ExportListParamsCategoryInAccountStatementBai2, ExportListParamsCategoryInTransactionCsv, ExportListParamsCategoryInBalanceCsv, ExportListParamsCategoryInBookkeepingAccountBalanceCsv, ExportListParamsCategoryInEntityCsv, ExportListParamsCategoryInVendorCsv, ExportListParamsCategoryInDashboardTableCsv:
		return true
	}
	return false
}

type ExportListParamsCreatedAt struct {
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

// URLQuery serializes [ExportListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r ExportListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ExportListParamsStatus struct {
	// Filter Exports for those with the specified status or statuses. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]ExportListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [ExportListParamsStatus]'s query parameters as `url.Values`.
func (r ExportListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ExportListParamsStatusIn string

const (
	ExportListParamsStatusInPending  ExportListParamsStatusIn = "pending"
	ExportListParamsStatusInComplete ExportListParamsStatusIn = "complete"
	ExportListParamsStatusInFailed   ExportListParamsStatusIn = "failed"
)

func (r ExportListParamsStatusIn) IsKnown() bool {
	switch r {
	case ExportListParamsStatusInPending, ExportListParamsStatusInComplete, ExportListParamsStatusInFailed:
		return true
	}
	return false
}
