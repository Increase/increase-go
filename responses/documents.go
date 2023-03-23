package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
)

type Document struct {
	// The Document identifier.
	ID string `json:"id,required"`
	// The type of document.
	Category DocumentCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Entity the document was generated for.
	EntityID string `json:"entity_id,required,nullable"`
	// The identifier of the File containing the Document's contents.
	FileID string `json:"file_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `document`.
	Type DocumentType `json:"type,required"`
	JSON DocumentJSON
}

type DocumentJSON struct {
	ID        pjson.Metadata
	Category  pjson.Metadata
	CreatedAt pjson.Metadata
	EntityID  pjson.Metadata
	FileID    pjson.Metadata
	Type      pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Document using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DocumentCategory string

const (
	DocumentCategoryForm_1099Int DocumentCategory = "form_1099_int"
)

type DocumentType string

const (
	DocumentTypeDocument DocumentType = "document"
)

type DocumentList struct {
	// The contents of the list.
	Data []Document `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       DocumentListJSON
}

type DocumentListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DocumentList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DocumentsPage struct {
	*pagination.Page[Document]
}

func (r *DocumentsPage) Document() *Document {
	return r.Current()
}

func (r *DocumentsPage) NextPage() (*DocumentsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &DocumentsPage{page}, nil
	}
}
