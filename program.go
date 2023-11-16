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

// ProgramService contains methods and other services that help with interacting
// with the increase API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewProgramService] method instead.
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
	path := fmt.Sprintf("programs/%s", programID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Programs
func (r *ProgramService) List(ctx context.Context, query ProgramListParams, opts ...option.RequestOption) (res *shared.Page[Program], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *ProgramService) ListAutoPaging(ctx context.Context, query ProgramListParams, opts ...option.RequestOption) *shared.PageAutoPager[Program] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Programs determine the compliance and commercial terms of Accounts. By default,
// you have a Commercial Banking program for managing your own funds. If you are
// lending or managing funds on behalf of your customers, or otherwise engaged in
// regulated activity, we will work together to create additional Programs for you.
type Program struct {
	// The Program identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Program
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
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
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Program) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A constant representing the object's type. For this resource it will always be
// `program`.
type ProgramType string

const (
	ProgramTypeProgram ProgramType = "program"
)

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
