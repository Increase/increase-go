package files

type FilesCreateParams struct {
	// The file contents. This should follow the specifications of
	// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
	// transfers for the multipart/form-data protocol.
	File *string `json:"file,omitempty"`

	// The description you choose to give the File.
	Description *string `json:"description,omitempty"`

	// What the File will be used for in Increase's systems.
	Purpose *string `json:"purpose,omitempty"`
}

// The file contents. This should follow the specifications of
// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
// transfers for the multipart/form-data protocol.
func (r *FilesCreateParams) GetFile() (File string) {
	if r != nil && r.File != nil {
		File = *r.File
	}
	return
}

// The description you choose to give the File.
func (r *FilesCreateParams) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// What the File will be used for in Increase's systems.
func (r *FilesCreateParams) GetPurpose() (Purpose string) {
	if r != nil && r.Purpose != nil {
		Purpose = *r.Purpose
	}
	return
}
