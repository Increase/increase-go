package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

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
