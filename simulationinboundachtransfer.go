// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// SimulationInboundACHTransferService contains methods and other services that
// help with interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimulationInboundACHTransferService] method instead.
type SimulationInboundACHTransferService struct {
	Options []option.RequestOption
}

// NewSimulationInboundACHTransferService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationInboundACHTransferService(opts ...option.RequestOption) (r *SimulationInboundACHTransferService) {
	r = &SimulationInboundACHTransferService{}
	r.Options = opts
	return
}

// Simulates an inbound ACH transfer to your account. This imitates initiating a
// transfer to an Increase account from a different financial institution. The
// transfer may be either a credit or a debit depending on if the `amount` is
// positive or negative. The result of calling this API will contain the created
// transfer. You can pass a `resolve_at` parameter to allow for a window to
// [action on the Inbound ACH Transfer](https://increase.com/documentation/receiving-ach-transfers).
// Alternatively, if you don't pass the `resolve_at` parameter the result will
// contain either a [Transaction](#transactions) or a
// [Declined Transaction](#declined-transactions) depending on whether or not the
// transfer is allowed.
func (r *SimulationInboundACHTransferService) New(ctx context.Context, body SimulationInboundACHTransferNewParams, opts ...option.RequestOption) (res *InboundACHTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "simulations/inbound_ach_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SimulationInboundACHTransferNewParams struct {
	// The identifier of the Account Number the inbound ACH Transfer is for.
	AccountNumberID param.Field[string] `json:"account_number_id,required"`
	// The transfer amount in cents. A positive amount originates a credit transfer
	// pushing funds to the receiving account. A negative amount originates a debit
	// transfer pulling funds from the receiving account.
	Amount param.Field[int64] `json:"amount,required"`
	// The description of the date of the transfer.
	CompanyDescriptiveDate param.Field[string] `json:"company_descriptive_date"`
	// Data associated with the transfer set by the sender.
	CompanyDiscretionaryData param.Field[string] `json:"company_discretionary_data"`
	// The description of the transfer set by the sender.
	CompanyEntryDescription param.Field[string] `json:"company_entry_description"`
	// The sender's company ID.
	CompanyID param.Field[string] `json:"company_id"`
	// The name of the sender.
	CompanyName param.Field[string] `json:"company_name"`
	// The ID of the receiver of the transfer.
	ReceiverIDNumber param.Field[string] `json:"receiver_id_number"`
	// The name of the receiver of the transfer.
	ReceiverName param.Field[string] `json:"receiver_name"`
	// The time at which the transfer should be resolved. If not provided will resolve
	// immediately.
	ResolveAt param.Field[time.Time] `json:"resolve_at" format:"date-time"`
	// The standard entry class code for the transfer.
	StandardEntryClassCode param.Field[SimulationInboundACHTransferNewParamsStandardEntryClassCode] `json:"standard_entry_class_code"`
}

func (r SimulationInboundACHTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The standard entry class code for the transfer.
type SimulationInboundACHTransferNewParamsStandardEntryClassCode string

const (
	// Corporate Credit and Debit (CCD).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit SimulationInboundACHTransferNewParamsStandardEntryClassCode = "corporate_credit_or_debit"
	// Corporate Trade Exchange (CTX).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeCorporateTradeExchange SimulationInboundACHTransferNewParamsStandardEntryClassCode = "corporate_trade_exchange"
	// Prearranged Payments and Deposits (PPD).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit SimulationInboundACHTransferNewParamsStandardEntryClassCode = "prearranged_payments_and_deposit"
	// Internet Initiated (WEB).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeInternetInitiated SimulationInboundACHTransferNewParamsStandardEntryClassCode = "internet_initiated"
	// Point of Sale (POS).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodePointOfSale SimulationInboundACHTransferNewParamsStandardEntryClassCode = "point_of_sale"
	// Telephone Initiated (TEL).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeTelephoneInitiated SimulationInboundACHTransferNewParamsStandardEntryClassCode = "telephone_initiated"
	// Customer Initiated (CIE).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeCustomerInitiated SimulationInboundACHTransferNewParamsStandardEntryClassCode = "customer_initiated"
	// Accounts Receivable (ARC).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeAccountsReceivable SimulationInboundACHTransferNewParamsStandardEntryClassCode = "accounts_receivable"
	// Machine Transfer (MTE).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeMachineTransfer SimulationInboundACHTransferNewParamsStandardEntryClassCode = "machine_transfer"
	// Shared Network Transaction (SHR).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeSharedNetworkTransaction SimulationInboundACHTransferNewParamsStandardEntryClassCode = "shared_network_transaction"
	// Represented Check (RCK).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeRepresentedCheck SimulationInboundACHTransferNewParamsStandardEntryClassCode = "represented_check"
	// Back Office Conversion (BOC).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeBackOfficeConversion SimulationInboundACHTransferNewParamsStandardEntryClassCode = "back_office_conversion"
	// Point of Purchase (POP).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodePointOfPurchase SimulationInboundACHTransferNewParamsStandardEntryClassCode = "point_of_purchase"
	// Check Truncation (TRC).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeCheckTruncation SimulationInboundACHTransferNewParamsStandardEntryClassCode = "check_truncation"
	// Destroyed Check (XCK).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeDestroyedCheck SimulationInboundACHTransferNewParamsStandardEntryClassCode = "destroyed_check"
	// International ACH Transaction (IAT).
	SimulationInboundACHTransferNewParamsStandardEntryClassCodeInternationalACHTransaction SimulationInboundACHTransferNewParamsStandardEntryClassCode = "international_ach_transaction"
)

func (r SimulationInboundACHTransferNewParamsStandardEntryClassCode) IsKnown() bool {
	switch r {
	case SimulationInboundACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit, SimulationInboundACHTransferNewParamsStandardEntryClassCodeCorporateTradeExchange, SimulationInboundACHTransferNewParamsStandardEntryClassCodePrearrangedPaymentsAndDeposit, SimulationInboundACHTransferNewParamsStandardEntryClassCodeInternetInitiated, SimulationInboundACHTransferNewParamsStandardEntryClassCodePointOfSale, SimulationInboundACHTransferNewParamsStandardEntryClassCodeTelephoneInitiated, SimulationInboundACHTransferNewParamsStandardEntryClassCodeCustomerInitiated, SimulationInboundACHTransferNewParamsStandardEntryClassCodeAccountsReceivable, SimulationInboundACHTransferNewParamsStandardEntryClassCodeMachineTransfer, SimulationInboundACHTransferNewParamsStandardEntryClassCodeSharedNetworkTransaction, SimulationInboundACHTransferNewParamsStandardEntryClassCodeRepresentedCheck, SimulationInboundACHTransferNewParamsStandardEntryClassCodeBackOfficeConversion, SimulationInboundACHTransferNewParamsStandardEntryClassCodePointOfPurchase, SimulationInboundACHTransferNewParamsStandardEntryClassCodeCheckTruncation, SimulationInboundACHTransferNewParamsStandardEntryClassCodeDestroyedCheck, SimulationInboundACHTransferNewParamsStandardEntryClassCodeInternationalACHTransaction:
		return true
	}
	return false
}
