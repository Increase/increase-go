package requests

import (
	"fmt"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
)

type CreateASupplementalDocumentForAnEntityParameters struct {
	// The identifier of the File containing the document.
	FileID fields.Field[string] `json:"file_id,required"`
}

// MarshalJSON serializes CreateASupplementalDocumentForAnEntityParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateASupplementalDocumentForAnEntityParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateASupplementalDocumentForAnEntityParameters) String() (result string) {
	return fmt.Sprintf("&CreateASupplementalDocumentForAnEntityParameters{FileID:%s}", r.FileID)
}
