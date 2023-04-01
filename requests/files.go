package requests

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/fields"
)

type CreateAFileParameters struct {
	// The file contents. This should follow the specifications of
	// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
	// transfers for the multipart/form-data protocol.
	File fields.Field[io.Reader] `form:"file,required" format:"binary"`
	// The description you choose to give the File.
	Description fields.Field[string] `form:"description"`
	// What the File will be used for in Increase's systems.
	Purpose fields.Field[CreateAFileParametersPurpose] `form:"purpose,required"`
}

func (r *CreateAFileParameters) MarshalMultipart() (data []byte, err error) {
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

func (r CreateAFileParameters) String() (result string) {
	return fmt.Sprintf("&CreateAFileParameters{File:%s Description:%s Purpose:%s}", r.File, r.Description, r.Purpose)
}

type CreateAFileParametersPurpose string

const (
	CreateAFileParametersPurposeCheckImageFront            CreateAFileParametersPurpose = "check_image_front"
	CreateAFileParametersPurposeCheckImageBack             CreateAFileParametersPurpose = "check_image_back"
	CreateAFileParametersPurposeFormSs_4                   CreateAFileParametersPurpose = "form_ss_4"
	CreateAFileParametersPurposeIdentityDocument           CreateAFileParametersPurpose = "identity_document"
	CreateAFileParametersPurposeOther                      CreateAFileParametersPurpose = "other"
	CreateAFileParametersPurposeTrustFormationDocument     CreateAFileParametersPurpose = "trust_formation_document"
	CreateAFileParametersPurposeDigitalWalletArtwork       CreateAFileParametersPurpose = "digital_wallet_artwork"
	CreateAFileParametersPurposeDigitalWalletAppIcon       CreateAFileParametersPurpose = "digital_wallet_app_icon"
	CreateAFileParametersPurposeEntitySupplementalDocument CreateAFileParametersPurpose = "entity_supplemental_document"
)

type FileListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     fields.Field[int64]                   `query:"limit"`
	CreatedAt fields.Field[FileListParamsCreatedAt] `query:"created_at"`
	Purpose   fields.Field[FileListParamsPurpose]   `query:"purpose"`
}

// URLQuery serializes FileListParams into a url.Values of the query parameters
// associated with this value
func (r *FileListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r FileListParams) String() (result string) {
	return fmt.Sprintf("&FileListParams{Cursor:%s Limit:%s CreatedAt:%s Purpose:%s}", r.Cursor, r.Limit, r.CreatedAt, r.Purpose)
}

type FileListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After fields.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before fields.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter fields.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore fields.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes FileListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r *FileListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r FileListParamsCreatedAt) String() (result string) {
	return fmt.Sprintf("&FileListParamsCreatedAt{After:%s Before:%s OnOrAfter:%s OnOrBefore:%s}", r.After, r.Before, r.OnOrAfter, r.OnOrBefore)
}

type FileListParamsPurpose struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]FileListParamsPurposeIn] `query:"in"`
}

// URLQuery serializes FileListParamsPurpose into a url.Values of the query
// parameters associated with this value
func (r *FileListParamsPurpose) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r FileListParamsPurpose) String() (result string) {
	return fmt.Sprintf("&FileListParamsPurpose{In:%s}", core.Fmt(r.In))
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
