package requests

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type FileNewParams struct {
	// The file contents. This should follow the specifications of
	// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
	// transfers for the multipart/form-data protocol.
	File field.Field[io.Reader] `form:"file,required" format:"binary"`
	// The description you choose to give the File.
	Description field.Field[string] `form:"description"`
	// What the File will be used for in Increase's systems.
	Purpose field.Field[FileNewParamsPurpose] `form:"purpose,required"`
}

func (r FileNewParams) MarshalMultipart() (data []byte, err error) {
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	{
		name := "anonymous_file"
		if nameable, ok := r.File.Value.(interface{ Name() string }); ok {
			name = nameable.Name()
		}
		part, err := writer.CreateFormFile("file", name)
		if err != nil {
			return nil, err
		}
		io.Copy(part, r.File.Value)
	}
	{
		bdy, err := pjson.Marshal(r.Description)
		if err != nil {
			return nil, err
		}
		writer.WriteField("description", string(bdy))
	}
	{
		bdy, err := pjson.Marshal(r.Purpose)
		if err != nil {
			return nil, err
		}
		writer.WriteField("purpose", string(bdy))
	}
	return body.Bytes(), nil
}

type FileNewParamsPurpose string

const (
	FileNewParamsPurposeCheckImageFront            FileNewParamsPurpose = "check_image_front"
	FileNewParamsPurposeCheckImageBack             FileNewParamsPurpose = "check_image_back"
	FileNewParamsPurposeFormSs_4                   FileNewParamsPurpose = "form_ss_4"
	FileNewParamsPurposeIdentityDocument           FileNewParamsPurpose = "identity_document"
	FileNewParamsPurposeOther                      FileNewParamsPurpose = "other"
	FileNewParamsPurposeTrustFormationDocument     FileNewParamsPurpose = "trust_formation_document"
	FileNewParamsPurposeDigitalWalletArtwork       FileNewParamsPurpose = "digital_wallet_artwork"
	FileNewParamsPurposeDigitalWalletAppIcon       FileNewParamsPurpose = "digital_wallet_app_icon"
	FileNewParamsPurposeEntitySupplementalDocument FileNewParamsPurpose = "entity_supplemental_document"
)

type FileListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     field.Field[int64]                   `query:"limit"`
	CreatedAt field.Field[FileListParamsCreatedAt] `query:"created_at"`
	Purpose   field.Field[FileListParamsPurpose]   `query:"purpose"`
}

// URLQuery serializes FileListParams into a url.Values of the query parameters
// associated with this value
func (r FileListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type FileListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes FileListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r FileListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type FileListParamsPurpose struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In field.Field[[]FileListParamsPurposeIn] `query:"in"`
}

// URLQuery serializes FileListParamsPurpose into a url.Values of the query
// parameters associated with this value
func (r FileListParamsPurpose) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type FileListParamsPurposeIn string

const (
	FileListParamsPurposeInCheckImageFront            FileListParamsPurposeIn = "check_image_front"
	FileListParamsPurposeInCheckImageBack             FileListParamsPurposeIn = "check_image_back"
	FileListParamsPurposeInForm_1099Int               FileListParamsPurposeIn = "form_1099_int"
	FileListParamsPurposeInFormSs_4                   FileListParamsPurposeIn = "form_ss_4"
	FileListParamsPurposeInIdentityDocument           FileListParamsPurposeIn = "identity_document"
	FileListParamsPurposeInIncreaseStatement          FileListParamsPurposeIn = "increase_statement"
	FileListParamsPurposeInOther                      FileListParamsPurposeIn = "other"
	FileListParamsPurposeInTrustFormationDocument     FileListParamsPurposeIn = "trust_formation_document"
	FileListParamsPurposeInDigitalWalletArtwork       FileListParamsPurposeIn = "digital_wallet_artwork"
	FileListParamsPurposeInDigitalWalletAppIcon       FileListParamsPurposeIn = "digital_wallet_app_icon"
	FileListParamsPurposeInEntitySupplementalDocument FileListParamsPurposeIn = "entity_supplemental_document"
	FileListParamsPurposeInExport                     FileListParamsPurposeIn = "export"
)
