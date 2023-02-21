package types

import (
	"fmt"
	"increase/core"
	"increase/core/pjson"
)

type CreateASupplementalDocumentForAnEntityParameters struct {
	// The identifier of the File containing the document.
	FileID     *string                `pjson:"file_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateASupplementalDocumentForAnEntityParameters using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *CreateASupplementalDocumentForAnEntityParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateASupplementalDocumentForAnEntityParameters into an
// array of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *CreateASupplementalDocumentForAnEntityParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier of the File containing the document.
func (r *CreateASupplementalDocumentForAnEntityParameters) GetFileID() (FileID string) {
	if r != nil && r.FileID != nil {
		FileID = *r.FileID
	}
	return
}

func (r CreateASupplementalDocumentForAnEntityParameters) String() (result string) {
	return fmt.Sprintf("&CreateASupplementalDocumentForAnEntityParameters{FileID:%s}", core.FmtP(r.FileID))
}
