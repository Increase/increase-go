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

// IntrafiExclusionService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntrafiExclusionService] method
// instead.
type IntrafiExclusionService struct {
	Options []option.RequestOption
}

// NewIntrafiExclusionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntrafiExclusionService(opts ...option.RequestOption) (r *IntrafiExclusionService) {
	r = &IntrafiExclusionService{}
	r.Options = opts
	return
}

// Create an IntraFi Exclusion
func (r *IntrafiExclusionService) New(ctx context.Context, body IntrafiExclusionNewParams, opts ...option.RequestOption) (res *IntrafiExclusion, err error) {
	opts = append(r.Options[:], opts...)
	path := "intrafi_exclusions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an IntraFi Exclusion
func (r *IntrafiExclusionService) Get(ctx context.Context, intrafiExclusionID string, opts ...option.RequestOption) (res *IntrafiExclusion, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("intrafi_exclusions/%s", intrafiExclusionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List IntraFi Exclusions.
func (r *IntrafiExclusionService) List(ctx context.Context, query IntrafiExclusionListParams, opts ...option.RequestOption) (res *shared.Page[IntrafiExclusion], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "intrafi_exclusions"
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

// List IntraFi Exclusions.
func (r *IntrafiExclusionService) ListAutoPaging(ctx context.Context, query IntrafiExclusionListParams, opts ...option.RequestOption) *shared.PageAutoPager[IntrafiExclusion] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive an IntraFi Exclusion
func (r *IntrafiExclusionService) Archive(ctx context.Context, intrafiExclusionID string, opts ...option.RequestOption) (res *IntrafiExclusion, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("intrafi_exclusions/%s/archive", intrafiExclusionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Certain institutions may be excluded per Entity when sweeping funds into the
// IntraFi network. This is useful when an Entity already has deposits at a
// particular bank, and does not want to sweep additional funds to it. It may take
// 5 business days for an exclusion to be processed.
type IntrafiExclusion struct {
	// The identifier of this exclusion request.
	ID string `json:"id,required"`
	// The name of the excluded institution.
	BankName string `json:"bank_name,required"`
	// The entity for which this institution is excluded.
	EntityID string `json:"entity_id,required"`
	// When this was exclusion was confirmed by IntraFi.
	ExcludedAt time.Time `json:"excluded_at,required,nullable" format:"date-time"`
	// The Federal Deposit Insurance Corporation's certificate number for the
	// institution.
	FdicCertificateNumber string `json:"fdic_certificate_number,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The status of the exclusion request.
	Status IntrafiExclusionStatus `json:"status,required"`
	// When this was exclusion was submitted to IntraFi by Increase.
	SubmittedAt time.Time `json:"submitted_at,required,nullable" format:"date-time"`
	// A constant representing the object's type. For this resource it will always be
	// `intrafi_exclusion`.
	Type IntrafiExclusionType `json:"type,required"`
	JSON intrafiExclusionJSON `json:"-"`
}

// intrafiExclusionJSON contains the JSON metadata for the struct
// [IntrafiExclusion]
type intrafiExclusionJSON struct {
	ID                    apijson.Field
	BankName              apijson.Field
	EntityID              apijson.Field
	ExcludedAt            apijson.Field
	FdicCertificateNumber apijson.Field
	IdempotencyKey        apijson.Field
	Status                apijson.Field
	SubmittedAt           apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IntrafiExclusion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the exclusion request.
type IntrafiExclusionStatus string

const (
	// The exclusion is being added to the IntraFi network.
	IntrafiExclusionStatusPending IntrafiExclusionStatus = "pending"
	// The exclusion has been added to the IntraFi network.
	IntrafiExclusionStatusCompleted IntrafiExclusionStatus = "completed"
	// The exclusion has been removed from the IntraFi network.
	IntrafiExclusionStatusArchived IntrafiExclusionStatus = "archived"
)

// A constant representing the object's type. For this resource it will always be
// `intrafi_exclusion`.
type IntrafiExclusionType string

const (
	IntrafiExclusionTypeIntrafiExclusion IntrafiExclusionType = "intrafi_exclusion"
)

type IntrafiExclusionNewParams struct {
	// The name of the financial institution to be excluded.
	BankName param.Field[string] `json:"bank_name,required"`
	// The identifier of the Entity whose deposits will be excluded.
	EntityID param.Field[string] `json:"entity_id,required"`
}

func (r IntrafiExclusionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntrafiExclusionListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter IntraFi Exclusions for those belonging to the specified Entity.
	EntityID param.Field[string] `query:"entity_id"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [IntrafiExclusionListParams]'s query parameters as
// `url.Values`.
func (r IntrafiExclusionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
