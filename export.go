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
func (r *ExportService) List(ctx context.Context, query ExportListParams, opts ...option.RequestOption) (res *pagination.Page[Export], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "exports"
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

// List Exports
func (r *ExportService) ListAutoPaging(ctx context.Context, query ExportListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Export] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Exports are generated files. Some exports can contain a lot of data, like a CSV
// of your transactions. Others can be a single document, like a tax form. Since
// they can take a while, they are generated asynchronously. We send a webhook when
// they are ready. For more information, please read our
// [Exports documentation](https://increase.com/documentation/exports).
type Export struct {
	// The Export identifier.
	ID string `json:"id,required"`
	// Details of the account statement BAI2 export. This field will be present when
	// the `category` is equal to `account_statement_bai2`.
	AccountStatementBai2 ExportAccountStatementBai2 `json:"account_statement_bai2,required,nullable"`
	// Details of the account statement OFX export. This field will be present when the
	// `category` is equal to `account_statement_ofx`.
	AccountStatementOfx ExportAccountStatementOfx `json:"account_statement_ofx,required,nullable"`
	// Details of the account verification letter export. This field will be present
	// when the `category` is equal to `account_verification_letter`.
	AccountVerificationLetter ExportAccountVerificationLetter `json:"account_verification_letter,required,nullable"`
	// Details of the balance CSV export. This field will be present when the
	// `category` is equal to `balance_csv`.
	BalanceCsv ExportBalanceCsv `json:"balance_csv,required,nullable"`
	// Details of the bookkeeping account balance CSV export. This field will be
	// present when the `category` is equal to `bookkeeping_account_balance_csv`.
	BookkeepingAccountBalanceCsv ExportBookkeepingAccountBalanceCsv `json:"bookkeeping_account_balance_csv,required,nullable"`
	// The category of the Export. We may add additional possible values for this enum
	// over time; your application should be able to handle that gracefully.
	Category ExportCategory `json:"category,required"`
	// The time the Export was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Details of the dashboard table CSV export. This field will be present when the
	// `category` is equal to `dashboard_table_csv`.
	DashboardTableCsv ExportDashboardTableCsv `json:"dashboard_table_csv,required,nullable"`
	// Details of the entity CSV export. This field will be present when the `category`
	// is equal to `entity_csv`.
	EntityCsv ExportEntityCsv `json:"entity_csv,required,nullable"`
	// Details of the Form 1099-INT export. This field will be present when the
	// `category` is equal to `form_1099_int`.
	Form1099Int ExportForm1099Int `json:"form_1099_int,required,nullable"`
	// Details of the Form 1099-MISC export. This field will be present when the
	// `category` is equal to `form_1099_misc`.
	Form1099Misc ExportForm1099Misc `json:"form_1099_misc,required,nullable"`
	// Details of the funding instructions export. This field will be present when the
	// `category` is equal to `funding_instructions`.
	FundingInstructions ExportFundingInstructions `json:"funding_instructions,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The result of the Export. This will be present when the Export's status
	// transitions to `complete`.
	Result ExportResult `json:"result,required,nullable"`
	// The status of the Export.
	Status ExportStatus `json:"status,required"`
	// Details of the transaction CSV export. This field will be present when the
	// `category` is equal to `transaction_csv`.
	TransactionCsv ExportTransactionCsv `json:"transaction_csv,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `export`.
	Type ExportType `json:"type,required"`
	// Details of the vendor CSV export. This field will be present when the `category`
	// is equal to `vendor_csv`.
	VendorCsv ExportVendorCsv `json:"vendor_csv,required,nullable"`
	JSON      exportJSON      `json:"-"`
}

// exportJSON contains the JSON metadata for the struct [Export]
type exportJSON struct {
	ID                           apijson.Field
	AccountStatementBai2         apijson.Field
	AccountStatementOfx          apijson.Field
	AccountVerificationLetter    apijson.Field
	BalanceCsv                   apijson.Field
	BookkeepingAccountBalanceCsv apijson.Field
	Category                     apijson.Field
	CreatedAt                    apijson.Field
	DashboardTableCsv            apijson.Field
	EntityCsv                    apijson.Field
	Form1099Int                  apijson.Field
	Form1099Misc                 apijson.Field
	FundingInstructions          apijson.Field
	IdempotencyKey               apijson.Field
	Result                       apijson.Field
	Status                       apijson.Field
	TransactionCsv               apijson.Field
	Type                         apijson.Field
	VendorCsv                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *Export) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportJSON) RawJSON() string {
	return r.raw
}

// Details of the account statement BAI2 export. This field will be present when
// the `category` is equal to `account_statement_bai2`.
type ExportAccountStatementBai2 struct {
	// Filter results by Account.
	AccountID string `json:"account_id,required,nullable"`
	// The date for which to retrieve the balance.
	EffectiveDate time.Time `json:"effective_date,required,nullable" format:"date"`
	// Filter results by Program.
	ProgramID string                         `json:"program_id,required,nullable"`
	JSON      exportAccountStatementBai2JSON `json:"-"`
}

// exportAccountStatementBai2JSON contains the JSON metadata for the struct
// [ExportAccountStatementBai2]
type exportAccountStatementBai2JSON struct {
	AccountID     apijson.Field
	EffectiveDate apijson.Field
	ProgramID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ExportAccountStatementBai2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportAccountStatementBai2JSON) RawJSON() string {
	return r.raw
}

// Details of the account statement OFX export. This field will be present when the
// `category` is equal to `account_statement_ofx`.
type ExportAccountStatementOfx struct {
	// The Account to create a statement for.
	AccountID string `json:"account_id,required"`
	// Filter transactions by their created date.
	CreatedAt ExportAccountStatementOfxCreatedAt `json:"created_at,required,nullable"`
	JSON      exportAccountStatementOfxJSON      `json:"-"`
}

// exportAccountStatementOfxJSON contains the JSON metadata for the struct
// [ExportAccountStatementOfx]
type exportAccountStatementOfxJSON struct {
	AccountID   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportAccountStatementOfx) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportAccountStatementOfxJSON) RawJSON() string {
	return r.raw
}

// Filter transactions by their created date.
type ExportAccountStatementOfxCreatedAt struct {
	// Filter results to transactions created after this time.
	After time.Time `json:"after,required,nullable" format:"date-time"`
	// Filter results to transactions created before this time.
	Before time.Time                              `json:"before,required,nullable" format:"date-time"`
	JSON   exportAccountStatementOfxCreatedAtJSON `json:"-"`
}

// exportAccountStatementOfxCreatedAtJSON contains the JSON metadata for the struct
// [ExportAccountStatementOfxCreatedAt]
type exportAccountStatementOfxCreatedAtJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportAccountStatementOfxCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportAccountStatementOfxCreatedAtJSON) RawJSON() string {
	return r.raw
}

// Details of the account verification letter export. This field will be present
// when the `category` is equal to `account_verification_letter`.
type ExportAccountVerificationLetter struct {
	// The Account Number to create a letter for.
	AccountNumberID string `json:"account_number_id,required"`
	// The date of the balance to include in the letter.
	BalanceDate time.Time                           `json:"balance_date,required,nullable" format:"date"`
	JSON        exportAccountVerificationLetterJSON `json:"-"`
}

// exportAccountVerificationLetterJSON contains the JSON metadata for the struct
// [ExportAccountVerificationLetter]
type exportAccountVerificationLetterJSON struct {
	AccountNumberID apijson.Field
	BalanceDate     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ExportAccountVerificationLetter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportAccountVerificationLetterJSON) RawJSON() string {
	return r.raw
}

// Details of the balance CSV export. This field will be present when the
// `category` is equal to `balance_csv`.
type ExportBalanceCsv struct {
	// Filter results by Account.
	AccountID string `json:"account_id,required,nullable"`
	// Filter balances by their created date.
	CreatedAt ExportBalanceCsvCreatedAt `json:"created_at,required,nullable"`
	JSON      exportBalanceCsvJSON      `json:"-"`
}

// exportBalanceCsvJSON contains the JSON metadata for the struct
// [ExportBalanceCsv]
type exportBalanceCsvJSON struct {
	AccountID   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportBalanceCsv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportBalanceCsvJSON) RawJSON() string {
	return r.raw
}

// Filter balances by their created date.
type ExportBalanceCsvCreatedAt struct {
	// Filter balances created after this time.
	After time.Time `json:"after,required,nullable" format:"date-time"`
	// Filter balances created before this time.
	Before time.Time                     `json:"before,required,nullable" format:"date-time"`
	JSON   exportBalanceCsvCreatedAtJSON `json:"-"`
}

// exportBalanceCsvCreatedAtJSON contains the JSON metadata for the struct
// [ExportBalanceCsvCreatedAt]
type exportBalanceCsvCreatedAtJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportBalanceCsvCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportBalanceCsvCreatedAtJSON) RawJSON() string {
	return r.raw
}

// Details of the bookkeeping account balance CSV export. This field will be
// present when the `category` is equal to `bookkeeping_account_balance_csv`.
type ExportBookkeepingAccountBalanceCsv struct {
	// Filter results by Bookkeeping Account.
	BookkeepingAccountID string `json:"bookkeeping_account_id,required,nullable"`
	// Filter balances by their created date.
	CreatedAt ExportBookkeepingAccountBalanceCsvCreatedAt `json:"created_at,required,nullable"`
	JSON      exportBookkeepingAccountBalanceCsvJSON      `json:"-"`
}

// exportBookkeepingAccountBalanceCsvJSON contains the JSON metadata for the struct
// [ExportBookkeepingAccountBalanceCsv]
type exportBookkeepingAccountBalanceCsvJSON struct {
	BookkeepingAccountID apijson.Field
	CreatedAt            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ExportBookkeepingAccountBalanceCsv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportBookkeepingAccountBalanceCsvJSON) RawJSON() string {
	return r.raw
}

// Filter balances by their created date.
type ExportBookkeepingAccountBalanceCsvCreatedAt struct {
	// Filter balances created after this time.
	After time.Time `json:"after,required,nullable" format:"date-time"`
	// Filter balances created before this time.
	Before time.Time                                       `json:"before,required,nullable" format:"date-time"`
	JSON   exportBookkeepingAccountBalanceCsvCreatedAtJSON `json:"-"`
}

// exportBookkeepingAccountBalanceCsvCreatedAtJSON contains the JSON metadata for
// the struct [ExportBookkeepingAccountBalanceCsvCreatedAt]
type exportBookkeepingAccountBalanceCsvCreatedAtJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportBookkeepingAccountBalanceCsvCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportBookkeepingAccountBalanceCsvCreatedAtJSON) RawJSON() string {
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
	ExportCategoryAccountVerificationLetter    ExportCategory = "account_verification_letter"
	ExportCategoryFundingInstructions          ExportCategory = "funding_instructions"
	ExportCategoryForm1099Int                  ExportCategory = "form_1099_int"
	ExportCategoryForm1099Misc                 ExportCategory = "form_1099_misc"
)

func (r ExportCategory) IsKnown() bool {
	switch r {
	case ExportCategoryAccountStatementOfx, ExportCategoryAccountStatementBai2, ExportCategoryTransactionCsv, ExportCategoryBalanceCsv, ExportCategoryBookkeepingAccountBalanceCsv, ExportCategoryEntityCsv, ExportCategoryVendorCsv, ExportCategoryDashboardTableCsv, ExportCategoryAccountVerificationLetter, ExportCategoryFundingInstructions, ExportCategoryForm1099Int, ExportCategoryForm1099Misc:
		return true
	}
	return false
}

// Details of the dashboard table CSV export. This field will be present when the
// `category` is equal to `dashboard_table_csv`.
type ExportDashboardTableCsv struct {
	JSON exportDashboardTableCsvJSON `json:"-"`
}

// exportDashboardTableCsvJSON contains the JSON metadata for the struct
// [ExportDashboardTableCsv]
type exportDashboardTableCsvJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportDashboardTableCsv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportDashboardTableCsvJSON) RawJSON() string {
	return r.raw
}

// Details of the entity CSV export. This field will be present when the `category`
// is equal to `entity_csv`.
type ExportEntityCsv struct {
	JSON exportEntityCsvJSON `json:"-"`
}

// exportEntityCsvJSON contains the JSON metadata for the struct [ExportEntityCsv]
type exportEntityCsvJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportEntityCsv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportEntityCsvJSON) RawJSON() string {
	return r.raw
}

// Details of the Form 1099-INT export. This field will be present when the
// `category` is equal to `form_1099_int`.
type ExportForm1099Int struct {
	// The Account the tax form is for.
	AccountID string `json:"account_id,required"`
	// Whether the tax form is a corrected form.
	Corrected bool `json:"corrected,required"`
	// A description of the tax form.
	Description string `json:"description,required"`
	// The tax year for the tax form.
	Year int64                 `json:"year,required"`
	JSON exportForm1099IntJSON `json:"-"`
}

// exportForm1099IntJSON contains the JSON metadata for the struct
// [ExportForm1099Int]
type exportForm1099IntJSON struct {
	AccountID   apijson.Field
	Corrected   apijson.Field
	Description apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportForm1099Int) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportForm1099IntJSON) RawJSON() string {
	return r.raw
}

// Details of the Form 1099-MISC export. This field will be present when the
// `category` is equal to `form_1099_misc`.
type ExportForm1099Misc struct {
	// The Account the tax form is for.
	AccountID string `json:"account_id,required"`
	// Whether the tax form is a corrected form.
	Corrected bool `json:"corrected,required"`
	// The tax year for the tax form.
	Year int64                  `json:"year,required"`
	JSON exportForm1099MiscJSON `json:"-"`
}

// exportForm1099MiscJSON contains the JSON metadata for the struct
// [ExportForm1099Misc]
type exportForm1099MiscJSON struct {
	AccountID   apijson.Field
	Corrected   apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportForm1099Misc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportForm1099MiscJSON) RawJSON() string {
	return r.raw
}

// Details of the funding instructions export. This field will be present when the
// `category` is equal to `funding_instructions`.
type ExportFundingInstructions struct {
	// The Account Number to create funding instructions for.
	AccountNumberID string                        `json:"account_number_id,required"`
	JSON            exportFundingInstructionsJSON `json:"-"`
}

// exportFundingInstructionsJSON contains the JSON metadata for the struct
// [ExportFundingInstructions]
type exportFundingInstructionsJSON struct {
	AccountNumberID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ExportFundingInstructions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportFundingInstructionsJSON) RawJSON() string {
	return r.raw
}

// The result of the Export. This will be present when the Export's status
// transitions to `complete`.
type ExportResult struct {
	// The File containing the contents of the Export.
	FileID string           `json:"file_id,required"`
	JSON   exportResultJSON `json:"-"`
}

// exportResultJSON contains the JSON metadata for the struct [ExportResult]
type exportResultJSON struct {
	FileID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportResultJSON) RawJSON() string {
	return r.raw
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

// Details of the transaction CSV export. This field will be present when the
// `category` is equal to `transaction_csv`.
type ExportTransactionCsv struct {
	// Filter results by Account.
	AccountID string `json:"account_id,required,nullable"`
	// Filter transactions by their created date.
	CreatedAt ExportTransactionCsvCreatedAt `json:"created_at,required,nullable"`
	JSON      exportTransactionCsvJSON      `json:"-"`
}

// exportTransactionCsvJSON contains the JSON metadata for the struct
// [ExportTransactionCsv]
type exportTransactionCsvJSON struct {
	AccountID   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportTransactionCsv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportTransactionCsvJSON) RawJSON() string {
	return r.raw
}

// Filter transactions by their created date.
type ExportTransactionCsvCreatedAt struct {
	// Filter transactions created after this time.
	After time.Time `json:"after,required,nullable" format:"date-time"`
	// Filter transactions created before this time.
	Before time.Time                         `json:"before,required,nullable" format:"date-time"`
	JSON   exportTransactionCsvCreatedAtJSON `json:"-"`
}

// exportTransactionCsvCreatedAtJSON contains the JSON metadata for the struct
// [ExportTransactionCsvCreatedAt]
type exportTransactionCsvCreatedAtJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportTransactionCsvCreatedAt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportTransactionCsvCreatedAtJSON) RawJSON() string {
	return r.raw
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

// Details of the vendor CSV export. This field will be present when the `category`
// is equal to `vendor_csv`.
type ExportVendorCsv struct {
	JSON exportVendorCsvJSON `json:"-"`
}

// exportVendorCsvJSON contains the JSON metadata for the struct [ExportVendorCsv]
type exportVendorCsvJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExportVendorCsv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exportVendorCsvJSON) RawJSON() string {
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
	// `account_verification_letter`.
	AccountVerificationLetter param.Field[ExportNewParamsAccountVerificationLetter] `json:"account_verification_letter"`
	// Options for the created export. Required if `category` is equal to
	// `balance_csv`.
	BalanceCsv param.Field[ExportNewParamsBalanceCsv] `json:"balance_csv"`
	// Options for the created export. Required if `category` is equal to
	// `bookkeeping_account_balance_csv`.
	BookkeepingAccountBalanceCsv param.Field[ExportNewParamsBookkeepingAccountBalanceCsv] `json:"bookkeeping_account_balance_csv"`
	// Options for the created export. Required if `category` is equal to `entity_csv`.
	EntityCsv param.Field[ExportNewParamsEntityCsv] `json:"entity_csv"`
	// Options for the created export. Required if `category` is equal to
	// `funding_instructions`.
	FundingInstructions param.Field[ExportNewParamsFundingInstructions] `json:"funding_instructions"`
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
	ExportNewParamsCategoryAccountVerificationLetter    ExportNewParamsCategory = "account_verification_letter"
	ExportNewParamsCategoryFundingInstructions          ExportNewParamsCategory = "funding_instructions"
)

func (r ExportNewParamsCategory) IsKnown() bool {
	switch r {
	case ExportNewParamsCategoryAccountStatementOfx, ExportNewParamsCategoryAccountStatementBai2, ExportNewParamsCategoryTransactionCsv, ExportNewParamsCategoryBalanceCsv, ExportNewParamsCategoryBookkeepingAccountBalanceCsv, ExportNewParamsCategoryEntityCsv, ExportNewParamsCategoryVendorCsv, ExportNewParamsCategoryAccountVerificationLetter, ExportNewParamsCategoryFundingInstructions:
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
// `account_verification_letter`.
type ExportNewParamsAccountVerificationLetter struct {
	// The Account Number to create a letter for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The date of the balance to include in the letter. Defaults to the current date.
	BalanceDate param.Field[time.Time] `json:"balance_date" format:"date"`
}

func (r ExportNewParamsAccountVerificationLetter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to
// `balance_csv`.
type ExportNewParamsBalanceCsv struct {
	// Filter exported Balances to the specified Account.
	AccountID param.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsBalanceCsvCreatedAt] `json:"created_at"`
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
	// Filter exported Bookkeeping Account Balances to the specified Bookkeeping
	// Account.
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
}

func (r ExportNewParamsEntityCsv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to
// `funding_instructions`.
type ExportNewParamsFundingInstructions struct {
	// The Account Number to create funding instructions for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
}

func (r ExportNewParamsFundingInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for the created export. Required if `category` is equal to
// `transaction_csv`.
type ExportNewParamsTransactionCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID param.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsTransactionCsvCreatedAt] `json:"created_at"`
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
	// Filter Exports for those with the specified category.
	Category  param.Field[ExportListParamsCategory]  `query:"category"`
	CreatedAt param.Field[ExportListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor       param.Field[string]                       `query:"cursor"`
	Form1099Int  param.Field[ExportListParamsForm1099Int]  `query:"form_1099_int"`
	Form1099Misc param.Field[ExportListParamsForm1099Misc] `query:"form_1099_misc"`
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

// Filter Exports for those with the specified category.
type ExportListParamsCategory string

const (
	ExportListParamsCategoryAccountStatementOfx          ExportListParamsCategory = "account_statement_ofx"
	ExportListParamsCategoryAccountStatementBai2         ExportListParamsCategory = "account_statement_bai2"
	ExportListParamsCategoryTransactionCsv               ExportListParamsCategory = "transaction_csv"
	ExportListParamsCategoryBalanceCsv                   ExportListParamsCategory = "balance_csv"
	ExportListParamsCategoryBookkeepingAccountBalanceCsv ExportListParamsCategory = "bookkeeping_account_balance_csv"
	ExportListParamsCategoryEntityCsv                    ExportListParamsCategory = "entity_csv"
	ExportListParamsCategoryVendorCsv                    ExportListParamsCategory = "vendor_csv"
	ExportListParamsCategoryDashboardTableCsv            ExportListParamsCategory = "dashboard_table_csv"
	ExportListParamsCategoryAccountVerificationLetter    ExportListParamsCategory = "account_verification_letter"
	ExportListParamsCategoryFundingInstructions          ExportListParamsCategory = "funding_instructions"
	ExportListParamsCategoryForm1099Int                  ExportListParamsCategory = "form_1099_int"
	ExportListParamsCategoryForm1099Misc                 ExportListParamsCategory = "form_1099_misc"
)

func (r ExportListParamsCategory) IsKnown() bool {
	switch r {
	case ExportListParamsCategoryAccountStatementOfx, ExportListParamsCategoryAccountStatementBai2, ExportListParamsCategoryTransactionCsv, ExportListParamsCategoryBalanceCsv, ExportListParamsCategoryBookkeepingAccountBalanceCsv, ExportListParamsCategoryEntityCsv, ExportListParamsCategoryVendorCsv, ExportListParamsCategoryDashboardTableCsv, ExportListParamsCategoryAccountVerificationLetter, ExportListParamsCategoryFundingInstructions, ExportListParamsCategoryForm1099Int, ExportListParamsCategoryForm1099Misc:
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

type ExportListParamsForm1099Int struct {
	// Filter Form 1099-INT Exports to those for the specified Account.
	AccountID param.Field[string] `query:"account_id"`
}

// URLQuery serializes [ExportListParamsForm1099Int]'s query parameters as
// `url.Values`.
func (r ExportListParamsForm1099Int) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ExportListParamsForm1099Misc struct {
	// Filter Form 1099-MISC Exports to those for the specified Account.
	AccountID param.Field[string] `query:"account_id"`
}

// URLQuery serializes [ExportListParamsForm1099Misc]'s query parameters as
// `url.Values`.
func (r ExportListParamsForm1099Misc) URLQuery() (v url.Values) {
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
