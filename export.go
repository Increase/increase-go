package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type ExportService struct {
	Options []option.RequestOption
}

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
	JSON ExportJSON
}

type ExportJSON struct {
	ID              apijson.Metadata
	CreatedAt       apijson.Metadata
	Category        apijson.Metadata
	Status          apijson.Metadata
	FileID          apijson.Metadata
	FileDownloadURL apijson.Metadata
	Type            apijson.Metadata
	raw             string
	Extras          map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Export using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
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
	Category field.Field[ExportNewParamsCategory] `json:"category,required"`
	// Options for the created export. Required if `category` is equal to
	// `transaction_csv`.
	TransactionCsv field.Field[ExportNewParamsTransactionCsv] `json:"transaction_csv"`
	// Options for the created export. Required if `category` is equal to
	// `balance_csv`.
	BalanceCsv field.Field[ExportNewParamsBalanceCsv] `json:"balance_csv"`
}

// MarshalJSON serializes ExportNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r ExportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExportNewParamsCategory string

const (
	ExportNewParamsCategoryTransactionCsv ExportNewParamsCategory = "transaction_csv"
	ExportNewParamsCategoryBalanceCsv     ExportNewParamsCategory = "balance_csv"
)

type ExportNewParamsTransactionCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID field.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt field.Field[ExportNewParamsTransactionCsvCreatedAt] `json:"created_at"`
}

type ExportNewParamsTransactionCsvCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

type ExportNewParamsBalanceCsv struct {
	// Filter exported Transactions to the specified Account.
	AccountID field.Field[string] `json:"account_id"`
	// Filter results by time range on the `created_at` attribute.
	CreatedAt field.Field[ExportNewParamsBalanceCsvCreatedAt] `json:"created_at"`
}

type ExportNewParamsBalanceCsvCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `json:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `json:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `json:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `json:"on_or_before" format:"date-time"`
}

type ExportListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes ExportListParams into a url.Values of the query parameters
// associated with this value
func (r ExportListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ExportListResponse struct {
	// The contents of the list.
	Data []Export `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       ExportListResponseJSON
}

type ExportListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExportListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
