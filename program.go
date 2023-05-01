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

type ProgramService struct {
	Options []option.RequestOption
}

func NewProgramService(opts ...option.RequestOption) (r *ProgramService) {
	r = &ProgramService{}
	r.Options = opts
	return
}

// Retrieve a Program
func (r *ProgramService) Get(ctx context.Context, program_id string, opts ...option.RequestOption) (res *Program, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("programs/%s", program_id)
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

type Program struct {
	// The name of the Program.
	Name string `json:"name,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Program
	// was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Program
	// was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The Program identifier.
	ID string `json:"id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `program`.
	Type ProgramType `json:"type,required"`
	JSON ProgramJSON
}

type ProgramJSON struct {
	Name      apijson.Metadata
	CreatedAt apijson.Metadata
	UpdatedAt apijson.Metadata
	ID        apijson.Metadata
	Type      apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Program using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Program) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProgramType string

const (
	ProgramTypeProgram ProgramType = "program"
)

type ProgramListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes ProgramListParams into a url.Values of the query parameters
// associated with this value
func (r ProgramListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ProgramListResponse struct {
	// The contents of the list.
	Data []Program `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       ProgramListResponseJSON
}

type ProgramListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ProgramListResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ProgramListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
