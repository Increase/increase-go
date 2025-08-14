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

// CardPushTransferService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardPushTransferService] method instead.
type CardPushTransferService struct {
	Options []option.RequestOption
}

// NewCardPushTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCardPushTransferService(opts ...option.RequestOption) (r *CardPushTransferService) {
	r = &CardPushTransferService{}
	r.Options = opts
	return
}

// Create a Card Push Transfer
func (r *CardPushTransferService) New(ctx context.Context, body CardPushTransferNewParams, opts ...option.RequestOption) (res *CardPushTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := "card_push_transfers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Card Push Transfer
func (r *CardPushTransferService) Get(ctx context.Context, cardPushTransferID string, opts ...option.RequestOption) (res *CardPushTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if cardPushTransferID == "" {
		err = errors.New("missing required card_push_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("card_push_transfers/%s", cardPushTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Card Push Transfers
func (r *CardPushTransferService) List(ctx context.Context, query CardPushTransferListParams, opts ...option.RequestOption) (res *pagination.Page[CardPushTransfer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "card_push_transfers"
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

// List Card Push Transfers
func (r *CardPushTransferService) ListAutoPaging(ctx context.Context, query CardPushTransferListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CardPushTransfer] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Approves a Card Push Transfer in a pending_approval state.
func (r *CardPushTransferService) Approve(ctx context.Context, cardPushTransferID string, opts ...option.RequestOption) (res *CardPushTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if cardPushTransferID == "" {
		err = errors.New("missing required card_push_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("card_push_transfers/%s/approve", cardPushTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancels a Card Push Transfer in a pending_approval state.
func (r *CardPushTransferService) Cancel(ctx context.Context, cardPushTransferID string, opts ...option.RequestOption) (res *CardPushTransfer, err error) {
	opts = append(r.Options[:], opts...)
	if cardPushTransferID == "" {
		err = errors.New("missing required card_push_transfer_id parameter")
		return
	}
	path := fmt.Sprintf("card_push_transfers/%s/cancel", cardPushTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Card Push Transfers send funds to a recipient's payment card in real-time.
type CardPushTransfer struct {
	// The Card Push Transfer's identifier.
	ID string `json:"id,required"`
	// If the transfer is accepted by the recipient bank, this will contain
	// supplemental details.
	Acceptance CardPushTransferAcceptance `json:"acceptance,required,nullable"`
	// The Account from which the transfer was sent.
	AccountID string `json:"account_id,required"`
	// The transfer amount in USD cents.
	Amount int64 `json:"amount,required"`
	// If your account requires approvals for transfers and the transfer was approved,
	// this will contain details of the approval.
	Approval CardPushTransferApproval `json:"approval,required,nullable"`
	// The Business Application Identifier describes the type of transaction being
	// performed. Your program must be approved for the specified Business Application
	// Identifier in order to use it.
	BusinessApplicationIdentifier CardPushTransferBusinessApplicationIdentifier `json:"business_application_identifier,required"`
	// If your account requires approvals for transfers and the transfer was not
	// approved, this will contain details of the cancellation.
	Cancellation CardPushTransferCancellation `json:"cancellation,required,nullable"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// What object created the transfer, either via the API or the dashboard.
	CreatedBy CardPushTransferCreatedBy `json:"created_by,required,nullable"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
	// currency.
	Currency CardPushTransferCurrency `json:"currency,required"`
	// If the transfer is rejected by the card network or the destination financial
	// institution, this will contain supplemental details.
	Decline CardPushTransferDecline `json:"decline,required,nullable"`
	// The idempotency key you chose for this object. This value is unique across
	// Increase and is used to ensure that a request is only processed once. Learn more
	// about [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey string `json:"idempotency_key,required,nullable"`
	// The merchant category code (MCC) of the merchant (generally your business)
	// sending the transfer. This is a four-digit code that describes the type of
	// business or service provided by the merchant. Your program must be approved for
	// the specified MCC in order to use it.
	MerchantCategoryCode string `json:"merchant_category_code,required"`
	// The city name of the merchant (generally your business) sending the transfer.
	MerchantCityName string `json:"merchant_city_name,required"`
	// The merchant name shows up as the statement descriptor for the transfer. This is
	// typically the name of your business or organization.
	MerchantName string `json:"merchant_name,required"`
	// For certain Business Application Identifiers, the statement descriptor is
	// `merchant_name_prefix*sender_name`, where the `merchant_name_prefix` is a one to
	// four character prefix that identifies the merchant.
	MerchantNamePrefix string `json:"merchant_name_prefix,required"`
	// The postal code of the merchant (generally your business) sending the transfer.
	MerchantPostalCode string `json:"merchant_postal_code,required"`
	// The state of the merchant (generally your business) sending the transfer.
	MerchantState string `json:"merchant_state,required"`
	// The name of the funds recipient.
	RecipientName string `json:"recipient_name,required"`
	// The city of the sender.
	SenderAddressCity string `json:"sender_address_city,required"`
	// The address line 1 of the sender.
	SenderAddressLine1 string `json:"sender_address_line1,required"`
	// The postal code of the sender.
	SenderAddressPostalCode string `json:"sender_address_postal_code,required"`
	// The state of the sender.
	SenderAddressState string `json:"sender_address_state,required"`
	// The name of the funds originator.
	SenderName string `json:"sender_name,required"`
	// The Account Number the recipient will see as having sent the transfer.
	SourceAccountNumberID string `json:"source_account_number_id,required"`
	// The lifecycle status of the transfer.
	Status CardPushTransferStatus `json:"status,required"`
	// After the transfer is submitted to the card network, this will contain
	// supplemental details.
	Submission CardPushTransferSubmission `json:"submission,required,nullable"`
	// A constant representing the object's type. For this resource it will always be
	// `card_push_transfer`.
	Type CardPushTransferType `json:"type,required"`
	JSON cardPushTransferJSON `json:"-"`
}

// cardPushTransferJSON contains the JSON metadata for the struct
// [CardPushTransfer]
type cardPushTransferJSON struct {
	ID                            apijson.Field
	Acceptance                    apijson.Field
	AccountID                     apijson.Field
	Amount                        apijson.Field
	Approval                      apijson.Field
	BusinessApplicationIdentifier apijson.Field
	Cancellation                  apijson.Field
	CreatedAt                     apijson.Field
	CreatedBy                     apijson.Field
	Currency                      apijson.Field
	Decline                       apijson.Field
	IdempotencyKey                apijson.Field
	MerchantCategoryCode          apijson.Field
	MerchantCityName              apijson.Field
	MerchantName                  apijson.Field
	MerchantNamePrefix            apijson.Field
	MerchantPostalCode            apijson.Field
	MerchantState                 apijson.Field
	RecipientName                 apijson.Field
	SenderAddressCity             apijson.Field
	SenderAddressLine1            apijson.Field
	SenderAddressPostalCode       apijson.Field
	SenderAddressState            apijson.Field
	SenderName                    apijson.Field
	SourceAccountNumberID         apijson.Field
	Status                        apijson.Field
	Submission                    apijson.Field
	Type                          apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *CardPushTransfer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferJSON) RawJSON() string {
	return r.raw
}

// If the transfer is accepted by the recipient bank, this will contain
// supplemental details.
type CardPushTransferAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was accepted by the issuing bank.
	AcceptedAt time.Time `json:"accepted_at,required" format:"date-time"`
	// The authorization identification response from the issuing bank.
	AuthorizationIdentificationResponse string `json:"authorization_identification_response,required"`
	// The result of the Card Verification Value 2 match.
	CardVerificationValue2Result CardPushTransferAcceptanceCardVerificationValue2Result `json:"card_verification_value2_result,required,nullable"`
	// A unique identifier for the transaction on the card network.
	NetworkTransactionIdentifier string                         `json:"network_transaction_identifier,required,nullable"`
	JSON                         cardPushTransferAcceptanceJSON `json:"-"`
}

// cardPushTransferAcceptanceJSON contains the JSON metadata for the struct
// [CardPushTransferAcceptance]
type cardPushTransferAcceptanceJSON struct {
	AcceptedAt                          apijson.Field
	AuthorizationIdentificationResponse apijson.Field
	CardVerificationValue2Result        apijson.Field
	NetworkTransactionIdentifier        apijson.Field
	raw                                 string
	ExtraFields                         map[string]apijson.Field
}

func (r *CardPushTransferAcceptance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferAcceptanceJSON) RawJSON() string {
	return r.raw
}

// The result of the Card Verification Value 2 match.
type CardPushTransferAcceptanceCardVerificationValue2Result string

const (
	CardPushTransferAcceptanceCardVerificationValue2ResultMatch   CardPushTransferAcceptanceCardVerificationValue2Result = "match"
	CardPushTransferAcceptanceCardVerificationValue2ResultNoMatch CardPushTransferAcceptanceCardVerificationValue2Result = "no_match"
)

func (r CardPushTransferAcceptanceCardVerificationValue2Result) IsKnown() bool {
	switch r {
	case CardPushTransferAcceptanceCardVerificationValue2ResultMatch, CardPushTransferAcceptanceCardVerificationValue2ResultNoMatch:
		return true
	}
	return false
}

// If your account requires approvals for transfers and the transfer was approved,
// this will contain details of the approval.
type CardPushTransferApproval struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was approved.
	ApprovedAt time.Time `json:"approved_at,required" format:"date-time"`
	// If the Transfer was approved by a user in the dashboard, the email address of
	// that user.
	ApprovedBy string                       `json:"approved_by,required,nullable"`
	JSON       cardPushTransferApprovalJSON `json:"-"`
}

// cardPushTransferApprovalJSON contains the JSON metadata for the struct
// [CardPushTransferApproval]
type cardPushTransferApprovalJSON struct {
	ApprovedAt  apijson.Field
	ApprovedBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPushTransferApproval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferApprovalJSON) RawJSON() string {
	return r.raw
}

// The Business Application Identifier describes the type of transaction being
// performed. Your program must be approved for the specified Business Application
// Identifier in order to use it.
type CardPushTransferBusinessApplicationIdentifier string

const (
	CardPushTransferBusinessApplicationIdentifierAccountToAccount           CardPushTransferBusinessApplicationIdentifier = "account_to_account"
	CardPushTransferBusinessApplicationIdentifierBusinessToBusiness         CardPushTransferBusinessApplicationIdentifier = "business_to_business"
	CardPushTransferBusinessApplicationIdentifierMoneyTransferBankInitiated CardPushTransferBusinessApplicationIdentifier = "money_transfer_bank_initiated"
	CardPushTransferBusinessApplicationIdentifierNonCardBillPayment         CardPushTransferBusinessApplicationIdentifier = "non_card_bill_payment"
	CardPushTransferBusinessApplicationIdentifierConsumerBillPayment        CardPushTransferBusinessApplicationIdentifier = "consumer_bill_payment"
	CardPushTransferBusinessApplicationIdentifierCardBillPayment            CardPushTransferBusinessApplicationIdentifier = "card_bill_payment"
	CardPushTransferBusinessApplicationIdentifierFundsDisbursement          CardPushTransferBusinessApplicationIdentifier = "funds_disbursement"
	CardPushTransferBusinessApplicationIdentifierFundsTransfer              CardPushTransferBusinessApplicationIdentifier = "funds_transfer"
	CardPushTransferBusinessApplicationIdentifierLoyaltyAndOffers           CardPushTransferBusinessApplicationIdentifier = "loyalty_and_offers"
	CardPushTransferBusinessApplicationIdentifierMerchantDisbursement       CardPushTransferBusinessApplicationIdentifier = "merchant_disbursement"
	CardPushTransferBusinessApplicationIdentifierMerchantPayment            CardPushTransferBusinessApplicationIdentifier = "merchant_payment"
	CardPushTransferBusinessApplicationIdentifierPersonToPerson             CardPushTransferBusinessApplicationIdentifier = "person_to_person"
	CardPushTransferBusinessApplicationIdentifierTopUp                      CardPushTransferBusinessApplicationIdentifier = "top_up"
	CardPushTransferBusinessApplicationIdentifierWalletTransfer             CardPushTransferBusinessApplicationIdentifier = "wallet_transfer"
)

func (r CardPushTransferBusinessApplicationIdentifier) IsKnown() bool {
	switch r {
	case CardPushTransferBusinessApplicationIdentifierAccountToAccount, CardPushTransferBusinessApplicationIdentifierBusinessToBusiness, CardPushTransferBusinessApplicationIdentifierMoneyTransferBankInitiated, CardPushTransferBusinessApplicationIdentifierNonCardBillPayment, CardPushTransferBusinessApplicationIdentifierConsumerBillPayment, CardPushTransferBusinessApplicationIdentifierCardBillPayment, CardPushTransferBusinessApplicationIdentifierFundsDisbursement, CardPushTransferBusinessApplicationIdentifierFundsTransfer, CardPushTransferBusinessApplicationIdentifierLoyaltyAndOffers, CardPushTransferBusinessApplicationIdentifierMerchantDisbursement, CardPushTransferBusinessApplicationIdentifierMerchantPayment, CardPushTransferBusinessApplicationIdentifierPersonToPerson, CardPushTransferBusinessApplicationIdentifierTopUp, CardPushTransferBusinessApplicationIdentifierWalletTransfer:
		return true
	}
	return false
}

// If your account requires approvals for transfers and the transfer was not
// approved, this will contain details of the cancellation.
type CardPushTransferCancellation struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Transfer was canceled.
	CanceledAt time.Time `json:"canceled_at,required" format:"date-time"`
	// If the Transfer was canceled by a user in the dashboard, the email address of
	// that user.
	CanceledBy string                           `json:"canceled_by,required,nullable"`
	JSON       cardPushTransferCancellationJSON `json:"-"`
}

// cardPushTransferCancellationJSON contains the JSON metadata for the struct
// [CardPushTransferCancellation]
type cardPushTransferCancellationJSON struct {
	CanceledAt  apijson.Field
	CanceledBy  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPushTransferCancellation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferCancellationJSON) RawJSON() string {
	return r.raw
}

// What object created the transfer, either via the API or the dashboard.
type CardPushTransferCreatedBy struct {
	// If present, details about the API key that created the transfer.
	APIKey CardPushTransferCreatedByAPIKey `json:"api_key,required,nullable"`
	// The type of object that created this transfer.
	Category CardPushTransferCreatedByCategory `json:"category,required"`
	// If present, details about the OAuth Application that created the transfer.
	OAuthApplication CardPushTransferCreatedByOAuthApplication `json:"oauth_application,required,nullable"`
	// If present, details about the User that created the transfer.
	User CardPushTransferCreatedByUser `json:"user,required,nullable"`
	JSON cardPushTransferCreatedByJSON `json:"-"`
}

// cardPushTransferCreatedByJSON contains the JSON metadata for the struct
// [CardPushTransferCreatedBy]
type cardPushTransferCreatedByJSON struct {
	APIKey           apijson.Field
	Category         apijson.Field
	OAuthApplication apijson.Field
	User             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CardPushTransferCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferCreatedByJSON) RawJSON() string {
	return r.raw
}

// If present, details about the API key that created the transfer.
type CardPushTransferCreatedByAPIKey struct {
	// The description set for the API key when it was created.
	Description string                              `json:"description,required,nullable"`
	JSON        cardPushTransferCreatedByAPIKeyJSON `json:"-"`
}

// cardPushTransferCreatedByAPIKeyJSON contains the JSON metadata for the struct
// [CardPushTransferCreatedByAPIKey]
type cardPushTransferCreatedByAPIKeyJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPushTransferCreatedByAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferCreatedByAPIKeyJSON) RawJSON() string {
	return r.raw
}

// The type of object that created this transfer.
type CardPushTransferCreatedByCategory string

const (
	CardPushTransferCreatedByCategoryAPIKey           CardPushTransferCreatedByCategory = "api_key"
	CardPushTransferCreatedByCategoryOAuthApplication CardPushTransferCreatedByCategory = "oauth_application"
	CardPushTransferCreatedByCategoryUser             CardPushTransferCreatedByCategory = "user"
)

func (r CardPushTransferCreatedByCategory) IsKnown() bool {
	switch r {
	case CardPushTransferCreatedByCategoryAPIKey, CardPushTransferCreatedByCategoryOAuthApplication, CardPushTransferCreatedByCategoryUser:
		return true
	}
	return false
}

// If present, details about the OAuth Application that created the transfer.
type CardPushTransferCreatedByOAuthApplication struct {
	// The name of the OAuth Application.
	Name string                                        `json:"name,required"`
	JSON cardPushTransferCreatedByOAuthApplicationJSON `json:"-"`
}

// cardPushTransferCreatedByOAuthApplicationJSON contains the JSON metadata for the
// struct [CardPushTransferCreatedByOAuthApplication]
type cardPushTransferCreatedByOAuthApplicationJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPushTransferCreatedByOAuthApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferCreatedByOAuthApplicationJSON) RawJSON() string {
	return r.raw
}

// If present, details about the User that created the transfer.
type CardPushTransferCreatedByUser struct {
	// The email address of the User.
	Email string                            `json:"email,required"`
	JSON  cardPushTransferCreatedByUserJSON `json:"-"`
}

// cardPushTransferCreatedByUserJSON contains the JSON metadata for the struct
// [CardPushTransferCreatedByUser]
type cardPushTransferCreatedByUserJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CardPushTransferCreatedByUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferCreatedByUserJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the transfer's
// currency.
type CardPushTransferCurrency string

const (
	CardPushTransferCurrencyCad CardPushTransferCurrency = "CAD"
	CardPushTransferCurrencyChf CardPushTransferCurrency = "CHF"
	CardPushTransferCurrencyEur CardPushTransferCurrency = "EUR"
	CardPushTransferCurrencyGbp CardPushTransferCurrency = "GBP"
	CardPushTransferCurrencyJpy CardPushTransferCurrency = "JPY"
	CardPushTransferCurrencyUsd CardPushTransferCurrency = "USD"
)

func (r CardPushTransferCurrency) IsKnown() bool {
	switch r {
	case CardPushTransferCurrencyCad, CardPushTransferCurrencyChf, CardPushTransferCurrencyEur, CardPushTransferCurrencyGbp, CardPushTransferCurrencyJpy, CardPushTransferCurrencyUsd:
		return true
	}
	return false
}

// If the transfer is rejected by the card network or the destination financial
// institution, this will contain supplemental details.
type CardPushTransferDecline struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer declined.
	DeclinedAt time.Time `json:"declined_at,required" format:"date-time"`
	// A unique identifier for the transaction on the card network.
	NetworkTransactionIdentifier string `json:"network_transaction_identifier,required,nullable"`
	// The reason why the transfer was declined.
	Reason CardPushTransferDeclineReason `json:"reason,required"`
	JSON   cardPushTransferDeclineJSON   `json:"-"`
}

// cardPushTransferDeclineJSON contains the JSON metadata for the struct
// [CardPushTransferDecline]
type cardPushTransferDeclineJSON struct {
	DeclinedAt                   apijson.Field
	NetworkTransactionIdentifier apijson.Field
	Reason                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CardPushTransferDecline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferDeclineJSON) RawJSON() string {
	return r.raw
}

// The reason why the transfer was declined.
type CardPushTransferDeclineReason string

const (
	CardPushTransferDeclineReasonDoNotHonor                                              CardPushTransferDeclineReason = "do_not_honor"
	CardPushTransferDeclineReasonActivityCountLimitExceeded                              CardPushTransferDeclineReason = "activity_count_limit_exceeded"
	CardPushTransferDeclineReasonReferToCardIssuer                                       CardPushTransferDeclineReason = "refer_to_card_issuer"
	CardPushTransferDeclineReasonReferToCardIssuerSpecialCondition                       CardPushTransferDeclineReason = "refer_to_card_issuer_special_condition"
	CardPushTransferDeclineReasonInvalidMerchant                                         CardPushTransferDeclineReason = "invalid_merchant"
	CardPushTransferDeclineReasonPickUpCard                                              CardPushTransferDeclineReason = "pick_up_card"
	CardPushTransferDeclineReasonError                                                   CardPushTransferDeclineReason = "error"
	CardPushTransferDeclineReasonPickUpCardSpecial                                       CardPushTransferDeclineReason = "pick_up_card_special"
	CardPushTransferDeclineReasonInvalidTransaction                                      CardPushTransferDeclineReason = "invalid_transaction"
	CardPushTransferDeclineReasonInvalidAmount                                           CardPushTransferDeclineReason = "invalid_amount"
	CardPushTransferDeclineReasonInvalidAccountNumber                                    CardPushTransferDeclineReason = "invalid_account_number"
	CardPushTransferDeclineReasonNoSuchIssuer                                            CardPushTransferDeclineReason = "no_such_issuer"
	CardPushTransferDeclineReasonReEnterTransaction                                      CardPushTransferDeclineReason = "re_enter_transaction"
	CardPushTransferDeclineReasonNoCreditAccount                                         CardPushTransferDeclineReason = "no_credit_account"
	CardPushTransferDeclineReasonPickUpCardLost                                          CardPushTransferDeclineReason = "pick_up_card_lost"
	CardPushTransferDeclineReasonPickUpCardStolen                                        CardPushTransferDeclineReason = "pick_up_card_stolen"
	CardPushTransferDeclineReasonClosedAccount                                           CardPushTransferDeclineReason = "closed_account"
	CardPushTransferDeclineReasonInsufficientFunds                                       CardPushTransferDeclineReason = "insufficient_funds"
	CardPushTransferDeclineReasonNoCheckingAccount                                       CardPushTransferDeclineReason = "no_checking_account"
	CardPushTransferDeclineReasonNoSavingsAccount                                        CardPushTransferDeclineReason = "no_savings_account"
	CardPushTransferDeclineReasonExpiredCard                                             CardPushTransferDeclineReason = "expired_card"
	CardPushTransferDeclineReasonTransactionNotPermittedToCardholder                     CardPushTransferDeclineReason = "transaction_not_permitted_to_cardholder"
	CardPushTransferDeclineReasonTransactionNotAllowedAtTerminal                         CardPushTransferDeclineReason = "transaction_not_allowed_at_terminal"
	CardPushTransferDeclineReasonSuspectedFraud                                          CardPushTransferDeclineReason = "suspected_fraud"
	CardPushTransferDeclineReasonActivityAmountLimitExceeded                             CardPushTransferDeclineReason = "activity_amount_limit_exceeded"
	CardPushTransferDeclineReasonRestrictedCard                                          CardPushTransferDeclineReason = "restricted_card"
	CardPushTransferDeclineReasonSecurityViolation                                       CardPushTransferDeclineReason = "security_violation"
	CardPushTransferDeclineReasonTransactionDoesNotFulfillAntiMoneyLaunderingRequirement CardPushTransferDeclineReason = "transaction_does_not_fulfill_anti_money_laundering_requirement"
	CardPushTransferDeclineReasonBlockedFirstUse                                         CardPushTransferDeclineReason = "blocked_first_use"
	CardPushTransferDeclineReasonCreditIssuerUnavailable                                 CardPushTransferDeclineReason = "credit_issuer_unavailable"
	CardPushTransferDeclineReasonNegativeCardVerificationValueResults                    CardPushTransferDeclineReason = "negative_card_verification_value_results"
	CardPushTransferDeclineReasonIssuerUnavailable                                       CardPushTransferDeclineReason = "issuer_unavailable"
	CardPushTransferDeclineReasonFinancialInstitutionCannotBeFound                       CardPushTransferDeclineReason = "financial_institution_cannot_be_found"
	CardPushTransferDeclineReasonTransactionCannotBeCompleted                            CardPushTransferDeclineReason = "transaction_cannot_be_completed"
	CardPushTransferDeclineReasonDuplicateTransaction                                    CardPushTransferDeclineReason = "duplicate_transaction"
	CardPushTransferDeclineReasonSystemMalfunction                                       CardPushTransferDeclineReason = "system_malfunction"
	CardPushTransferDeclineReasonAdditionalCustomerAuthenticationRequired                CardPushTransferDeclineReason = "additional_customer_authentication_required"
	CardPushTransferDeclineReasonSurchargeAmountNotPermitted                             CardPushTransferDeclineReason = "surcharge_amount_not_permitted"
	CardPushTransferDeclineReasonDeclineForCvv2Failure                                   CardPushTransferDeclineReason = "decline_for_cvv2_failure"
	CardPushTransferDeclineReasonStopPaymentOrder                                        CardPushTransferDeclineReason = "stop_payment_order"
	CardPushTransferDeclineReasonRevocationOfAuthorizationOrder                          CardPushTransferDeclineReason = "revocation_of_authorization_order"
	CardPushTransferDeclineReasonRevocationOfAllAuthorizationsOrder                      CardPushTransferDeclineReason = "revocation_of_all_authorizations_order"
)

func (r CardPushTransferDeclineReason) IsKnown() bool {
	switch r {
	case CardPushTransferDeclineReasonDoNotHonor, CardPushTransferDeclineReasonActivityCountLimitExceeded, CardPushTransferDeclineReasonReferToCardIssuer, CardPushTransferDeclineReasonReferToCardIssuerSpecialCondition, CardPushTransferDeclineReasonInvalidMerchant, CardPushTransferDeclineReasonPickUpCard, CardPushTransferDeclineReasonError, CardPushTransferDeclineReasonPickUpCardSpecial, CardPushTransferDeclineReasonInvalidTransaction, CardPushTransferDeclineReasonInvalidAmount, CardPushTransferDeclineReasonInvalidAccountNumber, CardPushTransferDeclineReasonNoSuchIssuer, CardPushTransferDeclineReasonReEnterTransaction, CardPushTransferDeclineReasonNoCreditAccount, CardPushTransferDeclineReasonPickUpCardLost, CardPushTransferDeclineReasonPickUpCardStolen, CardPushTransferDeclineReasonClosedAccount, CardPushTransferDeclineReasonInsufficientFunds, CardPushTransferDeclineReasonNoCheckingAccount, CardPushTransferDeclineReasonNoSavingsAccount, CardPushTransferDeclineReasonExpiredCard, CardPushTransferDeclineReasonTransactionNotPermittedToCardholder, CardPushTransferDeclineReasonTransactionNotAllowedAtTerminal, CardPushTransferDeclineReasonSuspectedFraud, CardPushTransferDeclineReasonActivityAmountLimitExceeded, CardPushTransferDeclineReasonRestrictedCard, CardPushTransferDeclineReasonSecurityViolation, CardPushTransferDeclineReasonTransactionDoesNotFulfillAntiMoneyLaunderingRequirement, CardPushTransferDeclineReasonBlockedFirstUse, CardPushTransferDeclineReasonCreditIssuerUnavailable, CardPushTransferDeclineReasonNegativeCardVerificationValueResults, CardPushTransferDeclineReasonIssuerUnavailable, CardPushTransferDeclineReasonFinancialInstitutionCannotBeFound, CardPushTransferDeclineReasonTransactionCannotBeCompleted, CardPushTransferDeclineReasonDuplicateTransaction, CardPushTransferDeclineReasonSystemMalfunction, CardPushTransferDeclineReasonAdditionalCustomerAuthenticationRequired, CardPushTransferDeclineReasonSurchargeAmountNotPermitted, CardPushTransferDeclineReasonDeclineForCvv2Failure, CardPushTransferDeclineReasonStopPaymentOrder, CardPushTransferDeclineReasonRevocationOfAuthorizationOrder, CardPushTransferDeclineReasonRevocationOfAllAuthorizationsOrder:
		return true
	}
	return false
}

// The lifecycle status of the transfer.
type CardPushTransferStatus string

const (
	CardPushTransferStatusPendingApproval   CardPushTransferStatus = "pending_approval"
	CardPushTransferStatusCanceled          CardPushTransferStatus = "canceled"
	CardPushTransferStatusPendingReviewing  CardPushTransferStatus = "pending_reviewing"
	CardPushTransferStatusRequiresAttention CardPushTransferStatus = "requires_attention"
	CardPushTransferStatusPendingSubmission CardPushTransferStatus = "pending_submission"
	CardPushTransferStatusSubmitted         CardPushTransferStatus = "submitted"
	CardPushTransferStatusComplete          CardPushTransferStatus = "complete"
	CardPushTransferStatusDeclined          CardPushTransferStatus = "declined"
)

func (r CardPushTransferStatus) IsKnown() bool {
	switch r {
	case CardPushTransferStatusPendingApproval, CardPushTransferStatusCanceled, CardPushTransferStatusPendingReviewing, CardPushTransferStatusRequiresAttention, CardPushTransferStatusPendingSubmission, CardPushTransferStatusSubmitted, CardPushTransferStatusComplete, CardPushTransferStatusDeclined:
		return true
	}
	return false
}

// After the transfer is submitted to the card network, this will contain
// supplemental details.
type CardPushTransferSubmission struct {
	// A 12-digit retrieval reference number that identifies the transfer. Usually a
	// combination of a timestamp and the trace number.
	RetrievalReferenceNumber string `json:"retrieval_reference_number,required"`
	// A unique reference for the transfer.
	SenderReference string `json:"sender_reference,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the transfer was submitted to card network.
	SubmittedAt time.Time `json:"submitted_at,required" format:"date-time"`
	// A 6-digit trace number that identifies the transfer within a small window of
	// time.
	TraceNumber string                         `json:"trace_number,required"`
	JSON        cardPushTransferSubmissionJSON `json:"-"`
}

// cardPushTransferSubmissionJSON contains the JSON metadata for the struct
// [CardPushTransferSubmission]
type cardPushTransferSubmissionJSON struct {
	RetrievalReferenceNumber apijson.Field
	SenderReference          apijson.Field
	SubmittedAt              apijson.Field
	TraceNumber              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CardPushTransferSubmission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cardPushTransferSubmissionJSON) RawJSON() string {
	return r.raw
}

// A constant representing the object's type. For this resource it will always be
// `card_push_transfer`.
type CardPushTransferType string

const (
	CardPushTransferTypeCardPushTransfer CardPushTransferType = "card_push_transfer"
)

func (r CardPushTransferType) IsKnown() bool {
	switch r {
	case CardPushTransferTypeCardPushTransfer:
		return true
	}
	return false
}

type CardPushTransferNewParams struct {
	// The transfer amount in USD cents. For Card Push transfers, must be positive.
	Amount param.Field[int64] `json:"amount,required"`
	// The Business Application Identifier describes the type of transaction being
	// performed. Your program must be approved for the specified Business Application
	// Identifier in order to use it.
	BusinessApplicationIdentifier param.Field[CardPushTransferNewParamsBusinessApplicationIdentifier] `json:"business_application_identifier,required"`
	// The Increase identifier for the Card Token that represents the card number
	// you're pushing funds to.
	CardTokenID param.Field[string] `json:"card_token_id,required"`
	// The merchant category code (MCC) of the merchant (generally your business)
	// sending the transfer. This is a four-digit code that describes the type of
	// business or service provided by the merchant. Your program must be approved for
	// the specified MCC in order to use it.
	MerchantCategoryCode param.Field[string] `json:"merchant_category_code,required"`
	// The city name of the merchant (generally your business) sending the transfer.
	MerchantCityName param.Field[string] `json:"merchant_city_name,required"`
	// The merchant name shows up as the statement descriptor for the transfer. This is
	// typically the name of your business or organization.
	MerchantName param.Field[string] `json:"merchant_name,required"`
	// For certain Business Application Identifiers, the statement descriptor is
	// `merchant_name_prefix*sender_name`, where the `merchant_name_prefix` is a one to
	// four character prefix that identifies the merchant.
	MerchantNamePrefix param.Field[string] `json:"merchant_name_prefix,required"`
	// The postal code of the merchant (generally your business) sending the transfer.
	MerchantPostalCode param.Field[string] `json:"merchant_postal_code,required"`
	// The state of the merchant (generally your business) sending the transfer.
	MerchantState param.Field[string] `json:"merchant_state,required"`
	// The name of the funds recipient.
	RecipientName param.Field[string] `json:"recipient_name,required"`
	// The city of the sender.
	SenderAddressCity param.Field[string] `json:"sender_address_city,required"`
	// The address line 1 of the sender.
	SenderAddressLine1 param.Field[string] `json:"sender_address_line1,required"`
	// The postal code of the sender.
	SenderAddressPostalCode param.Field[string] `json:"sender_address_postal_code,required"`
	// The state of the sender.
	SenderAddressState param.Field[string] `json:"sender_address_state,required"`
	// The name of the funds originator.
	SenderName param.Field[string] `json:"sender_name,required"`
	// The identifier of the Account Number from which to send the transfer.
	SourceAccountNumberID param.Field[string] `json:"source_account_number_id,required"`
	// Whether the transfer requires explicit approval via the dashboard or API.
	RequireApproval param.Field[bool] `json:"require_approval"`
}

func (r CardPushTransferNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Business Application Identifier describes the type of transaction being
// performed. Your program must be approved for the specified Business Application
// Identifier in order to use it.
type CardPushTransferNewParamsBusinessApplicationIdentifier string

const (
	CardPushTransferNewParamsBusinessApplicationIdentifierAccountToAccount           CardPushTransferNewParamsBusinessApplicationIdentifier = "account_to_account"
	CardPushTransferNewParamsBusinessApplicationIdentifierBusinessToBusiness         CardPushTransferNewParamsBusinessApplicationIdentifier = "business_to_business"
	CardPushTransferNewParamsBusinessApplicationIdentifierMoneyTransferBankInitiated CardPushTransferNewParamsBusinessApplicationIdentifier = "money_transfer_bank_initiated"
	CardPushTransferNewParamsBusinessApplicationIdentifierNonCardBillPayment         CardPushTransferNewParamsBusinessApplicationIdentifier = "non_card_bill_payment"
	CardPushTransferNewParamsBusinessApplicationIdentifierConsumerBillPayment        CardPushTransferNewParamsBusinessApplicationIdentifier = "consumer_bill_payment"
	CardPushTransferNewParamsBusinessApplicationIdentifierCardBillPayment            CardPushTransferNewParamsBusinessApplicationIdentifier = "card_bill_payment"
	CardPushTransferNewParamsBusinessApplicationIdentifierFundsDisbursement          CardPushTransferNewParamsBusinessApplicationIdentifier = "funds_disbursement"
	CardPushTransferNewParamsBusinessApplicationIdentifierFundsTransfer              CardPushTransferNewParamsBusinessApplicationIdentifier = "funds_transfer"
	CardPushTransferNewParamsBusinessApplicationIdentifierLoyaltyAndOffers           CardPushTransferNewParamsBusinessApplicationIdentifier = "loyalty_and_offers"
	CardPushTransferNewParamsBusinessApplicationIdentifierMerchantDisbursement       CardPushTransferNewParamsBusinessApplicationIdentifier = "merchant_disbursement"
	CardPushTransferNewParamsBusinessApplicationIdentifierMerchantPayment            CardPushTransferNewParamsBusinessApplicationIdentifier = "merchant_payment"
	CardPushTransferNewParamsBusinessApplicationIdentifierPersonToPerson             CardPushTransferNewParamsBusinessApplicationIdentifier = "person_to_person"
	CardPushTransferNewParamsBusinessApplicationIdentifierTopUp                      CardPushTransferNewParamsBusinessApplicationIdentifier = "top_up"
	CardPushTransferNewParamsBusinessApplicationIdentifierWalletTransfer             CardPushTransferNewParamsBusinessApplicationIdentifier = "wallet_transfer"
)

func (r CardPushTransferNewParamsBusinessApplicationIdentifier) IsKnown() bool {
	switch r {
	case CardPushTransferNewParamsBusinessApplicationIdentifierAccountToAccount, CardPushTransferNewParamsBusinessApplicationIdentifierBusinessToBusiness, CardPushTransferNewParamsBusinessApplicationIdentifierMoneyTransferBankInitiated, CardPushTransferNewParamsBusinessApplicationIdentifierNonCardBillPayment, CardPushTransferNewParamsBusinessApplicationIdentifierConsumerBillPayment, CardPushTransferNewParamsBusinessApplicationIdentifierCardBillPayment, CardPushTransferNewParamsBusinessApplicationIdentifierFundsDisbursement, CardPushTransferNewParamsBusinessApplicationIdentifierFundsTransfer, CardPushTransferNewParamsBusinessApplicationIdentifierLoyaltyAndOffers, CardPushTransferNewParamsBusinessApplicationIdentifierMerchantDisbursement, CardPushTransferNewParamsBusinessApplicationIdentifierMerchantPayment, CardPushTransferNewParamsBusinessApplicationIdentifierPersonToPerson, CardPushTransferNewParamsBusinessApplicationIdentifierTopUp, CardPushTransferNewParamsBusinessApplicationIdentifierWalletTransfer:
		return true
	}
	return false
}

type CardPushTransferListParams struct {
	// Filter Card Push Transfers to ones belonging to the specified Account.
	AccountID param.Field[string]                              `query:"account_id"`
	CreatedAt param.Field[CardPushTransferListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Filter records to the one with the specified `idempotency_key` you chose for
	// that object. This value is unique across Increase and is used to ensure that a
	// request is only processed once. Learn more about
	// [idempotency](https://increase.com/documentation/idempotency-keys).
	IdempotencyKey param.Field[string] `query:"idempotency_key"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  param.Field[int64]                            `query:"limit"`
	Status param.Field[CardPushTransferListParamsStatus] `query:"status"`
}

// URLQuery serializes [CardPushTransferListParams]'s query parameters as
// `url.Values`.
func (r CardPushTransferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardPushTransferListParamsCreatedAt struct {
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

// URLQuery serializes [CardPushTransferListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r CardPushTransferListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardPushTransferListParamsStatus struct {
	// Filter Card Push Transfers by status. For GET requests, this should be encoded
	// as a comma-delimited string, such as `?in=one,two,three`.
	In param.Field[[]CardPushTransferListParamsStatusIn] `query:"in"`
}

// URLQuery serializes [CardPushTransferListParamsStatus]'s query parameters as
// `url.Values`.
func (r CardPushTransferListParamsStatus) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CardPushTransferListParamsStatusIn string

const (
	CardPushTransferListParamsStatusInPendingApproval   CardPushTransferListParamsStatusIn = "pending_approval"
	CardPushTransferListParamsStatusInCanceled          CardPushTransferListParamsStatusIn = "canceled"
	CardPushTransferListParamsStatusInPendingReviewing  CardPushTransferListParamsStatusIn = "pending_reviewing"
	CardPushTransferListParamsStatusInRequiresAttention CardPushTransferListParamsStatusIn = "requires_attention"
	CardPushTransferListParamsStatusInPendingSubmission CardPushTransferListParamsStatusIn = "pending_submission"
	CardPushTransferListParamsStatusInSubmitted         CardPushTransferListParamsStatusIn = "submitted"
	CardPushTransferListParamsStatusInComplete          CardPushTransferListParamsStatusIn = "complete"
	CardPushTransferListParamsStatusInDeclined          CardPushTransferListParamsStatusIn = "declined"
)

func (r CardPushTransferListParamsStatusIn) IsKnown() bool {
	switch r {
	case CardPushTransferListParamsStatusInPendingApproval, CardPushTransferListParamsStatusInCanceled, CardPushTransferListParamsStatusInPendingReviewing, CardPushTransferListParamsStatusInRequiresAttention, CardPushTransferListParamsStatusInPendingSubmission, CardPushTransferListParamsStatusInSubmitted, CardPushTransferListParamsStatusInComplete, CardPushTransferListParamsStatusInDeclined:
		return true
	}
	return false
}
