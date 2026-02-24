// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/apiquery"
	"github.com/Increase/increase-go/internal/param"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
	"github.com/Increase/increase-go/packages/pagination"
)

// InboundACHTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboundACHTransferService] method instead.
type InboundACHTransferService struct {
	Options []option.RequestOption
}

// NewInboundACHTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInboundACHTransferService(opts ...option.RequestOption) (r *InboundACHTransferService) {
	r = &InboundACHTransferService{}
	r.Options = opts
	return
}

// Retrieve an Inbound ACH Transfer
func (r *InboundACHTransferService) Get(ctx context.Context, inboundACHTransferID string, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundACHTransferID == "" {
		err = errors.New("missing required inbound_ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_ach_transfers/%s", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inbound ACH Transfers
func (r *InboundACHTransferService) List(ctx context.Context, query InboundACHTransferListParams, opts ...option.RequestOption) (res *pagination.Page[InboundACHTransfer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "inbound_ach_transfers"
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

// List Inbound ACH Transfers
func (r *InboundACHTransferService) ListAutoPaging(ctx context.Context, query InboundACHTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InboundACHTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create a notification of change for an Inbound ACH Transfer
func (r *InboundACHTransferService) NewNotificationOfChange(ctx context.Context, inboundACHTransferID string, body InboundACHTransferNewNotificationOfChangeParams, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundACHTransferID == "" {
		err = errors.New("missing required inbound_ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_ach_transfers/%s/create_notification_of_change", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Decline an Inbound ACH Transfer
func (r *InboundACHTransferService) Decline(ctx context.Context, inboundACHTransferID string, body InboundACHTransferDeclineParams, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundACHTransferID == "" {
		err = errors.New("missing required inbound_ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_ach_transfers/%s/decline", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return an Inbound ACH Transfer
func (r *InboundACHTransferService) TransferReturn(ctx context.Context, inboundACHTransferID string, body InboundACHTransferTransferReturnParams, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboundACHTransferID == "" {
		err = errors.New("missing required inbound_ach_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("inbound_ach_transfers/%s/transfer_return", inboundACHTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// An Inbound ACH Transfer is an ACH transfer initiated outside of Increase to your
// account.
type InboundACHTransfer struct {
	// The inbound ACH transfer's identifier.
	ID string `json:"id" api:"required"`
	// If your transfer is accepted, this will contain details of the acceptance.
	Acceptance InboundACHTransferAcceptance `json:"acceptance" api:"required,nullable"`
	// The Account to which the transfer belongs.
	AccountID string `json:"account_id" api:"required"`
	// The identifier of the Account Number to which this transfer was sent.
	AccountNumberID string `json:"account_number_id" api:"required"`
	// Additional information sent from the originator.
	Addenda InboundACHTransferAddenda `json:"addenda" api:"required,nullable"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount" api:"required"`
	// The time at which the transfer will be automatically resolved.
	AutomaticallyResolvesAt time.Time `json:"automatically_resolves_at" api:"required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the inbound ACH transfer was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// If your transfer is declined, this will contain details of the decline.
	Decline InboundACHTransferDecline `json:"decline" api:"required,nullable"`
	// The direction of the transfer.
	Direction InboundACHTransferDirection `json:"direction" api:"required"`
	// The effective date of the transfer. This is sent by the sending bank and is a
	// factor in determining funds availability.
	EffectiveDate time.Time `json:"effective_date" api:"required" format:"date"`
	// If the Inbound ACH Transfer has a Standard Entry Class Code of IAT, this will
	// contain fields pertaining to the International ACH Transaction.
	InternationalAddenda InboundACHTransferInternationalAddenda `json:"international_addenda" api:"required,nullable"`
	// If you initiate a notification of change in response to the transfer, this will
	// contain its details.
	NotificationOfChange InboundACHTransferNotificationOfChange `json:"notification_of_change" api:"required,nullable"`
	// The descriptive date of the transfer.
	OriginatorCompanyDescriptiveDate string `json:"originator_company_descriptive_date" api:"required,nullable"`
	// The additional information included with the transfer.
	OriginatorCompanyDiscretionaryData string `json:"originator_company_discretionary_data" api:"required,nullable"`
	// The description of the transfer.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description" api:"required"`
	// The id of the company that initiated the transfer.
	OriginatorCompanyID string `json:"originator_company_id" api:"required"`
	// The name of the company that initiated the transfer.
	OriginatorCompanyName string `json:"originator_company_name" api:"required"`
	// The American Banking Association (ABA) routing number of the bank originating
	// the transfer.
	OriginatorRoutingNumber string `json:"originator_routing_number" api:"required"`
	// The id of the receiver of the transfer.
	ReceiverIDNumber string `json:"receiver_id_number" api:"required,nullable"`
	// The name of the receiver of the transfer.
	ReceiverName string `json:"receiver_name" api:"required,nullable"`
	// A subhash containing information about when and how the transfer settled at the
	// Federal Reserve.
	Settlement InboundACHTransferSettlement `json:"settlement" api:"required"`
	// The Standard Entry Class (SEC) code of the transfer.
	StandardEntryClassCode InboundACHTransferStandardEntryClassCode `json:"standard_entry_class_code" api:"required"`
	// The status of the transfer.
	Status InboundACHTransferStatus `json:"status" api:"required"`
	// A 15 digit number set by the sending bank and transmitted to the receiving bank.
	// Along with the amount, date, and originating routing number, this can be used to
	// identify the ACH transfer. ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach-returns#ach-returns).
	TraceNumber string `json:"trace_number" api:"required"`
	// If your transfer is returned, this will contain details of the return.
	TransferReturn InboundACHTransferTransferReturn `json:"transfer_return" api:"required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_ach_transfer`.
	Type        InboundACHTransferType `json:"type" api:"required"`
	ExtraFields map[string]interface{} `json:"-" api:"extrafields"`
	JSON        inboundACHTransferJSON `json:"-"`
}

// inboundACHTransferJSON contains the JSON metadata for the struct
// [InboundACHTransfer]
type inboundACHTransferJSON struct {
	ID                                 apijson.Field
	Acceptance                         apijson.Field
	AccountID                          apijson.Field
	AccountNumberID                    apijson.Field
	Addenda                            apijson.Field
	Amount                             apijson.Field
	AutomaticallyResolvesAt            apijson.Field
	CreatedAt                          apijson.Field
	Decline                            apijson.Field
	Direction                          apijson.Field
	EffectiveDate                      apijson.Field
	InternationalAddenda               apijson.Field
	NotificationOfChange               apijson.Field
	OriginatorCompanyDescriptiveDate   apijson.Field
	OriginatorCompanyDiscretionaryData apijson.Field
	OriginatorCompanyEntryDescription  apijson.Field
	OriginatorCompanyID                apijson.Field
	OriginatorCompanyName              apijson.Field
	OriginatorRoutingNumber            apijson.Field
	ReceiverIDNumber                   apijson.Field
	ReceiverName                       apijson.Field
	Settlement                         apijson.Field
	StandardEntryClassCode             apijson.Field
	Status                             apijson.Field
	TraceNumber                        apijson.Field
	TransferReturn                     apijson.Field
	Type                               apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *InboundACHTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferJSON) RawJSON() string {
	return r.raw
}

// If your transfer is accepted, this will contain details of the acceptance.
type InboundACHTransferAcceptance struct {
	// The time at which the transfer was accepted.
	AcceptedAt time.Time `json:"accepted_at" api:"required" format:"date-time"`
	// The id of the transaction for the accepted transfer.
	TransactionID string                           `json:"transaction_id" api:"required"`
	JSON          inboundACHTransferAcceptanceJSON `json:"-"`
}

// inboundACHTransferAcceptanceJSON contains the JSON metadata for the struct
// [InboundACHTransferAcceptance]
type inboundACHTransferAcceptanceJSON struct {
	AcceptedAt    apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundACHTransferAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferAcceptanceJSON) RawJSON() string {
	return r.raw
}

// Additional information sent from the originator.
type InboundACHTransferAddenda struct {
	// The type of addendum.
	Category InboundACHTransferAddendaCategory `json:"category" api:"required"`
	// Unstructured `payment_related_information` passed through by the originator.
	Freeform InboundACHTransferAddendaFreeform `json:"freeform" api:"required,nullable"`
	JSON     inboundACHTransferAddendaJSON     `json:"-"`
}

// inboundACHTransferAddendaJSON contains the JSON metadata for the struct
// [InboundACHTransferAddenda]
type inboundACHTransferAddendaJSON struct {
	Category    apijson.Field
	Freeform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundACHTransferAddenda) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferAddendaJSON) RawJSON() string {
	return r.raw
}

// The type of addendum.
type InboundACHTransferAddendaCategory string

const (
	InboundACHTransferAddendaCategoryFreeform InboundACHTransferAddendaCategory = "freeform"
)

func (r InboundACHTransferAddendaCategory) IsKnown() bool {
	switch r {
	case InboundACHTransferAddendaCategoryFreeform:
		return true
	}
	return false
}

// Unstructured `payment_related_information` passed through by the originator.
type InboundACHTransferAddendaFreeform struct {
	// Each entry represents an addendum received from the originator.
	Entries []InboundACHTransferAddendaFreeformEntry `json:"entries" api:"required"`
	JSON    inboundACHTransferAddendaFreeformJSON    `json:"-"`
}

// inboundACHTransferAddendaFreeformJSON contains the JSON metadata for the struct
// [InboundACHTransferAddendaFreeform]
type inboundACHTransferAddendaFreeformJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InboundACHTransferAddendaFreeform) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferAddendaFreeformJSON) RawJSON() string {
	return r.raw
}

type InboundACHTransferAddendaFreeformEntry struct {
	// The payment related information passed in the addendum.
	PaymentRelatedInformation string                                     `json:"payment_related_information" api:"required"`
	JSON                      inboundACHTransferAddendaFreeformEntryJSON `json:"-"`
}

// inboundACHTransferAddendaFreeformEntryJSON contains the JSON metadata for the
// struct [InboundACHTransferAddendaFreeformEntry]
type inboundACHTransferAddendaFreeformEntryJSON struct {
	PaymentRelatedInformation apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InboundACHTransferAddendaFreeformEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferAddendaFreeformEntryJSON) RawJSON() string {
	return r.raw
}

// If your transfer is declined, this will contain details of the decline.
type InboundACHTransferDecline struct {
	// The time at which the transfer was declined.
	DeclinedAt time.Time `json:"declined_at" api:"required" format:"date-time"`
	// The id of the transaction for the declined transfer.
	DeclinedTransactionID string `json:"declined_transaction_id" api:"required"`
	// The reason for the transfer decline.
	Reason InboundACHTransferDeclineReason `json:"reason" api:"required"`
	JSON   inboundACHTransferDeclineJSON   `json:"-"`
}

// inboundACHTransferDeclineJSON contains the JSON metadata for the struct
// [InboundACHTransferDecline]
type inboundACHTransferDeclineJSON struct {
	DeclinedAt            apijson.Field
	DeclinedTransactionID apijson.Field
	Reason                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InboundACHTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// The reason for the transfer decline.
type InboundACHTransferDeclineReason string

const (
	InboundACHTransferDeclineReasonACHRouteCanceled                                            InboundACHTransferDeclineReason = "ach_route_canceled"
	InboundACHTransferDeclineReasonACHRouteDisabled                                            InboundACHTransferDeclineReason = "ach_route_disabled"
	InboundACHTransferDeclineReasonBreachesLimit                                               InboundACHTransferDeclineReason = "breaches_limit"
	InboundACHTransferDeclineReasonEntityNotActive                                             InboundACHTransferDeclineReason = "entity_not_active"
	InboundACHTransferDeclineReasonGroupLocked                                                 InboundACHTransferDeclineReason = "group_locked"
	InboundACHTransferDeclineReasonTransactionNotAllowed                                       InboundACHTransferDeclineReason = "transaction_not_allowed"
	InboundACHTransferDeclineReasonUserInitiated                                               InboundACHTransferDeclineReason = "user_initiated"
	InboundACHTransferDeclineReasonInsufficientFunds                                           InboundACHTransferDeclineReason = "insufficient_funds"
	InboundACHTransferDeclineReasonReturnedPerOdfiRequest                                      InboundACHTransferDeclineReason = "returned_per_odfi_request"
	InboundACHTransferDeclineReasonAuthorizationRevokedByCustomer                              InboundACHTransferDeclineReason = "authorization_revoked_by_customer"
	InboundACHTransferDeclineReasonPaymentStopped                                              InboundACHTransferDeclineReason = "payment_stopped"
	InboundACHTransferDeclineReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   InboundACHTransferDeclineReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundACHTransferDeclineReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferDeclineReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	InboundACHTransferDeclineReasonBeneficiaryOrAccountHolderDeceased                          InboundACHTransferDeclineReason = "beneficiary_or_account_holder_deceased"
	InboundACHTransferDeclineReasonCreditEntryRefusedByReceiver                                InboundACHTransferDeclineReason = "credit_entry_refused_by_receiver"
	InboundACHTransferDeclineReasonDuplicateEntry                                              InboundACHTransferDeclineReason = "duplicate_entry"
	InboundACHTransferDeclineReasonCorporateCustomerAdvisedNotAuthorized                       InboundACHTransferDeclineReason = "corporate_customer_advised_not_authorized"
)

func (r InboundACHTransferDeclineReason) IsKnown() bool {
	switch r {
	case InboundACHTransferDeclineReasonACHRouteCanceled, InboundACHTransferDeclineReasonACHRouteDisabled, InboundACHTransferDeclineReasonBreachesLimit, InboundACHTransferDeclineReasonEntityNotActive, InboundACHTransferDeclineReasonGroupLocked, InboundACHTransferDeclineReasonTransactionNotAllowed, InboundACHTransferDeclineReasonUserInitiated, InboundACHTransferDeclineReasonInsufficientFunds, InboundACHTransferDeclineReasonReturnedPerOdfiRequest, InboundACHTransferDeclineReasonAuthorizationRevokedByCustomer, InboundACHTransferDeclineReasonPaymentStopped, InboundACHTransferDeclineReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, InboundACHTransferDeclineReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, InboundACHTransferDeclineReasonBeneficiaryOrAccountHolderDeceased, InboundACHTransferDeclineReasonCreditEntryRefusedByReceiver, InboundACHTransferDeclineReasonDuplicateEntry, InboundACHTransferDeclineReasonCorporateCustomerAdvisedNotAuthorized:
		return true
	}
	return false
}

// The direction of the transfer.
type InboundACHTransferDirection string

const (
	InboundACHTransferDirectionCredit InboundACHTransferDirection = "credit"
	InboundACHTransferDirectionDebit  InboundACHTransferDirection = "debit"
)

func (r InboundACHTransferDirection) IsKnown() bool {
	switch r {
	case InboundACHTransferDirectionCredit, InboundACHTransferDirectionDebit:
		return true
	}
	return false
}

// If the Inbound ACH Transfer has a Standard Entry Class Code of IAT, this will
// contain fields pertaining to the International ACH Transaction.
type InboundACHTransferInternationalAddenda struct {
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the destination country.
	DestinationCountryCode string `json:"destination_country_code" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// destination bank account.
	DestinationCurrencyCode string `json:"destination_currency_code" api:"required"`
	// A description of how the foreign exchange rate was calculated.
	ForeignExchangeIndicator InboundACHTransferInternationalAddendaForeignExchangeIndicator `json:"foreign_exchange_indicator" api:"required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference" api:"required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator" api:"required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount" api:"required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number" api:"required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode InboundACHTransferInternationalAddendaInternationalTransactionTypeCode `json:"international_transaction_type_code" api:"required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// originating bank account.
	OriginatingCurrencyCode string `json:"originating_currency_code" api:"required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the originating branch country.
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country" api:"required"`
	// An identifier for the originating bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	OriginatingDepositoryFinancialInstitutionID string `json:"originating_depository_financial_institution_id" api:"required"`
	// An instruction of how to interpret the
	// `originating_depository_financial_institution_id` field for this Transaction.
	OriginatingDepositoryFinancialInstitutionIDQualifier InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier" api:"required"`
	// The name of the originating bank. Sometimes this will refer to an American bank
	// and obscure the correspondent foreign bank.
	OriginatingDepositoryFinancialInstitutionName string `json:"originating_depository_financial_institution_name" api:"required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorCity string `json:"originator_city" api:"required"`
	// A portion of the originator address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the originator country.
	OriginatorCountry string `json:"originator_country" api:"required"`
	// An identifier for the originating company. This is generally stable across
	// multiple ACH transfers.
	OriginatorIdentification string `json:"originator_identification" api:"required"`
	// Either the name of the originator or an intermediary money transmitter.
	OriginatorName string `json:"originator_name" api:"required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorPostalCode string `json:"originator_postal_code" api:"required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStateOrProvince string `json:"originator_state_or_province" api:"required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStreetAddress string `json:"originator_street_address" api:"required"`
	// A description field set by the originator.
	PaymentRelatedInformation string `json:"payment_related_information" api:"required,nullable"`
	// A description field set by the originator.
	PaymentRelatedInformation2 string `json:"payment_related_information2" api:"required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverCity string `json:"receiver_city" api:"required"`
	// A portion of the receiver address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the receiver country.
	ReceiverCountry string `json:"receiver_country" api:"required"`
	// An identification number the originator uses for the receiver.
	ReceiverIdentificationNumber string `json:"receiver_identification_number" api:"required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverPostalCode string `json:"receiver_postal_code" api:"required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStateOrProvince string `json:"receiver_state_or_province" api:"required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStreetAddress string `json:"receiver_street_address" api:"required"`
	// The name of the receiver of the transfer. This is not verified by Increase.
	ReceivingCompanyOrIndividualName string `json:"receiving_company_or_individual_name" api:"required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the receiving bank country.
	ReceivingDepositoryFinancialInstitutionCountry string `json:"receiving_depository_financial_institution_country" api:"required"`
	// An identifier for the receiving bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	ReceivingDepositoryFinancialInstitutionID string `json:"receiving_depository_financial_institution_id" api:"required"`
	// An instruction of how to interpret the
	// `receiving_depository_financial_institution_id` field for this Transaction.
	ReceivingDepositoryFinancialInstitutionIDQualifier InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier" api:"required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string                                     `json:"receiving_depository_financial_institution_name" api:"required"`
	JSON                                        inboundACHTransferInternationalAddendaJSON `json:"-"`
}

// inboundACHTransferInternationalAddendaJSON contains the JSON metadata for the
// struct [InboundACHTransferInternationalAddenda]
type inboundACHTransferInternationalAddendaJSON struct {
	DestinationCountryCode                                 apijson.Field
	DestinationCurrencyCode                                apijson.Field
	ForeignExchangeIndicator                               apijson.Field
	ForeignExchangeReference                               apijson.Field
	ForeignExchangeReferenceIndicator                      apijson.Field
	ForeignPaymentAmount                                   apijson.Field
	ForeignTraceNumber                                     apijson.Field
	InternationalTransactionTypeCode                       apijson.Field
	OriginatingCurrencyCode                                apijson.Field
	OriginatingDepositoryFinancialInstitutionBranchCountry apijson.Field
	OriginatingDepositoryFinancialInstitutionID            apijson.Field
	OriginatingDepositoryFinancialInstitutionIDQualifier   apijson.Field
	OriginatingDepositoryFinancialInstitutionName          apijson.Field
	OriginatorCity                                         apijson.Field
	OriginatorCountry                                      apijson.Field
	OriginatorIdentification                               apijson.Field
	OriginatorName                                         apijson.Field
	OriginatorPostalCode                                   apijson.Field
	OriginatorStateOrProvince                              apijson.Field
	OriginatorStreetAddress                                apijson.Field
	PaymentRelatedInformation                              apijson.Field
	PaymentRelatedInformation2                             apijson.Field
	ReceiverCity                                           apijson.Field
	ReceiverCountry                                        apijson.Field
	ReceiverIdentificationNumber                           apijson.Field
	ReceiverPostalCode                                     apijson.Field
	ReceiverStateOrProvince                                apijson.Field
	ReceiverStreetAddress                                  apijson.Field
	ReceivingCompanyOrIndividualName                       apijson.Field
	ReceivingDepositoryFinancialInstitutionCountry         apijson.Field
	ReceivingDepositoryFinancialInstitutionID              apijson.Field
	ReceivingDepositoryFinancialInstitutionIDQualifier     apijson.Field
	ReceivingDepositoryFinancialInstitutionName            apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *InboundACHTransferInternationalAddenda) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferInternationalAddendaJSON) RawJSON() string {
	return r.raw
}

// A description of how the foreign exchange rate was calculated.
type InboundACHTransferInternationalAddendaForeignExchangeIndicator string

const (
	InboundACHTransferInternationalAddendaForeignExchangeIndicatorFixedToVariable InboundACHTransferInternationalAddendaForeignExchangeIndicator = "fixed_to_variable"
	InboundACHTransferInternationalAddendaForeignExchangeIndicatorVariableToFixed InboundACHTransferInternationalAddendaForeignExchangeIndicator = "variable_to_fixed"
	InboundACHTransferInternationalAddendaForeignExchangeIndicatorFixedToFixed    InboundACHTransferInternationalAddendaForeignExchangeIndicator = "fixed_to_fixed"
)

func (r InboundACHTransferInternationalAddendaForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case InboundACHTransferInternationalAddendaForeignExchangeIndicatorFixedToVariable, InboundACHTransferInternationalAddendaForeignExchangeIndicatorVariableToFixed, InboundACHTransferInternationalAddendaForeignExchangeIndicatorFixedToFixed:
		return true
	}
	return false
}

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicator string

const (
	InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicatorForeignExchangeRate            InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicatorBlank                          InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicator = "blank"
)

func (r InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicator) IsKnown() bool {
	switch r {
	case InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicatorForeignExchangeRate, InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber, InboundACHTransferInternationalAddendaForeignExchangeReferenceIndicatorBlank:
		return true
	}
	return false
}

// The type of transfer. Set by the originator.
type InboundACHTransferInternationalAddendaInternationalTransactionTypeCode string

const (
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeAnnuity                  InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "annuity"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeBusinessOrCommercial     InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "business_or_commercial"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeDeposit                  InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "deposit"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeLoan                     InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "loan"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeMiscellaneous            InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "miscellaneous"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeMortgage                 InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "mortgage"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodePension                  InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "pension"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeRemittance               InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "remittance"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeRentOrLease              InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "rent_or_lease"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeSalaryOrPayroll          InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "salary_or_payroll"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeTax                      InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "tax"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeAccountsReceivable       InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "accounts_receivable"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeBackOfficeConversion     InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "back_office_conversion"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeMachineTransfer          InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "machine_transfer"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodePointOfPurchase          InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "point_of_purchase"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodePointOfSale              InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "point_of_sale"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeRepresentedCheck         InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "represented_check"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeSharedNetworkTransaction InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "shared_network_transaction"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeTelphoneInitiated        InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "telphone_initiated"
	InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeInternetInitiated        InboundACHTransferInternationalAddendaInternationalTransactionTypeCode = "internet_initiated"
)

func (r InboundACHTransferInternationalAddendaInternationalTransactionTypeCode) IsKnown() bool {
	switch r {
	case InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeAnnuity, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeBusinessOrCommercial, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeDeposit, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeLoan, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeMiscellaneous, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeMortgage, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodePension, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeRemittance, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeRentOrLease, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeSalaryOrPayroll, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeTax, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeAccountsReceivable, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeBackOfficeConversion, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeMachineTransfer, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodePointOfPurchase, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodePointOfSale, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeRepresentedCheck, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeSharedNetworkTransaction, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeTelphoneInitiated, InboundACHTransferInternationalAddendaInternationalTransactionTypeCodeInternetInitiated:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifierBicCode                      InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifierIban                         InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifierBicCode, InboundACHTransferInternationalAddendaOriginatingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifierBicCode                      InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifierIban                         InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifierBicCode, InboundACHTransferInternationalAddendaReceivingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// If you initiate a notification of change in response to the transfer, this will
// contain its details.
type InboundACHTransferNotificationOfChange struct {
	// The new account number provided in the notification of change.
	UpdatedAccountNumber string `json:"updated_account_number" api:"required,nullable"`
	// The new routing number provided in the notification of change.
	UpdatedRoutingNumber string                                     `json:"updated_routing_number" api:"required,nullable"`
	JSON                 inboundACHTransferNotificationOfChangeJSON `json:"-"`
}

// inboundACHTransferNotificationOfChangeJSON contains the JSON metadata for the
// struct [InboundACHTransferNotificationOfChange]
type inboundACHTransferNotificationOfChangeJSON struct {
	UpdatedAccountNumber apijson.Field
	UpdatedRoutingNumber apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InboundACHTransferNotificationOfChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferNotificationOfChangeJSON) RawJSON() string {
	return r.raw
}

// A subhash containing information about when and how the transfer settled at the
// Federal Reserve.
type InboundACHTransferSettlement struct {
	// When the funds for this transfer settle at the recipient bank at the Federal
	// Reserve.
	SettledAt time.Time `json:"settled_at" api:"required" format:"date-time"`
	// The settlement schedule this transfer follows.
	SettlementSchedule InboundACHTransferSettlementSettlementSchedule `json:"settlement_schedule" api:"required"`
	JSON               inboundACHTransferSettlementJSON               `json:"-"`
}

// inboundACHTransferSettlementJSON contains the JSON metadata for the struct
// [InboundACHTransferSettlement]
type inboundACHTransferSettlementJSON struct {
	SettledAt          apijson.Field
	SettlementSchedule apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InboundACHTransferSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferSettlementJSON) RawJSON() string {
	return r.raw
}

// The settlement schedule this transfer follows.
type InboundACHTransferSettlementSettlementSchedule string

const (
	InboundACHTransferSettlementSettlementScheduleSameDay     InboundACHTransferSettlementSettlementSchedule = "same_day"
	InboundACHTransferSettlementSettlementScheduleFutureDated InboundACHTransferSettlementSettlementSchedule = "future_dated"
)

func (r InboundACHTransferSettlementSettlementSchedule) IsKnown() bool {
	switch r {
	case InboundACHTransferSettlementSettlementScheduleSameDay, InboundACHTransferSettlementSettlementScheduleFutureDated:
		return true
	}
	return false
}

// The Standard Entry Class (SEC) code of the transfer.
type InboundACHTransferStandardEntryClassCode string

const (
	InboundACHTransferStandardEntryClassCodeCorporateCreditOrDebit        InboundACHTransferStandardEntryClassCode = "corporate_credit_or_debit"
	InboundACHTransferStandardEntryClassCodeCorporateTradeExchange        InboundACHTransferStandardEntryClassCode = "corporate_trade_exchange"
	InboundACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit InboundACHTransferStandardEntryClassCode = "prearranged_payments_and_deposit"
	InboundACHTransferStandardEntryClassCodeInternetInitiated             InboundACHTransferStandardEntryClassCode = "internet_initiated"
	InboundACHTransferStandardEntryClassCodePointOfSale                   InboundACHTransferStandardEntryClassCode = "point_of_sale"
	InboundACHTransferStandardEntryClassCodeTelephoneInitiated            InboundACHTransferStandardEntryClassCode = "telephone_initiated"
	InboundACHTransferStandardEntryClassCodeCustomerInitiated             InboundACHTransferStandardEntryClassCode = "customer_initiated"
	InboundACHTransferStandardEntryClassCodeAccountsReceivable            InboundACHTransferStandardEntryClassCode = "accounts_receivable"
	InboundACHTransferStandardEntryClassCodeMachineTransfer               InboundACHTransferStandardEntryClassCode = "machine_transfer"
	InboundACHTransferStandardEntryClassCodeSharedNetworkTransaction      InboundACHTransferStandardEntryClassCode = "shared_network_transaction"
	InboundACHTransferStandardEntryClassCodeRepresentedCheck              InboundACHTransferStandardEntryClassCode = "represented_check"
	InboundACHTransferStandardEntryClassCodeBackOfficeConversion          InboundACHTransferStandardEntryClassCode = "back_office_conversion"
	InboundACHTransferStandardEntryClassCodePointOfPurchase               InboundACHTransferStandardEntryClassCode = "point_of_purchase"
	InboundACHTransferStandardEntryClassCodeCheckTruncation               InboundACHTransferStandardEntryClassCode = "check_truncation"
	InboundACHTransferStandardEntryClassCodeDestroyedCheck                InboundACHTransferStandardEntryClassCode = "destroyed_check"
	InboundACHTransferStandardEntryClassCodeInternationalACHTransaction   InboundACHTransferStandardEntryClassCode = "international_ach_transaction"
)

func (r InboundACHTransferStandardEntryClassCode) IsKnown() bool {
	switch r {
	case InboundACHTransferStandardEntryClassCodeCorporateCreditOrDebit, InboundACHTransferStandardEntryClassCodeCorporateTradeExchange, InboundACHTransferStandardEntryClassCodePrearrangedPaymentsAndDeposit, InboundACHTransferStandardEntryClassCodeInternetInitiated, InboundACHTransferStandardEntryClassCodePointOfSale, InboundACHTransferStandardEntryClassCodeTelephoneInitiated, InboundACHTransferStandardEntryClassCodeCustomerInitiated, InboundACHTransferStandardEntryClassCodeAccountsReceivable, InboundACHTransferStandardEntryClassCodeMachineTransfer, InboundACHTransferStandardEntryClassCodeSharedNetworkTransaction, InboundACHTransferStandardEntryClassCodeRepresentedCheck, InboundACHTransferStandardEntryClassCodeBackOfficeConversion, InboundACHTransferStandardEntryClassCodePointOfPurchase, InboundACHTransferStandardEntryClassCodeCheckTruncation, InboundACHTransferStandardEntryClassCodeDestroyedCheck, InboundACHTransferStandardEntryClassCodeInternationalACHTransaction:
		return true
	}
	return false
}

// The status of the transfer.
type InboundACHTransferStatus string

const (
	InboundACHTransferStatusPending  InboundACHTransferStatus = "pending"
	InboundACHTransferStatusDeclined InboundACHTransferStatus = "declined"
	InboundACHTransferStatusAccepted InboundACHTransferStatus = "accepted"
	InboundACHTransferStatusReturned InboundACHTransferStatus = "returned"
)

func (r InboundACHTransferStatus) IsKnown() bool {
	switch r {
	case InboundACHTransferStatusPending, InboundACHTransferStatusDeclined, InboundACHTransferStatusAccepted, InboundACHTransferStatusReturned:
		return true
	}
	return false
}

// If your transfer is returned, this will contain details of the return.
type InboundACHTransferTransferReturn struct {
	// The reason for the transfer return.
	Reason InboundACHTransferTransferReturnReason `json:"reason" api:"required"`
	// The time at which the transfer was returned.
	ReturnedAt time.Time `json:"returned_at" api:"required" format:"date-time"`
	// The id of the transaction for the returned transfer.
	TransactionID string                               `json:"transaction_id" api:"required"`
	JSON          inboundACHTransferTransferReturnJSON `json:"-"`
}

// inboundACHTransferTransferReturnJSON contains the JSON metadata for the struct
// [InboundACHTransferTransferReturn]
type inboundACHTransferTransferReturnJSON struct {
	Reason        apijson.Field
	ReturnedAt    apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InboundACHTransferTransferReturn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inboundACHTransferTransferReturnJSON) RawJSON() string {
	return r.raw
}

// The reason for the transfer return.
type InboundACHTransferTransferReturnReason string

const (
	InboundACHTransferTransferReturnReasonInsufficientFunds                                           InboundACHTransferTransferReturnReason = "insufficient_funds"
	InboundACHTransferTransferReturnReasonReturnedPerOdfiRequest                                      InboundACHTransferTransferReturnReason = "returned_per_odfi_request"
	InboundACHTransferTransferReturnReasonAuthorizationRevokedByCustomer                              InboundACHTransferTransferReturnReason = "authorization_revoked_by_customer"
	InboundACHTransferTransferReturnReasonPaymentStopped                                              InboundACHTransferTransferReturnReason = "payment_stopped"
	InboundACHTransferTransferReturnReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   InboundACHTransferTransferReturnReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundACHTransferTransferReturnReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferTransferReturnReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	InboundACHTransferTransferReturnReasonBeneficiaryOrAccountHolderDeceased                          InboundACHTransferTransferReturnReason = "beneficiary_or_account_holder_deceased"
	InboundACHTransferTransferReturnReasonCreditEntryRefusedByReceiver                                InboundACHTransferTransferReturnReason = "credit_entry_refused_by_receiver"
	InboundACHTransferTransferReturnReasonDuplicateEntry                                              InboundACHTransferTransferReturnReason = "duplicate_entry"
	InboundACHTransferTransferReturnReasonCorporateCustomerAdvisedNotAuthorized                       InboundACHTransferTransferReturnReason = "corporate_customer_advised_not_authorized"
)

func (r InboundACHTransferTransferReturnReason) IsKnown() bool {
	switch r {
	case InboundACHTransferTransferReturnReasonInsufficientFunds, InboundACHTransferTransferReturnReasonReturnedPerOdfiRequest, InboundACHTransferTransferReturnReasonAuthorizationRevokedByCustomer, InboundACHTransferTransferReturnReasonPaymentStopped, InboundACHTransferTransferReturnReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, InboundACHTransferTransferReturnReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, InboundACHTransferTransferReturnReasonBeneficiaryOrAccountHolderDeceased, InboundACHTransferTransferReturnReasonCreditEntryRefusedByReceiver, InboundACHTransferTransferReturnReasonDuplicateEntry, InboundACHTransferTransferReturnReasonCorporateCustomerAdvisedNotAuthorized:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_ach_transfer`.
type InboundACHTransferType string

const (
	InboundACHTransferTypeInboundACHTransfer InboundACHTransferType = "inbound_ach_transfer"
)

func (r InboundACHTransferType) IsKnown() bool {
	switch r {
	case InboundACHTransferTypeInboundACHTransfer:
		return true
	}
	return false
}

type InboundACHTransferListParams struct {
	// Filter Inbound ACH Transfers to ones belonging to the specified Account.
	AccountID param.Field[string] `query:"account_id"`
	// Filter Inbound ACH Transfers to ones belonging to the specified Account Number.
	AccountNumberID param.Field[string]                                `query:"account_number_id"`
	CreatedAt       param.Field[InboundACHTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                              `query:"limit"`
	Status param.Field[InboundACHTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [InboundACHTransferListParams]'s query parameters as
// `url.Values`.
func (r InboundACHTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundACHTransferListParamsCreatedAt struct {
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

// URLQuery serializes [InboundACHTransferListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r InboundACHTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundACHTransferListParamsStatus struct {
	// Filter Inbound ACH Transfers to those with the specified status. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]InboundACHTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [InboundACHTransferListParamsStatus]'s query parameters as
// `url.Values`.
func (r InboundACHTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InboundACHTransferListParamsStatusIn string

const (
	InboundACHTransferListParamsStatusInPending  InboundACHTransferListParamsStatusIn = "pending"
	InboundACHTransferListParamsStatusInDeclined InboundACHTransferListParamsStatusIn = "declined"
	InboundACHTransferListParamsStatusInAccepted InboundACHTransferListParamsStatusIn = "accepted"
	InboundACHTransferListParamsStatusInReturned InboundACHTransferListParamsStatusIn = "returned"
)

func (r InboundACHTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case InboundACHTransferListParamsStatusInPending, InboundACHTransferListParamsStatusInDeclined, InboundACHTransferListParamsStatusInAccepted, InboundACHTransferListParamsStatusInReturned:
		return true
	}
	return false
}

type InboundACHTransferNewNotificationOfChangeParams struct {
	// The updated account number to send in the notification of change.
	UpdatedAccountNumber param.Field[string] `json:"updated_account_number"`
	// The updated routing number to send in the notification of change.
	UpdatedRoutingNumber param.Field[string] `json:"updated_routing_number"`
}

func (r InboundACHTransferNewNotificationOfChangeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InboundACHTransferDeclineParams struct {
	// The reason why this transfer will be returned. If this parameter is unset, the
	// return codes will be `payment_stopped` for debits and
	// `credit_entry_refused_by_receiver` for credits.
	Reason param.Field[InboundACHTransferDeclineParamsReason] `json:"reason"`
}

func (r InboundACHTransferDeclineParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why this transfer will be returned. If this parameter is unset, the
// return codes will be `payment_stopped` for debits and
// `credit_entry_refused_by_receiver` for credits.
type InboundACHTransferDeclineParamsReason string

const (
	InboundACHTransferDeclineParamsReasonInsufficientFunds                                           InboundACHTransferDeclineParamsReason = "insufficient_funds"
	InboundACHTransferDeclineParamsReasonReturnedPerOdfiRequest                                      InboundACHTransferDeclineParamsReason = "returned_per_odfi_request"
	InboundACHTransferDeclineParamsReasonAuthorizationRevokedByCustomer                              InboundACHTransferDeclineParamsReason = "authorization_revoked_by_customer"
	InboundACHTransferDeclineParamsReasonPaymentStopped                                              InboundACHTransferDeclineParamsReason = "payment_stopped"
	InboundACHTransferDeclineParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   InboundACHTransferDeclineParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundACHTransferDeclineParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferDeclineParamsReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	InboundACHTransferDeclineParamsReasonBeneficiaryOrAccountHolderDeceased                          InboundACHTransferDeclineParamsReason = "beneficiary_or_account_holder_deceased"
	InboundACHTransferDeclineParamsReasonCreditEntryRefusedByReceiver                                InboundACHTransferDeclineParamsReason = "credit_entry_refused_by_receiver"
	InboundACHTransferDeclineParamsReasonDuplicateEntry                                              InboundACHTransferDeclineParamsReason = "duplicate_entry"
	InboundACHTransferDeclineParamsReasonCorporateCustomerAdvisedNotAuthorized                       InboundACHTransferDeclineParamsReason = "corporate_customer_advised_not_authorized"
)

func (r InboundACHTransferDeclineParamsReason) IsKnown() bool {
	switch r {
	case InboundACHTransferDeclineParamsReasonInsufficientFunds, InboundACHTransferDeclineParamsReasonReturnedPerOdfiRequest, InboundACHTransferDeclineParamsReasonAuthorizationRevokedByCustomer, InboundACHTransferDeclineParamsReasonPaymentStopped, InboundACHTransferDeclineParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, InboundACHTransferDeclineParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, InboundACHTransferDeclineParamsReasonBeneficiaryOrAccountHolderDeceased, InboundACHTransferDeclineParamsReasonCreditEntryRefusedByReceiver, InboundACHTransferDeclineParamsReasonDuplicateEntry, InboundACHTransferDeclineParamsReasonCorporateCustomerAdvisedNotAuthorized:
		return true
	}
	return false
}

type InboundACHTransferTransferReturnParams struct {
	// The reason why this transfer will be returned. The most usual return codes are
	// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
	Reason param.Field[InboundACHTransferTransferReturnParamsReason] `json:"reason" api:"required"`
}

func (r InboundACHTransferTransferReturnParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason why this transfer will be returned. The most usual return codes are
// `payment_stopped` for debits and `credit_entry_refused_by_receiver` for credits.
type InboundACHTransferTransferReturnParamsReason string

const (
	InboundACHTransferTransferReturnParamsReasonInsufficientFunds                                           InboundACHTransferTransferReturnParamsReason = "insufficient_funds"
	InboundACHTransferTransferReturnParamsReasonAuthorizationRevokedByCustomer                              InboundACHTransferTransferReturnParamsReason = "authorization_revoked_by_customer"
	InboundACHTransferTransferReturnParamsReasonPaymentStopped                                              InboundACHTransferTransferReturnParamsReason = "payment_stopped"
	InboundACHTransferTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete   InboundACHTransferTransferReturnParamsReason = "customer_advised_unauthorized_improper_ineligible_or_incomplete"
	InboundACHTransferTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity InboundACHTransferTransferReturnParamsReason = "representative_payee_deceased_or_unable_to_continue_in_that_capacity"
	InboundACHTransferTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased                          InboundACHTransferTransferReturnParamsReason = "beneficiary_or_account_holder_deceased"
	InboundACHTransferTransferReturnParamsReasonCreditEntryRefusedByReceiver                                InboundACHTransferTransferReturnParamsReason = "credit_entry_refused_by_receiver"
	InboundACHTransferTransferReturnParamsReasonDuplicateEntry                                              InboundACHTransferTransferReturnParamsReason = "duplicate_entry"
	InboundACHTransferTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized                       InboundACHTransferTransferReturnParamsReason = "corporate_customer_advised_not_authorized"
)

func (r InboundACHTransferTransferReturnParamsReason) IsKnown() bool {
	switch r {
	case InboundACHTransferTransferReturnParamsReasonInsufficientFunds, InboundACHTransferTransferReturnParamsReasonAuthorizationRevokedByCustomer, InboundACHTransferTransferReturnParamsReasonPaymentStopped, InboundACHTransferTransferReturnParamsReasonCustomerAdvisedUnauthorizedImproperIneligibleOrIncomplete, InboundACHTransferTransferReturnParamsReasonRepresentativePayeeDeceasedOrUnableToContinueInThatCapacity, InboundACHTransferTransferReturnParamsReasonBeneficiaryOrAccountHolderDeceased, InboundACHTransferTransferReturnParamsReasonCreditEntryRefusedByReceiver, InboundACHTransferTransferReturnParamsReasonDuplicateEntry, InboundACHTransferTransferReturnParamsReasonCorporateCustomerAdvisedNotAuthorized:
		return true
	}
	return false
}
