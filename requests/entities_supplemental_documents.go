package requests

import (
	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
)

type EntitySupplementalDocumentNewParams struct {
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes EntitySupplementalDocumentNewParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r EntitySupplementalDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
