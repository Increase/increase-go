// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// InboundWireTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundWireTransferService] method instead.
type InboundWireTransferService struct {
	Options []option.RequestOption
}

// NewInboundWireTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboundWireTransferService(opts ...option.RequestOption) (r *InboundWireTransferService) {
	r = &InboundWireTransferService{}
	r.Options = opts
	return
}

// Retrieve an Inbound Wire Transfer
func (r *InboundWireTransferService) Get(ctx context.Context, inboundWireTransferID string, opts ...option.RequestOption) (res *InboundWireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if inboundWireTransferID == "" {
		err = errors.New("missing required inbound_wire_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_wire_transfers/%s", inboundWireTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound Wire Transfers
func (r *InboundWireTransferService) List(ctx context.Context, query InboundWireTransferListParams, opts ...option.RequestOption) (res *pagination.Page[InboundWireTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_wire_transfers"
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

// List Inbound Wire Transfers
func (r *InboundWireTransferService) ListAutoPaging(ctx context.Context, query InboundWireTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundWireTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Reverse an Inbound Wire Transfer
func (r *InboundWireTransferService) Reverse(ctx context.Context, inboundWireTransferID string, body InboundWireTransferReverseParams, opts ...option.RequestOption) (res *InboundWireTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if inboundWireTransferID == "" {
		err = errors.New("missing required inbound_wire_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_wire_transfers/%s/reverse", inboundWireTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// An Inbound Wire Transfer is a wire transfer initiated outside of Increase to
// your account.
type InboundWireTransfer struct {
	// The inbound wire transfer's identifier.
	ID string `json:"id,required"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id,required"`
	// The identifier of the Account Number to which this transfer was sent.
	AccountNumberID string `json:"account_number_id,required"`
	// The amount in USD cents.
	Amount int64 `json:"amount,required"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine1 string `json:"beneficiary_address_line1,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine2 string `json:"beneficiary_address_line2,required,nullable"`
	// A free-form address field set by the sender.
	BeneficiaryAddressLine3 string `json:"beneficiary_address_line3,required,nullable"`
	// A name set by the sender.
	BeneficiaryName string `json:"beneficiary_name,required,nullable"`
	// A free-form reference string set by the sender, to help identify the transfer.
	BeneficiaryReference string `json:"beneficiary_reference,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the inbound wire transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// An Increase-constructed description of the transfer.
	Description string `json:"description,required"`
	// A unique identifier available to the originating and receiving banks, commonly
	// abbreviated as IMAD. It is created when the wire is submitted to the Fedwire
	// service and is helpful when debugging wires with the originating bank.
	InputMessageAccountabilityData string `json:"input_message_accountability_data,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine1 string `json:"originator_address_line1,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine2 string `json:"originator_address_line2,required,nullable"`
	// The address of the wire originator, set by the sending bank.
	OriginatorAddressLine3 string `json:"originator_address_line3,required,nullable"`
	// The originator of the wire, set by the sending bank.
	OriginatorName string `json:"originator_name,required,nullable"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number,required,nullable"`
	// An Increase-created concatenation of the Originator-to-Beneficiary lines.
	OriginatorToBeneficiaryInformation string `json:"originator_to_beneficiary_information,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine1 string `json:"originator_to_beneficiary_information_line1,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine2 string `json:"originator_to_beneficiary_information_line2,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine3 string `json:"originator_to_beneficiary_information_line3,required,nullable"`
	// A free-form message set by the wire originator.
	OriginatorToBeneficiaryInformationLine4 string `json:"originator_to_beneficiary_information_line4,required,nullable"`
	// Information about the reversal of the inbound wire transfer if it has been
	// reversed.
	Reversal InboundWireTransferReversal `json:"reversal,required,nullable"`
	// The sending bank's reference number for the wire transfer.
	SenderReference string `json:"sender_reference,required,nullable"`
	// The status of the transfer.
	Status InboundWireTransferStatus `json:"status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_wire_transfer`.
	Type InboundWireTransferType `json:"type,required"`
	JSON inboundWireTransferJSON `json:"-"`
}

// inboundWireTransferJSON contains the JSON metadata for the struct
// [InboundWireTransfer]
type inboundWireTransferJSON struct {
	ID                                      apijson.Field
	AccountID                               apijson.Field
	AccountNumberID                         apijson.Field
	Amount                                  apijson.Field
	BeneficiaryAddressLine1                 apijson.Field
	BeneficiaryAddressLine2                 apijson.Field
	BeneficiaryAddressLine3                 apijson.Field
	BeneficiaryName                         apijson.Field
	BeneficiaryReference                    apijson.Field
	CreatedAt                               apijson.Field
	Description                             apijson.Field
	InputMessageAccountabilityData          apijson.Field
	OriginatorAddressLine1                  apijson.Field
	OriginatorAddressLine2                  apijson.Field
	OriginatorAddressLine3                  apijson.Field
	OriginatorName                          apijson.Field
	OriginatorRoutingNumber                 apijson.Field
	OriginatorToBeneficiaryInformation      apijson.Field
	OriginatorToBeneficiaryInformationLine1 apijson.Field
	OriginatorToBeneficiaryInformationLine2 apijson.Field
	OriginatorToBeneficiaryInformationLine3 apijson.Field
	OriginatorToBeneficiaryInformationLine4 apijson.Field
	Reversal                                apijson.Field
	SenderReference                         apijson.Field
	Status                                  apijson.Field
	Type                                    apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *InboundWireTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundWireTransferJSON) RawJSON() string {
	return r.raw
}

// Information about the reversal of the inbound wire transfer if it has been
// reversed.
type InboundWireTransferReversal struct {
	// The reason for the reversal.
	Reason InboundWireTransferReversalReason `json:"reason,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was reversed.
	ReversedAt time.Time                       `json:"reversed_at,required" format:"date-time"`
	JSON       inboundWireTransferReversalJSON `json:"-"`
}

// inboundWireTransferReversalJSON contains the JSON metadata for the struct
// [InboundWireTransferReversal]
type inboundWireTransferReversalJSON struct {
	Reason      apijson.Field
	ReversedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundWireTransferReversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundWireTransferReversalJSON) RawJSON() string {
	return r.raw
}

// The reason for the reversal.
type InboundWireTransferReversalReason string

const (
	InboundWireTransferReversalReasonDuplicate       InboundWireTransferReversalReason = "duplicate"
	InboundWireTransferReversalReasonCreditorRequest InboundWireTransferReversalReason = "creditor_request"
)

func (r InboundWireTransferReversalReason) IsKnown() bool {
	switch r {
	case InboundWireTransferReversalReasonDuplicate, InboundWireTransferReversalReasonCreditorRequest:
		return true
	}
	return false
}

// The status of the transfer.
type InboundWireTransferStatus string

const (
	InboundWireTransferStatusPending  InboundWireTransferStatus = "pending"
	InboundWireTransferStatusAccepted InboundWireTransferStatus = "accepted"
	InboundWireTransferStatusDeclined InboundWireTransferStatus = "declined"
	InboundWireTransferStatusReversed InboundWireTransferStatus = "reversed"
)

func (r InboundWireTransferStatus) IsKnown() bool {
	switch r {
	case InboundWireTransferStatusPending, InboundWireTransferStatusAccepted, InboundWireTransferStatusDeclined, InboundWireTransferStatusReversed:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_wire_transfer`.
type InboundWireTransferType string

const (
	InboundWireTransferTypeInboundWireTransfer InboundWireTransferType = "inbound_wire_transfer"
)

func (r InboundWireTransferType) IsKnown() bool {
	switch r {
	case InboundWireTransferTypeInboundWireTransfer:
		return true
	}
	return false
}

type InboundWireTransferListParams struct {
	// Filter Inbound Wire Transfers to ones belonging to the specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound Wire Transfers to ones belonging to the specified Account Number.
	AccountNumberID param.Field[string]                                 `query:"account_number_id"`
	CreatedAt       param.Field[InboundWireTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                               `query:"limit"`
	Status param.Field[InboundWireTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [InboundWireTransferListParams]'s query parameters as
// `url.Values`.
func (r InboundWireTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundWireTransferListParamsCreatedAt struct {
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

// URLQuery serializes [InboundWireTransferListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r InboundWireTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundWireTransferListParamsStatus struct {
	// Filter Inbound Wire Transfers to those with the specified status. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]InboundWireTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [InboundWireTransferListParamsStatus]'s query parameters as
// `url.Values`.
func (r InboundWireTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundWireTransferListParamsStatusIn string

const (
	InboundWireTransferListParamsStatusInPending  InboundWireTransferListParamsStatusIn = "pending"
	InboundWireTransferListParamsStatusInAccepted InboundWireTransferListParamsStatusIn = "accepted"
	InboundWireTransferListParamsStatusInDeclined InboundWireTransferListParamsStatusIn = "declined"
	InboundWireTransferListParamsStatusInReversed InboundWireTransferListParamsStatusIn = "reversed"
)

func (r InboundWireTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case InboundWireTransferListParamsStatusInPending, InboundWireTransferListParamsStatusInAccepted, InboundWireTransferListParamsStatusInDeclined, InboundWireTransferListParamsStatusInReversed:
		return true
	}
	return false
}

type InboundWireTransferReverseParams struct {
	// Reason for the reversal.
	Reason param.Field[InboundWireTransferReverseParamsReason] `json:"reason,required"`
}

func (r InboundWireTransferReverseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reason for the reversal.
type InboundWireTransferReverseParamsReason string

const (
	InboundWireTransferReverseParamsReasonDuplicate       InboundWireTransferReverseParamsReason = "duplicate"
	InboundWireTransferReverseParamsReasonCreditorRequest InboundWireTransferReverseParamsReason = "creditor_request"
)

func (r InboundWireTransferReverseParamsReason) IsKnown() bool {
	switch r {
	case InboundWireTransferReverseParamsReasonDuplicate, InboundWireTransferReverseParamsReasonCreditorRequest:
		return true
	}
	return false
}
