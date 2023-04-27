package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

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
