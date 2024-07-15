// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationInboundInternationalACHTransferService contains methods and other
// services that help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundInternationalACHTransferService] method instead.
type SimulationInboundInternationalACHTransferService struct {
	Options []option.RequestOption
}

// NewSimulationInboundInternationalACHTransferService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewSimulationInboundInternationalACHTransferService(opts ...option.RequestOption) (r *SimulationInboundInternationalACHTransferService) {
	r = &SimulationInboundInternationalACHTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound international ACH transfer to your account. This imitates
// initiating a transfer to an Increase account from a different financial
// institution. The transfer may be either a credit or a debit depending on if the
// `amount` is positive or negative. The result of calling this API will contain
// the created transfer. .
func (r *SimulationInboundInternationalACHTransferService) New(ctx context.Context, body SimulationInboundInternationalACHTransferNewParams, opts ...option.RequestOption) (res *SimulationInboundInternationalACHTransferNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_international_ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// An Inbound International ACH Transfer is created when an "IAT" ACH transfer is
// initiated at another bank and received by Increase. There are additional fields
// on this object that are not present on all Inbound ACH Transfer object.
type SimulationInboundInternationalACHTransferNewResponse struct {
	// The amount in the minor unit of the destination account currency. For dollars,
	// for example, this is cents.
	Amount int64 `json:"amount,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the destination country.
	DestinationCountryCode string `json:"destination_country_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// destination bank account.
	DestinationCurrencyCode string `json:"destination_currency_code,required"`
	// A description of how the foreign exchange rate was calculated.
	ForeignExchangeIndicator SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicator `json:"foreign_exchange_indicator,required"`
	// Depending on the `foreign_exchange_reference_indicator`, an exchange rate or a
	// reference to a well-known rate.
	ForeignExchangeReference string `json:"foreign_exchange_reference,required,nullable"`
	// An instruction of how to interpret the `foreign_exchange_reference` field for
	// this Transaction.
	ForeignExchangeReferenceIndicator SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicator `json:"foreign_exchange_reference_indicator,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount int64 `json:"foreign_payment_amount,required"`
	// A reference number in the foreign banking infrastructure.
	ForeignTraceNumber string `json:"foreign_trace_number,required,nullable"`
	// The type of transfer. Set by the originator.
	InternationalTransactionTypeCode SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode `json:"international_transaction_type_code,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// originating bank account.
	OriginatingCurrencyCode string `json:"originating_currency_code,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the originating branch country.
	OriginatingDepositoryFinancialInstitutionBranchCountry string `json:"originating_depository_financial_institution_branch_country,required"`
	// An identifier for the originating bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	OriginatingDepositoryFinancialInstitutionID string `json:"originating_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `originating_depository_financial_institution_id` field for this Transaction.
	OriginatingDepositoryFinancialInstitutionIDQualifier SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifier `json:"originating_depository_financial_institution_id_qualifier,required"`
	// The name of the originating bank. Sometimes this will refer to an American bank
	// and obscure the correspondent foreign bank.
	OriginatingDepositoryFinancialInstitutionName string `json:"originating_depository_financial_institution_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorCity string `json:"originator_city,required"`
	// A description field set by the originator.
	OriginatorCompanyEntryDescription string `json:"originator_company_entry_description,required"`
	// A portion of the originator address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the originator country.
	OriginatorCountry string `json:"originator_country,required"`
	// An identifier for the originating company. This is generally stable across
	// multiple ACH transfers.
	OriginatorIdentification string `json:"originator_identification,required"`
	// Either the name of the originator or an intermediary money transmitter.
	OriginatorName string `json:"originator_name,required"`
	// A portion of the originator address. This may be incomplete.
	OriginatorPostalCode string `json:"originator_postal_code,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStateOrProvince string `json:"originator_state_or_province,required,nullable"`
	// A portion of the originator address. This may be incomplete.
	OriginatorStreetAddress string `json:"originator_street_address,required"`
	// A description field set by the originator.
	PaymentRelatedInformation string `json:"payment_related_information,required,nullable"`
	// A description field set by the originator.
	PaymentRelatedInformation2 string `json:"payment_related_information2,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverCity string `json:"receiver_city,required"`
	// A portion of the receiver address. The
	// [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2 country
	// code of the receiver country.
	ReceiverCountry string `json:"receiver_country,required"`
	// An identification number the originator uses for the receiver.
	ReceiverIdentificationNumber string `json:"receiver_identification_number,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverPostalCode string `json:"receiver_postal_code,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStateOrProvince string `json:"receiver_state_or_province,required,nullable"`
	// A portion of the receiver address. This may be incomplete.
	ReceiverStreetAddress string `json:"receiver_street_address,required"`
	// The name of the receiver of the transfer. This is not verified by Increase.
	ReceivingCompanyOrIndividualName string `json:"receiving_company_or_individual_name,required"`
	// The [ISO 3166](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), Alpha-2
	// country code of the receiving bank country.
	ReceivingDepositoryFinancialInstitutionCountry string `json:"receiving_depository_financial_institution_country,required"`
	// An identifier for the receiving bank. One of an International Bank Account
	// Number (IBAN) bank identifier, SWIFT Bank Identification Code (BIC), or a
	// domestic identifier like a US Routing Number.
	ReceivingDepositoryFinancialInstitutionID string `json:"receiving_depository_financial_institution_id,required"`
	// An instruction of how to interpret the
	// `receiving_depository_financial_institution_id` field for this Transaction.
	ReceivingDepositoryFinancialInstitutionIDQualifier SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifier `json:"receiving_depository_financial_institution_id_qualifier,required"`
	// The name of the receiving bank, as set by the sending financial institution.
	ReceivingDepositoryFinancialInstitutionName string `json:"receiving_depository_financial_institution_name,required"`
	// A 15 digit number recorded in the Nacha file and available to both the
	// originating and receiving bank. Along with the amount, date, and originating
	// routing number, this can be used to identify the ACH transfer at either bank.
	// ACH trace numbers are not unique, but are
	// [used to correlate returns](https://increase.com/documentation/ach-returns#ach-returns).
	TraceNumber string `json:"trace_number,required"`
	// A constant representing the object's type. For this resource it will always be
	// `inbound_international_ach_transfer`.
	Type SimulationInboundInternationalACHTransferNewResponseType `json:"type,required"`
	JSON simulationInboundInternationalACHTransferNewResponseJSON `json:"-"`
}

// simulationInboundInternationalACHTransferNewResponseJSON contains the JSON
// metadata for the struct [SimulationInboundInternationalACHTransferNewResponse]
type simulationInboundInternationalACHTransferNewResponseJSON struct {
	Amount                                                 apijson.Field
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
	OriginatorCompanyEntryDescription                      apijson.Field
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
	TraceNumber                                            apijson.Field
	Type                                                   apijson.Field
	raw                                                    string
	ExtraFields                                            map[string]apijson.Field
}

func (r *SimulationInboundInternationalACHTransferNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationInboundInternationalACHTransferNewResponseJSON) RawJSON() string {
	return r.raw
}

// A description of how the foreign exchange rate was calculated.
type SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicator string

const (
	// The originator chose an amount in their own currency. The settled amount in USD
	// was converted using the exchange rate.
	SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicatorFixedToVariable SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicator = "fixed_to_variable"
	// The originator chose an amount to settle in USD. The originator's amount was
	// variable; known only after the foreign exchange conversion.
	SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicatorVariableToFixed SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicator = "variable_to_fixed"
	// The amount was originated and settled as a fixed amount in USD. There is no
	// foreign exchange conversion.
	SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicatorFixedToFixed SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicator = "fixed_to_fixed"
)

func (r SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicatorFixedToVariable, SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicatorVariableToFixed, SimulationInboundInternationalACHTransferNewResponseForeignExchangeIndicatorFixedToFixed:
		return true
	}
	return false
}

// An instruction of how to interpret the `foreign_exchange_reference` field for
// this Transaction.
type SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicator string

const (
	// The ACH file contains a foreign exchange rate.
	SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicatorForeignExchangeRate SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicator = "foreign_exchange_rate"
	// The ACH file contains a reference to a well-known foreign exchange rate.
	SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicator = "foreign_exchange_reference_number"
	// There is no foreign exchange for this transfer, so the
	// `foreign_exchange_reference` field is blank.
	SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicatorBlank SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicator = "blank"
)

func (r SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicator) IsKnown() bool {
	switch r {
	case SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicatorForeignExchangeRate, SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicatorForeignExchangeReferenceNumber, SimulationInboundInternationalACHTransferNewResponseForeignExchangeReferenceIndicatorBlank:
		return true
	}
	return false
}

// The type of transfer. Set by the originator.
type SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode string

const (
	// Sent as `ANN` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeAnnuity SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "annuity"
	// Sent as `BUS` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeBusinessOrCommercial SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "business_or_commercial"
	// Sent as `DEP` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeDeposit SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "deposit"
	// Sent as `LOA` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeLoan SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "loan"
	// Sent as `MIS` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeMiscellaneous SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "miscellaneous"
	// Sent as `MOR` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeMortgage SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "mortgage"
	// Sent as `PEN` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodePension SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "pension"
	// Sent as `REM` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeRemittance SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "remittance"
	// Sent as `RLS` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeRentOrLease SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "rent_or_lease"
	// Sent as `SAL` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeSalaryOrPayroll SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "salary_or_payroll"
	// Sent as `TAX` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeTax SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "tax"
	// Sent as `ARC` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeAccountsReceivable SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "accounts_receivable"
	// Sent as `BOC` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeBackOfficeConversion SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "back_office_conversion"
	// Sent as `MTE` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeMachineTransfer SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "machine_transfer"
	// Sent as `POP` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodePointOfPurchase SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "point_of_purchase"
	// Sent as `POS` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodePointOfSale SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "point_of_sale"
	// Sent as `RCK` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeRepresentedCheck SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "represented_check"
	// Sent as `SHR` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeSharedNetworkTransaction SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "shared_network_transaction"
	// Sent as `TEL` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeTelphoneInitiated SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "telphone_initiated"
	// Sent as `WEB` in the Nacha file.
	SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeInternetInitiated SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode = "internet_initiated"
)

func (r SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCode) IsKnown() bool {
	switch r {
	case SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeAnnuity, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeBusinessOrCommercial, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeDeposit, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeLoan, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeMiscellaneous, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeMortgage, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodePension, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeRemittance, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeRentOrLease, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeSalaryOrPayroll, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeTax, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeAccountsReceivable, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeBackOfficeConversion, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeMachineTransfer, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodePointOfPurchase, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodePointOfSale, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeRepresentedCheck, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeSharedNetworkTransaction, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeTelphoneInitiated, SimulationInboundInternationalACHTransferNewResponseInternationalTransactionTypeCodeInternetInitiated:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `originating_depository_financial_institution_id` field for this Transaction.
type SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifierBicCode SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifierIban SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifierBicCode, SimulationInboundInternationalACHTransferNewResponseOriginatingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// An instruction of how to interpret the
// `receiving_depository_financial_institution_id` field for this Transaction.
type SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifier string

const (
	// A domestic clearing system number. In the US, for example, this is the American
	// Banking Association (ABA) routing number.
	SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifier = "national_clearing_system_number"
	// The SWIFT Bank Identifier Code (BIC) of the bank.
	SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifierBicCode SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifier = "bic_code"
	// An International Bank Account Number.
	SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifierIban SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifier = "iban"
)

func (r SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifier) IsKnown() bool {
	switch r {
	case SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifierNationalClearingSystemNumber, SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifierBicCode, SimulationInboundInternationalACHTransferNewResponseReceivingDepositoryFinancialInstitutionIDQualifierIban:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `inbound_international_ach_transfer`.
type SimulationInboundInternationalACHTransferNewResponseType string

const (
	SimulationInboundInternationalACHTransferNewResponseTypeInboundInternationalACHTransfer SimulationInboundInternationalACHTransferNewResponseType = "inbound_international_ach_transfer"
)

func (r SimulationInboundInternationalACHTransferNewResponseType) IsKnown() bool {
	switch r {
	case SimulationInboundInternationalACHTransferNewResponseTypeInboundInternationalACHTransfer:
		return true
	}
	return false
}

type SimulationInboundInternationalACHTransferNewParams struct {
	// The identifier of the Account Number the inbound international ACH Transfer is
	// for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount param.Field[int64] `json:"amount,required"`
	// The amount in the minor unit of the foreign payment currency. For dollars, for
	// example, this is cents.
	ForeignPaymentAmount param.Field[int64] `json:"foreign_payment_amount,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code for the
	// originating bank account.
	OriginatingCurrencyCode param.Field[string] `json:"originating_currency_code,required"`
	// A description field set by the originator.
	OriginatorCompanyEntryDescription param.Field[string] `json:"originator_company_entry_description"`
	// Either the name of the originator or an intermediary money transmitter.
	OriginatorName param.Field[string] `json:"originator_name"`
	// An identification number the originator uses for the receiver.
	ReceiverIdentificationNumber param.Field[string] `json:"receiver_identification_number"`
	// The name of the receiver of the transfer.
	ReceivingCompanyOrIndividualName param.Field[string] `json:"receiving_company_or_individual_name"`
}

func (r SimulationInboundInternationalACHTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}