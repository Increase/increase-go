package files

type FilesListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     *int64                   `json:"limit,omitempty"`
	CreatedAt *FilesCreatedAtListQuery `json:"created_at,omitempty"`
	Purpose   *FilesPurposeListQuery   `json:"purpose,omitempty"`
}

// Return the page of entries after this one.
func (r *FilesListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *FilesListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *FilesListQuery) GetCreatedAt() (CreatedAt FilesCreatedAtListQuery) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *FilesListQuery) GetPurpose() (Purpose FilesPurposeListQuery) {
	if r != nil && r.Purpose != nil {
		Purpose = *r.Purpose
	}
	return
}
