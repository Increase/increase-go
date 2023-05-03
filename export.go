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

// ExportService contains methods and other services that help with interacting
// with the increase API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewExportService] method instead.
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
	opts = append(r.Options[:], opts...)
	path := "exports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an Export
func (r *ExportService) Get(ctx context.Context, export_id string, opts ...option.RequestOption) (res *Export, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("exports/%s", export_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Exports
func (r *ExportService) List(ctx context.Context, query ExportListParams, opts ...option.RequestOption) (res *shared.Page[Export], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *ExportService) ListAutoPaging(ctx context.Context, query ExportListParams, opts ...option.RequestOption) *shared.PageAutoPager[Export] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Exports are batch summaries of your Increase data. You can make them from the
// API or dashboard. Since they can take a while, they are generated
// asynchronously. We send a webhook when they are ready. For more information,
// please read our
// [Exports documentation](https://increase.com/documentation/exports).
type Export struct {
	// The Export identifier.
	ID string `json:"id,required"`
	// The time the Export was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The category of the Export. We may add additional possible values for this enum
	// over time; your application should be able to handle that gracefully.
	Category ExportCategory `json:"category,required"`
	// The status of the Export.
	Status ExportStatus `json:"status,required"`
	// The File containing the contents of the Export. This will be present when the
	// Export's status transitions to `complete`.
	FileID string `json:"file_id,required,nullable"`
	// A URL at which the Export's file can be downloaded. This will be present when
	// the Export's status transitions to `complete`.
	FileDownloadURL string `json:"file_download_url,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `export`.
	Type ExportType `json:"type,required"`
	JSON exportJSON
}

// exportJSON contains the JSON metadata for the struct [Export]
type exportJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Category        apijson.Field
	Status          apijson.Field
	FileID          apijson.Field
	FileDownloadURL apijson.Field
	Type            apijson.Field
	raw             string
	Extras          map[string]apijson.Field
}

func (r *Export) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExportCategory string

const (
	ExportCategoryTransactionCsv ExportCategory = "transaction_csv"
	ExportCategoryBalanceCsv     ExportCategory = "balance_csv"
)

type ExportStatus string

const (
	ExportStatusPending  ExportStatus = "pending"
	ExportStatusComplete ExportStatus = "complete"
)

type ExportType string

const (
	ExportTypeExport ExportType = "export"
)

type ExportNewParams struct {
	// The type of Export to create.
	Category param.Field[ExportNewParamsCategory] `json:"category,required"`
	// Options for the created export. Required if `category` is equal to
	// `transaction_csv`.
	TransactionCsv param.Field[ExportNewParamsTransactionCsv] `json:"transaction_csv"`
	// Options for the created export. Required if `category` is equal to
	// `balance_csv`.
	BalanceCsv param.Field[ExportNewParamsBalanceCsv] `json:"balance_csv"`
}

func (r ExportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExportNewParamsCategory string

const (
	ExportNewParamsCategoryTransactionCsv ExportNewParamsCategory = "transaction_csv"
	ExportNewParamsCategoryBalanceCsv     ExportNewParamsCategory = "balance_csv"
)

// Options for the created export. Required if `category` is equal to
// `transaction_csv`.
type ExportNewParamsTransactionCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID param.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsTransactionCsvCreatedAt] `json:"created_at"`
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

// Options for the created export. Required if `category` is equal to
// `balance_csv`.
type ExportNewParamsBalanceCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID param.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt param.Field[ExportNewParamsBalanceCsvCreatedAt] `json:"created_at"`
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

type ExportListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ExportListParams]'s query parameters as `url.Values`.
func (r ExportListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// A list of Export objects
type ExportListResponse struct {
	// The contents of the list.
	Data []Export `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       exportListResponseJSON
}

// exportListResponseJSON contains the JSON metadata for the struct
// [ExportListResponse]
type exportListResponseJSON struct {
	Data       apijson.Field
	NextCursor apijson.Field
	raw        string
	Extras     map[string]apijson.Field
}

func (r *ExportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
