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
	"github.com/Increase/increase-go/internal/pagination"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// ProgramService contains methods and other services that help with interacting
// with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProgramService] method instead.
type ProgramService struct {
	Options []option.RequestOption
}

// NewProgramService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProgramService(opts ...option.RequestOption) (r *ProgramService) {
	r = &ProgramService{}
	r.Options = opts
	return
}

// Retrieve a Program
func (r *ProgramService) Get(ctx context.Context, programID string, opts ...option.RequestOption) (res *Program, err error) {
	opts = append(r.Options[:], opts...)
	if programID == "" {
		err = errors.New("missing required program_id parameter")
		return
	}
	path := fmt.Sprintf("programs/%s", programID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Programs
func (r *ProgramService) List(ctx context.Context, query ProgramListParams, opts ...option.RequestOption) (res *pagination.Page[Program], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "programs"
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

// List Programs
func (r *ProgramService) ListAutoPaging(ctx context.Context, query ProgramListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Program] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Programs determine the compliance and commercial terms of Accounts. By default,
// you have a Commercial Banking program for managing your own funds. If you are
// lending or managing funds on behalf of your customers, or otherwise engaged in
// regulated activity, we will work together to create additional Programs for you.
type Program struct {
	// The Program identifier.
	ID string `json:"id,required"`
	// The Bank the Program is with.
	Bank ProgramBank `json:"bank,required"`
	// The Program billing account.
	BillingAccountID string `json:"billing_account_id,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Program
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The default configuration for digital cards attached to this Program.
	DefaultDigitalCardProfileID string `json:"default_digital_card_profile_id,required,nullable"`
	// The Interest Rate currently being earned on the accounts in this program, as a
	// string containing a decimal number. For example, a 1% interest rate would be
	// represented as "0.01".
	InterestRate string `json:"interest_rate,required"`
	// The name of the Program.
	Name string `json:"name,required"`
	// A constant representing the object's type. For this resource it will always be
	// `program`.
	Type ProgramType `json:"type,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Program
	// was last updated.
	UpdatedAt time.Time   `json:"updated_at,required" format:"date-time"`
	JSON      programJSON `json:"-"`
}

// programJSON contains the JSON metadata for the struct [Program]
type programJSON struct {
	ID                          apijson.Field
	Bank                        apijson.Field
	BillingAccountID            apijson.Field
	CreatedAt                   apijson.Field
	DefaultDigitalCardProfileID apijson.Field
	InterestRate                apijson.Field
	Name                        apijson.Field
	Type                        apijson.Field
	UpdatedAt                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Program) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r programJSON) RawJSON() string {
	return r.raw
}

// The Bank the Program is with.
type ProgramBank string

const (
	// Blue Ridge Bank, N.A.
	ProgramBankBlueRidgeBank ProgramBank = "blue_ridge_bank"
	// First Internet Bank of Indiana
	ProgramBankFirstInternetBank ProgramBank = "first_internet_bank"
	// Grasshopper Bank
	ProgramBankGrasshopperBank ProgramBank = "grasshopper_bank"
)

func (r ProgramBank) IsKnown() bool {
	switch r {
	case ProgramBankBlueRidgeBank, ProgramBankFirstInternetBank, ProgramBankGrasshopperBank:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `program`.
type ProgramType string

const (
	ProgramTypeProgram ProgramType = "program"
)

func (r ProgramType) IsKnown() bool {
	switch r {
	case ProgramTypeProgram:
		return true
	}
	return false
}

type ProgramListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ProgramListParams]'s query parameters as `url.Values`.
func (r ProgramListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
