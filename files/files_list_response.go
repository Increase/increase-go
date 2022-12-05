package files

type FilesListResponse struct {
	// The contents of the list.
	Data *[]File `json:"data,omitempty"`

	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor,omitempty"`
}

// The contents of the list.
func (r *FilesListResponse) GetData() (Data []File) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *FilesListResponse) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}
