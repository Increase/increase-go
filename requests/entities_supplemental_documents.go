package requests

import (
	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
)

type SupplementalDocumentNewParams struct {
	// The identifier of the File containing the document.
	FileID field.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes SupplementalDocumentNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r SupplementalDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}
