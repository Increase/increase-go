// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apiform"
	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// FileService contains methods and other services that help with interacting with
// the increase API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a File
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Files
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *shared.Page[File], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "files"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Files
func (r *FileService) ListAutoPaging(ctx context.Context, query FileListParams, opts ...option.RequestOption) *shared.PageAutoPager[File] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Files are objects that represent a file hosted on Increase's servers. The file
// may have been uploaded by you (for example, when uploading a check image) or it
// may have been created by Increase (for example, an autogenerated statement PDF).
type File struct {
	// The File's identifier.
	ID string `json:"id,required"`
	// The time the File was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A description of the File.
	Description string `json:"description,required,nullable"`
	// Whether the File was generated by Increase or by you and sent to Increase.
	Direction FileDirection `json:"direction,required"`
	// A URL from where the File can be downloaded at this point in time. The location
	// of this URL may change over time.
	DownloadURL string `json:"download_url,required,nullable"`
	// The filename that was provided upon upload or generated by Increase.
	Filename string `json:"filename,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The MIME type of the file.
	MimeType string `json:"mime_type,required"`
	// What the File will be used for. We may add additional possible values for this
	// enum over time; your application should be able to handle such additions
	// gracefully.
	Purpose FilePurpose `json:"purpose,required"`
	// A constant representing the object's type. For this resource it will always be
	// `file`.
	Type FileType `json:"type,required"`
	JSON fileJSON `json:"-"`
}

// fileJSON contains the JSON metadata for the struct [File]
type fileJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	Direction      apijson.Field
	DownloadURL    apijson.Field
	Filename       apijson.Field
	IdempotencyKey apijson.Field
	MimeType       apijson.Field
	Purpose        apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *File) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the File was generated by Increase or by you and sent to Increase.
type FileDirection string

const (
	// This File was sent by you to Increase.
	FileDirectionToIncrease FileDirection = "to_increase"
	// This File was generated by Increase.
	FileDirectionFromIncrease FileDirection = "from_increase"
)

// What the File will be used for. We may add additional possible values for this
// enum over time; your application should be able to handle such additions
// gracefully.
type FilePurpose string

const (
	// An image of the front of a check, used for check deposits.
	FilePurposeCheckImageFront FilePurpose = "check_image_front"
	// An image of the back of a check, used for check deposits.
	FilePurposeCheckImageBack FilePurpose = "check_image_back"
	// An image of a check that was mailed to a recipient.
	FilePurposeMailedCheckImage FilePurpose = "mailed_check_image"
	// IRS Form 1099-INT.
	FilePurposeForm1099Int FilePurpose = "form_1099_int"
	// IRS Form SS-4.
	FilePurposeFormSS4 FilePurpose = "form_ss_4"
	// An image of a government-issued ID.
	FilePurposeIdentityDocument FilePurpose = "identity_document"
	// A statement generated by Increase.
	FilePurposeIncreaseStatement FilePurpose = "increase_statement"
	// A file purpose not covered by any of the other cases.
	FilePurposeOther FilePurpose = "other"
	// A legal document forming a trust.
	FilePurposeTrustFormationDocument FilePurpose = "trust_formation_document"
	// A card image to be rendered inside digital wallet apps. This must be a 1536x969
	// pixel PNG.
	FilePurposeDigitalWalletArtwork FilePurpose = "digital_wallet_artwork"
	// An icon for you app to be rendered inside digital wallet apps. This must be a
	// 100x100 pixel PNG.
	FilePurposeDigitalWalletAppIcon FilePurpose = "digital_wallet_app_icon"
	// A card image to be printed on the front of a physical card. This must be a
	// 2100x1340 pixel PNG with no other color but black.
	FilePurposePhysicalCardFront FilePurpose = "physical_card_front"
	// The image to be printed on the back of a physical card.
	FilePurposePhysicalCardBack FilePurpose = "physical_card_back"
	// An image representing the entirety of the carrier used for a physical card. This
	// must be a 2550x3300 pixel PNG with no other color but black.
	FilePurposePhysicalCardCarrier FilePurpose = "physical_card_carrier"
	// A document requested by Increase.
	FilePurposeDocumentRequest FilePurpose = "document_request"
	// A supplemental document associated an an Entity.
	FilePurposeEntitySupplementalDocument FilePurpose = "entity_supplemental_document"
	// The results of an Export you requested via the dashboard or API.
	FilePurposeExport FilePurpose = "export"
	// An attachment to an Unusual Activity Report.
	FilePurposeUnusualActivityReportAttachment FilePurpose = "unusual_activity_report_attachment"
)

// A constant representing the object's type. For this resource it will always be
// `file`.
type FileType string

const (
	FileTypeFile FileType = "file"
)

type FileNewParams struct {
	// The file contents. This should follow the specifications of
	// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
	// transfers for the multipart/form-data protocol.
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// What the File will be used for in Increase's systems.
	Purpose param.Field[FileNewParamsPurpose] `json:"purpose,required"`
	// The description you choose to give the File.
	Description param.Field[string] `json:"description"`
}

func (r FileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// What the File will be used for in Increase's systems.
type FileNewParamsPurpose string

const (
	// An image of the front of a check, used for check deposits.
	FileNewParamsPurposeCheckImageFront FileNewParamsPurpose = "check_image_front"
	// An image of the back of a check, used for check deposits.
	FileNewParamsPurposeCheckImageBack FileNewParamsPurpose = "check_image_back"
	// An image of a check that was mailed to a recipient.
	FileNewParamsPurposeMailedCheckImage FileNewParamsPurpose = "mailed_check_image"
	// IRS Form SS-4.
	FileNewParamsPurposeFormSS4 FileNewParamsPurpose = "form_ss_4"
	// An image of a government-issued ID.
	FileNewParamsPurposeIdentityDocument FileNewParamsPurpose = "identity_document"
	// A file purpose not covered by any of the other cases.
	FileNewParamsPurposeOther FileNewParamsPurpose = "other"
	// A legal document forming a trust.
	FileNewParamsPurposeTrustFormationDocument FileNewParamsPurpose = "trust_formation_document"
	// A card image to be rendered inside digital wallet apps. This must be a 1536x969
	// pixel PNG.
	FileNewParamsPurposeDigitalWalletArtwork FileNewParamsPurpose = "digital_wallet_artwork"
	// An icon for you app to be rendered inside digital wallet apps. This must be a
	// 100x100 pixel PNG.
	FileNewParamsPurposeDigitalWalletAppIcon FileNewParamsPurpose = "digital_wallet_app_icon"
	// A card image to be printed on the front of a physical card. This must be a
	// 2100x1340 pixel PNG with no other color but black.
	FileNewParamsPurposePhysicalCardFront FileNewParamsPurpose = "physical_card_front"
	// An image representing the entirety of the carrier used for a physical card. This
	// must be a 2550x3300 pixel PNG with no other color but black.
	FileNewParamsPurposePhysicalCardCarrier FileNewParamsPurpose = "physical_card_carrier"
	// A document requested by Increase.
	FileNewParamsPurposeDocumentRequest FileNewParamsPurpose = "document_request"
	// A supplemental document associated an an Entity.
	FileNewParamsPurposeEntitySupplementalDocument FileNewParamsPurpose = "entity_supplemental_document"
	// An attachment to an Unusual Activity Report.
	FileNewParamsPurposeUnusualActivityReportAttachment FileNewParamsPurpose = "unusual_activity_report_attachment"
)

type FileListParams struct {
	CreatedAt param.Field[FileListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit   param.Field[int64]                 `query:"limit"`
	Purpose param.Field[FileListParamsPurpose] `query:"purpose"`
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FileListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [FileListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r FileListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FileListParamsPurpose struct {
	// Filter Files for those with the specified purpose or purposes. For GET requests,
	// this should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]FileListParamsPurposeIn] `query:"in"`
}

// URLQuery serializes [FileListParamsPurpose]'s query parameters as `url.Values`.
func (r FileListParamsPurpose) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FileListParamsPurposeIn string

const (
	// An image of the front of a check, used for check deposits.
	FileListParamsPurposeInCheckImageFront FileListParamsPurposeIn = "check_image_front"
	// An image of the back of a check, used for check deposits.
	FileListParamsPurposeInCheckImageBack FileListParamsPurposeIn = "check_image_back"
	// An image of a check that was mailed to a recipient.
	FileListParamsPurposeInMailedCheckImage FileListParamsPurposeIn = "mailed_check_image"
	// IRS Form 1099-INT.
	FileListParamsPurposeInForm1099Int FileListParamsPurposeIn = "form_1099_int"
	// IRS Form SS-4.
	FileListParamsPurposeInFormSS4 FileListParamsPurposeIn = "form_ss_4"
	// An image of a government-issued ID.
	FileListParamsPurposeInIdentityDocument FileListParamsPurposeIn = "identity_document"
	// A statement generated by Increase.
	FileListParamsPurposeInIncreaseStatement FileListParamsPurposeIn = "increase_statement"
	// A file purpose not covered by any of the other cases.
	FileListParamsPurposeInOther FileListParamsPurposeIn = "other"
	// A legal document forming a trust.
	FileListParamsPurposeInTrustFormationDocument FileListParamsPurposeIn = "trust_formation_document"
	// A card image to be rendered inside digital wallet apps. This must be a 1536x969
	// pixel PNG.
	FileListParamsPurposeInDigitalWalletArtwork FileListParamsPurposeIn = "digital_wallet_artwork"
	// An icon for you app to be rendered inside digital wallet apps. This must be a
	// 100x100 pixel PNG.
	FileListParamsPurposeInDigitalWalletAppIcon FileListParamsPurposeIn = "digital_wallet_app_icon"
	// A card image to be printed on the front of a physical card. This must be a
	// 2100x1340 pixel PNG with no other color but black.
	FileListParamsPurposeInPhysicalCardFront FileListParamsPurposeIn = "physical_card_front"
	// The image to be printed on the back of a physical card.
	FileListParamsPurposeInPhysicalCardBack FileListParamsPurposeIn = "physical_card_back"
	// An image representing the entirety of the carrier used for a physical card. This
	// must be a 2550x3300 pixel PNG with no other color but black.
	FileListParamsPurposeInPhysicalCardCarrier FileListParamsPurposeIn = "physical_card_carrier"
	// A document requested by Increase.
	FileListParamsPurposeInDocumentRequest FileListParamsPurposeIn = "document_request"
	// A supplemental document associated an an Entity.
	FileListParamsPurposeInEntitySupplementalDocument FileListParamsPurposeIn = "entity_supplemental_document"
	// The results of an Export you requested via the dashboard or API.
	FileListParamsPurposeInExport FileListParamsPurposeIn = "export"
	// An attachment to an Unusual Activity Report.
	FileListParamsPurposeInUnusualActivityReportAttachment FileListParamsPurposeIn = "unusual_activity_report_attachment"
)
